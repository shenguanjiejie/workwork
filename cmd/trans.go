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
	"regexp"
	"workwork/cmd/model"

	"github.com/shenguanjiejie/go-tools/v3"
	"github.com/spf13/cobra"
)

// transCmd represents the trans command
var transCmd = &cobra.Command{
	Use:   "trans",
	Short: "Chinese/English Translate tool.<br>汉英/英汉翻译",
	Long:  `更详细的单词释义音标等, 后续考虑加入`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			return
		}
		chineseReg, err := regexp.Compile(`.*(\p{Han})+.*`)

		if err != nil {
			tools.Info(err)
			return
		}

		for _, arg := range args {
			isChinese := chineseReg.MatchString(arg)

			values := make(url.Values, 0)
			values.Set("q", arg)
			values.Set("from", "auto")
			if isChinese {
				values.Set("to", "en")
			} else {
				values.Set("to", "zh")
			}
			values.Set("appid", "20220706001266131")
			values.Set("salt", "1435660288")
			values.Set("sign", tools.MD5(values.Get("appid")+arg+values.Get("salt")+"_C0GKlwvijHwoj9GpWFH"))
			transInfo := new(model.TransInfo)
			err := tools.Get("https://fanyi-api.baidu.com/api/trans/vip/translate", values, transInfo)
			if err != nil {
				tools.Info(err)
				return
			}

			if len(transInfo.TransResult) == 0 {
				return
			}
			fmt.Println(transInfo.TransResult[0].Dst)
		}

		//TODO RJ 2022-07-07 00:20:24 接入有道翻译
	},
}

func init() {
	rootCmd.AddCommand(transCmd)

	command := new(model.Command[any])
	command.Title = transCmd.Short
	command.SubTitle = transCmd.Long
	command.Use = transCmd.Use
	model.Commands = append(model.Commands, command)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// go run main.go trans "不要问为什么翻译的功能不够给力, 因为给力的翻译api都收费~"
