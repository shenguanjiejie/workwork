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
	"workwork/cmd/model"

	"github.com/shenguanjiejie/go-tools"
	"github.com/spf13/cobra"
)

// readmeCmd represents the readme command
var readmeCmd = &cobra.Command{
	Use:   "readme",
	Short: "readme",
	Long:  `readme`,
	Run: func(cmd *cobra.Command, args []string) {
		version := "0.0.1"
		download := "or<br>[Download](https://github.com/shenguanjiejie/workwork/releases), Unpack, and move \"ww\" to /usr/local/bin.\n到[release页面](https://github.com/shenguanjiejie/workwork/releases)下载后, 解压出ww文件, 放在/usr/local/bin目录下即可\n"
		readme := fmt.Sprintf("# %s\n%s\n", rootCmd.Use, rootCmd.Short)
		readme = readme + fmt.Sprintf("# Install\n##### Mac\n```shell\nbrew install shenguanjiejie/tap/workwork\n```\nor\n```shell\ncurl -LO https://github.com/shenguanjiejie/workwork/releases/download/v%s/workwork_%s_darwin.tar.gz && tar -zxvf ./workwork_%s_darwin.tar.gz && mv ./ww /usr/local/bin && rm ./workwork_%s_darwin.tar.gz\n```\n%s", version, version, version, version, download)
		readme = readme + fmt.Sprintf("##### Linux\n```shell\ncurl -LO https://github.com/shenguanjiejie/workwork/releases/download/v%s/workwork_%s_linux_x86_64.tar.gz && tar -zxvf ./workwork_%s_linux_x86_64.tar.gz && mv ./ww /usr/local/bin && rm ./workwork_%s_linux_x86_64.tar.gz\n```\n%s", version, version, version, version, download)

		for _, cmd := range model.Commands {
			if cmd.SubTitle == "" {
				readme = readme + fmt.Sprintf("# %s\n%s\n", cmd.Use, cmd.Title)
			} else {
				readme = readme + fmt.Sprintf("# %s\n%s<br>%s\n", cmd.Use, cmd.Title, cmd.SubTitle)
			}

			if len(cmd.FlagIntArr)+len(cmd.FlagStringArr)+len(cmd.FlagBoolArr) > 0 {
				readme = readme + `
|params(参数)|shorthand(缩写)|default(默认值)|usage(说明)|
|---|---|---|---|
`
			}

			for _, flagInfo := range cmd.FlagIntArr {
				readme = readme + fmt.Sprintf("|--%s|-%s|%d|%s|\n", flagInfo.Name, flagInfo.Shorthand, flagInfo.Value, flagInfo.Usage)
			}

			for _, flagInfo := range cmd.FlagStringArr {
				readme = readme + fmt.Sprintf("|--%s|-%s|%s|%s|\n", flagInfo.Name, flagInfo.Shorthand, flagInfo.Value, flagInfo.Usage)
			}

			for _, flagInfo := range cmd.FlagBoolArr {
				readme = readme + fmt.Sprintf("|--%s|-%s|%v|%s|\n", flagInfo.Name, flagInfo.Shorthand, flagInfo.Value, flagInfo.Usage)
			}

			readme = readme + fmt.Sprintf("\n![%s](resources/%s.png)\n", cmd.Use, cmd.Use)
		}

		readme = readme + `
# TODO:
1. 默认保存路径配置, 默认读取文件路径配置. (Default I/O path config)
2. 单元测试. (Unit testing)
3. Alfred支持. (Alfred support)
4. color command, 色值转换(类似"time")
5. ...
`
		err := os.WriteFile("README.md", []byte(readme), 0755)
		if err != nil {
			tools.Slogln(err)
			return
		}
		fmt.Println(readme)
	},
}

func init() {
	rootCmd.AddCommand(readmeCmd)
	rootCmd.Hidden = true
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readmeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readmeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
