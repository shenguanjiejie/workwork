/*
Copyright © 2022 shenguanjiejie <835166018@qq.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
	"workwork/cmd/model"

	"github.com/shenguanjiejie/go-tools"
	"github.com/spf13/cobra"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "Base64 encode/decode. <br>Base64编码/解码",
	Long:  `Multiple encode/decode support, split with space.<br>支持用空格隔开, 一次进行多个编码/解码`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString(model.FileFlag.Name)
		if err != nil {
			tools.Slogln(err)
			return
		}
		decodeB, err := cmd.Flags().GetBool(model.DecodeFlag.Name)
		if err != nil {
			tools.Slogln(err)
			return
		}

		imageB, err := cmd.Flags().GetBool(model.ImageFlagBase64.Name)
		if err != nil {
			tools.Slogln(err)
			return
		}

		// RJ 对文件进行base64编码
		if file != "" {
			fileBytes, err := os.ReadFile(file)
			if err != nil {
				tools.Slogln(err)
				return
			}

			if decodeB {
				base64Decode(fileBytes, imageB)
			} else {
				base64Encode(fileBytes, imageB)
			}

			return
		}

		if len(args) == 0 {
			fmt.Println("未输入需要编码的内容")
		}

		for _, arg := range args {
			arg = strings.TrimSpace(arg)

			if decodeB {
				base64Decode([]byte(arg), imageB)
			} else {
				base64Encode([]byte(arg), imageB)
			}
		}
	},
}

func base64Decode(content []byte, imageB bool) error {

	imageReg := regexp.MustCompile(`^data:image\/(?:gif|png|jpeg|bmp|webp|svg\+xml)(?:;charset=utf-8)?;base64,`)
	if imageReg == nil {
		err := errors.New("imageReg初始化失败")
		tools.Slogln(err)
		return err
	}

	reg := regexp.MustCompile(`(^data:image\/(?:gif|png|jpeg|bmp|webp|svg\+xml)(?:;charset=utf-8)?;base64,)?(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4})$`)
	if reg == nil {
		err := errors.New("reg初始化失败")
		tools.Slogln(err)
		return err
	}

	match := reg.Match(content)
	if !match {
		err := errors.New("非base64格式内容, 无法进行decode")
		tools.Slogln(err)
		return err
	}

	// RJ 2022-07-04 19:26:10 去掉图片base64的前缀
	prefixByte := imageReg.Find(content)
	content = content[len(prefixByte):]
	decodeResult, err := base64.StdEncoding.DecodeString(string(content))
	if err != nil {
		tools.Slogln(err)
		return err
	}

	// RJ 2022-07-04 18:56:17 图片格式, 存储图片到本地
	if imageB {
		fName := time.Now().Format("20060102150405") + ".png"

		f, err := os.OpenFile(fName, os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.Write(decodeResult)
		fPath, err := os.Getwd()
		if err != nil {
			tools.Slogln(err)
			return err
		}
		fmt.Println("图片已保存到: " + fPath + "/" + fName)

		return nil
	}

	finalStr := string(decodeResult)
	fmt.Println(finalStr)

	return nil
}

func base64Encode(content []byte, imageB bool) {

	encodeStr := base64.StdEncoding.EncodeToString(content)
	// RJ 2022-07-04 18:38:45 转为图片格式的base64
	if imageB {
		encodeStr = "data:image/png;base64," + encodeStr
	}
	fmt.Println(encodeStr)
}

func init() {
	rootCmd.AddCommand(base64Cmd)

	command := new(model.Command[any])
	command.Title = base64Cmd.Short
	command.SubTitle = base64Cmd.Long
	command.Use = base64Cmd.Use
	model.AddFlag(base64Cmd, command, model.FileFlag)
	model.AddFlag(base64Cmd, command, model.DecodeFlag)
	model.AddFlag(base64Cmd, command, model.ImageFlagBase64)
	model.Commands = append(model.Commands, command)
}

// ww base64 123
// go run main.go base64 -d MTIz
// go run main.go base64 -f '/Users/rj/Desktop/2.png'
// go run main.go base64 -i -f '/Users/rj/Desktop/2.png'
// go run main.go base64 -d -i "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAsAAAAECAYAAABY+sXzAAAAAXNSR0IArs4c6QAAAIRlWElmTU0AKgAAAAgABQESAAMAAAABAAEAAAEaAAUAAAABAAAASgEbAAUAAAABAAAAUgEoAAMAAAABAAIAAIdpAAQAAAABAAAAWgAAAAAAAABIAAAAAQAAAEgAAAABAAOgAQADAAAAAQABAACgAgAEAAAAAQAAAAugAwAEAAAAAQAAAAQAAAAA8NSg1AAAAAlwSFlzAAALEwAACxMBAJqcGAAAAVlpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IlhNUCBDb3JlIDYuMC4wIj4KICAgPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4KICAgICAgPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIKICAgICAgICAgICAgeG1sbnM6dGlmZj0iaHR0cDovL25zLmFkb2JlLmNvbS90aWZmLzEuMC8iPgogICAgICAgICA8dGlmZjpPcmllbnRhdGlvbj4xPC90aWZmOk9yaWVudGF0aW9uPgogICAgICA8L3JkZjpEZXNjcmlwdGlvbj4KICAgPC9yZGY6UkRGPgo8L3g6eG1wbWV0YT4KGV7hBwAAALVJREFUCB1jcGj9EuPe/DmbYep/nv8MDKwMQPBssy8XiAYBneV5VgxrXZRAbBbOf/84t3Pwrq5+9bD8vQLDs9dB9pwiApt/3ljfzWp3ouTnla9HVc8es+FR17a9zsTM9P8nw99/obx/fn5lfcXwn4GFlY+BiYHp7192PnZmBg7237wcnD8YmP6wMH5jYPj/nzGg+YMyyJoXgQxiQAHG/8cYhI6tOsZp1fK91bT7iwtI7klgtjAAs/tBHgdMJN4AAAAASUVORK5CYII="
// go run main.go base64 -d -i -f '/Users/rj/Desktop/base64.txt'
