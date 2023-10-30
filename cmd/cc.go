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

	"github.com/shenguanjiejie/go-tools/v3"
	"github.com/spf13/cobra"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

// ccCmd represents the cc command
var ccCmd = &cobra.Command{
	Use:   "cc",
	Short: "Calculater<br>计算器",
	Long:  `Based on Go interpreter, supports various operators, logical operations, and bitwise operations. <br>使用Go解释器实现,支持各种运算符,逻辑运算以及位运算`,
	Run: func(cmd *cobra.Command, args []string) {

		// intNumber := `([1-9]\d+)`
		// floatNumber := `(\d+\.\d+)`
		// number := fmt.Sprintf("(%s|%s)", intNumber, floatNumber)

		// operator := `[+\-*/\^]`
		// expression := fmt.Sprintf("(%s%s(%s%s)*%s)", number,operator,number,operator,number)
		// bracket := fmt.Sprintf(`(\(%s?\))`,expression)
		// bracketSquare := fmt.Sprintf(`(\(%s?\))`,expression)
		if len(args) == 0 {
			fmt.Println("至少要输入一个表达式")
			return
		}

		i := interp.New(interp.Options{})

		i.Use(stdlib.Symbols)

		_, err := i.Eval(`import "fmt"`)
		if err != nil {
			tools.Info(err)
			return
		}

		for _, arg := range args {
			evalStr := fmt.Sprintf("fmt.Println(%s)", arg)
			_, err = i.Eval(evalStr)
			if err != nil {
				tools.Info(err)
				continue
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(ccCmd)

	command := new(model.Command[any])
	command.Title = ccCmd.Short
	command.SubTitle = ccCmd.Long
	command.Use = ccCmd.Use
	model.Commands = append(model.Commands, command)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ccCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ccCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// go run main.go cc '1.5+2*(1+2)' '100==100' '1+1>3' '1==1 && 1==2' '5%2' '1<<3'
