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
	"regexp"
	"strconv"
	"time"
	"workwork/cmd/model"

	"github.com/shenguanjiejie/go-tools"
	"github.com/spf13/cobra"
)

var timeFormats = []string{
	"20060102150405",
	"2006-01-02 15:04:05",
	"2006/01/02 15:04:05",
	"2006年01月02日 15时04分05秒",
	"2006年1月2日 3:4:5",
	"2006-01-02T15:04:05.999Z", // UTC, ISO 8601, YYYY-MM-DDThh:mm:ss.sssZ
	time.ANSIC,
}

var weekMap = map[time.Weekday]string{
	time.Sunday:    "周天",
	time.Monday:    "周一",
	time.Tuesday:   "周二",
	time.Wednesday: "周三",
	time.Thursday:  "周四",
	time.Friday:    "周五",
	time.Saturday:  "周六",
}

// `^[0-9]{4}$`
// `^[0-9]{4}(-|\/)([1-9]|0[1-9]|1[0-2])$`
// `^[0-9]{4}(-|\/)([1-9]|0[1-9]|1[0-2])(-|\/)([1-9]|0[1-9]|[1-2][0-9]|3[0-1])$`
// `^[0-9]{4}(-|\/)([1-9]|0[1-9]|1[0-2])(-|\/)([1-9]|0[1-9]|[1-2][0-9]|3[0-1])\s(0|[0-1][0-9]|2[0-3])$`
// `^[0-9]{4}(-|\/)([1-9]|0[1-9]|1[0-2])(-|\/)([1-9]|0[1-9]|[1-2][0-9]|3[0-1])\s(0|[0-1][0-9]|2[0-3]):?(0|[0-5][0-9]|60)$`
// `^[0-9]{4}(-|\/)([1-9]|0[1-9]|1[0-2])(-|\/)([1-9]|0[1-9]|[1-2][0-9]|3[0-1])\s(0|[0-1][0-9]|2[0-3]):?((0|[0-5][0-9]):?(0|[0-5][0-9])|6000|60:00)$`

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Time format tool.<br>时间转换工具",
	Long:  `Multiple support, split with space.<br>支持用空格隔开, 一次进行多个时间转换`,
	Run: func(cmd *cobra.Command, args []string) {

		unixB, err := cmd.Flags().GetBool(model.UnixFlag.Name)
		if err != nil {
			tools.Slogln(err)
			return
		}

		date := time.Now()

		if len(args) == 0 {
			printDate(date)
		}

		for _, arg := range args {
			if arg == "" {
				tools.Slogln("错误的时间格式")
				continue
			}

			argInt, err := strconv.Atoi(arg)

			// RJ 2022-07-05 21:04:23 是数字
			if err == nil {
				length := len(arg)
				if unixB {
					// RJ 2022-07-05 21:25:36 时间戳转换
					if length == 13 {
						date = time.UnixMilli(int64(argInt))
					} else if length <= 10 {
						date = time.Unix(int64(argInt), 0)
					} else {
						tools.Slogln("错误的时间戳格式")
						continue
					}
				} else {
					// RJ 2022-07-05 21:25:49 20060102150405格式时间转换
					layout := timeFormats[0][:length]
					date, err = time.ParseInLocation(layout, arg, time.Local)
					if err != nil {
						tools.Slogln("错误的时间格式")
						continue
					}
				}
				printDate(date)
				continue
			}

			dateStr := ""
			subStrArrAll := regexp.MustCompile(`[-/ \:年月日时分秒]{1}`).Split(arg, -1)
			for _, subStr := range subStrArrAll {
				if subStr != "" {
					dateStr = dateStr + subStr
				}
			}
			layout := timeFormats[0][:len(dateStr)]
			date, err = time.ParseInLocation(layout, dateStr, time.Local)
			if err != nil {
				tools.Slogln("错误的时间格式")
				continue
			}
			printDate(date)
		}
	},
}

func printDate(date time.Time) {
	fmt.Print(date.Unix())
	fmt.Print("    ")
	for i, format := range timeFormats {
		if i == 4 {
			if date.Hour() > 10 {
				fmt.Print(date.Format("2006年1月2日 15:4:5"))
			} else {
				fmt.Print(date.Format(format))
			}
			fmt.Print(" " + weekMap[date.Weekday()])
		} else if i == 5 {
			utcDate := date.UTC()
			fmt.Print(utcDate.Format(format))
		} else {
			fmt.Print(date.Format(format))
		}

		if i != len(timeFormats)-1 {
			fmt.Print("    ")
		}
	}
	fmt.Println()
}

func init() {
	rootCmd.AddCommand(timeCmd)

	command := new(model.Command[any])
	command.Title = timeCmd.Short
	command.SubTitle = timeCmd.Long
	command.Use = timeCmd.Use
	model.AddFlag(timeCmd, command, model.UnixFlag)
	model.Commands = append(model.Commands, command)
}

// go run main.go time  20010101 '2002-02-02 02:02' '2003年03月03日 03分' '2004/04/04 04:04:04'
// go run main.go time 20010101020202
// go run main.go time -u 978314522
// go run main.go time -u 978314522333
