package cmd

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

import (
	"fmt"
	"os"

	"github.com/shenguanjiejie/go-tools/v3"
	"github.com/shenguanjiejie/workwork/cmd/model"
	"github.com/spf13/cobra"
)

// md5Cmd represents the md5 command
var md5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "md5",
	Long:  `Multiple encode support, split with space. and file supported too.<br>支持用空格隔开,传入多个字符串md5,支持文件md5`,
	Run: func(cmd *cobra.Command, args []string) {

		file, err := cmd.Flags().GetString(model.FileFlag.Name)
		if err != nil {
			tools.Info(err)
			return
		}

		if file != "" {
			// RJ 对文件进行MD5加密
			fileBytes, err := os.ReadFile(file)
			if err != nil {
				tools.Info(err)
				return
			}

			fmt.Println(tools.MD5(string(fileBytes)))
			return
		}

		if len(args) == 0 {
			fmt.Println("未输入需要加密的内容")
		}

		for _, arg := range args {
			fmt.Println(tools.MD5(arg))
		}
	},
}

func init() {
	rootCmd.AddCommand(md5Cmd)
	command := new(model.Command[any])
	command.Title = md5Cmd.Short
	command.SubTitle = md5Cmd.Long
	command.Use = md5Cmd.Use
	model.AddFlag(md5Cmd, command, model.FileFlag)
	model.Commands = append(model.Commands, command)
}
