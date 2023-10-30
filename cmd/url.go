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
	"net/url"
	"workwork/cmd/model"

	"github.com/shenguanjiejie/go-tools/v3"
	"github.com/spf13/cobra"
)

// urlCmd represents the url command
var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "URL encode/decode.<br>URL编码/解码",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			tools.Info("没有输入URL")
			return
		}

		urlStr := args[0]

		decodeB, err := cmd.Flags().GetBool(model.DecodeFlag.Name)
		if err != nil {
			tools.Info(err)
			return
		}
		if decodeB {
			decodedStr, err := url.QueryUnescape(urlStr)
			if err != nil {
				tools.Info(err)
				return
			}
			fmt.Println(decodedStr)

			decodedPathVar, err := url.PathUnescape(urlStr)
			if err != nil {
				tools.Info(err)
				return
			}
			fmt.Println(decodedPathVar)
		} else {
			encodedStr := url.QueryEscape(urlStr)
			fmt.Println(encodedStr)
			encodedPathVar := url.PathEscape(urlStr)
			fmt.Println(encodedPathVar)
		}
	},
}

func init() {
	rootCmd.AddCommand(urlCmd)

	command := new(model.Command[any])
	command.Title = urlCmd.Short
	command.SubTitle = urlCmd.Long
	command.Use = urlCmd.Use
	model.AddFlag(urlCmd, command, model.DecodeFlag)
	model.Commands = append(model.Commands, command)
}

// url https://www.google.com.hk/search\?q\=中文搜索测试
// go run main.go url -d "https://www.google.com.hk/search\?q\=%E4%B8%AD%E6%96%87%E6%90%9C%E7%B4%A2%E6%B5%8B%E8%AF%95"
