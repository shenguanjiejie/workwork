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
	"fmt"
	"os"
	"regexp"
	"workwork/cmd/model"

	"github.com/shenguanjiejie/go-tools"
	"github.com/spf13/cobra"
)

// regexCmd represents the regex command
var regexCmd = &cobra.Command{
	Use:   "regex",
	Short: "Regex test tool.<br>正则表达式测试工具",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// tools.Slogln(args)
		wanna, err := cmd.Flags().GetString(model.WannaFlag.Name)
		if err != nil {
			tools.Slogln(err)
			return
		}

		if wanna != "" {
			wannaRegexStr := ""
			for _, s := range wanna {
				wannaRegexStr = wannaRegexStr + string(s) + ".*"
			}
			reg := regexp.MustCompile(`\{[(\n\s+//.+)|\n]*\n\s+title\: '.*` + wannaRegexStr + `\n.+\n.+(\n\s+counterExamples:.+)?[(\n\s+//.+)|\n]*\n\},?`)
			if reg == nil {
				tools.Slogln("正则表达式初始化失败")
				return
			}

			all := reg.FindAllString(model.RegexAll, 20)
			for _, str := range all {
				fmt.Println(str)
			}
			return
		}

		match, err := cmd.Flags().GetBool(model.MatchFlag.Name)
		if err != nil {
			tools.Slogln(err)
			return
		}

		count, err := cmd.Flags().GetInt(model.FindFlagCount.Name)
		if err != nil {
			tools.Slogln(err)
			return
		}

		file, err := cmd.Flags().GetString(model.FileFlag.Name)
		if err != nil {
			tools.Slogln(err)
			return
		}

		content := ""
		if file != "" {
			if len(args) == 0 {
				fmt.Println("需要指定正则表达式")
				return
			}
			fileBytes, err := os.ReadFile(file)
			if err != nil {
				tools.Slogln(err)
				return
			}

			content = string(fileBytes)
		} else {
			if len(args) < 2 {
				fmt.Println("需要指定正则表达式和所要匹配的内容")
				return
			}
			content = args[1]
		}

		regexStr := args[0]
		// tools.Slogln(regexStr, content)
		reg, err := regexp.Compile(regexStr)
		if err != nil {
			tools.Slogln(err)
			return
		}

		if match {
			match := reg.MatchString(content)
			fmt.Println(match)
		} else {
			results := reg.FindAllString(content, count)
			for i, result := range results {
				fmt.Println(i+1, ": "+result)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(regexCmd)

	command := new(model.Command[any])
	command.Title = regexCmd.Short
	command.SubTitle = regexCmd.Long
	command.Use = regexCmd.Use
	model.AddFlag(regexCmd, command, model.FileFlag)
	model.AddFlag(regexCmd, command, model.WannaFlag)
	model.AddFlag(regexCmd, command, model.MatchFlag)
	model.AddFlag(regexCmd, command, model.FindFlagCount)
	model.Commands = append(model.Commands, command)
}

// go run main.go regex -w 身份证1代
// go run main.go regex "Don't .+, Be H[a-z]+" "Don't W orry, Be Happy~~~"
// go run main.go regex -c 10 "\\w+" "Dont Worry, Be Happy~~~"
