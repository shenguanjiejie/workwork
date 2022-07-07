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
	"workwork/cmd/model"

	"github.com/spf13/cobra"
)

// readmeCmd represents the readme command
var readmeCmd = &cobra.Command{
	Use:   "readme",
	Short: "readme",
	Long:  `readme`,
	Run: func(cmd *cobra.Command, args []string) {
		readme := `
## workwork
开发/测试常用工具的命令行实现
`

		for _, cmd := range model.Commands {
			readme = readme + fmt.Sprintf("# %s\n%s\n%s\n", cmd.Use, cmd.Title, cmd.SubTitle)
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
