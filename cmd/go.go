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
	"workwork/cmd/model"

	"github.com/shenguanjiejie/go-tools"
	"github.com/spf13/cobra"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "使用字符串, 像脚本一样直接运行Go代码",
	Long:  `fmt,time,os,math包默认引入, 其他包需单独import`,
	Run: func(cmd *cobra.Command, args []string) {
		i := interp.New(interp.Options{})

		i.Use(stdlib.Symbols)

		_, err := i.Eval(`
		import (
			"fmt"
			"time"
			"os"
			"math"
		)`)
		if err != nil {
			tools.Slogln(err)
			return
		}

		for _, arg := range args {
			_, err = i.Eval(arg)
			if err != nil {
				tools.Slogln(err)
				continue
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(goCmd)

	command := new(model.Command)
	command.Title = goCmd.Short
	command.SubTitle = goCmd.Long
	command.Use = goCmd.Use
	model.Commands = append(model.Commands, command)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// go run main.go go 'fmt.Println(time.Now().Unix())'
// go run main.go go 'import "strconv"' 'fmt.Println(strconv.Itoa(2333))'
