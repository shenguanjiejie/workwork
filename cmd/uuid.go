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
	"workwork/cmd/model"

	"github.com/google/uuid"
	"github.com/shenguanjiejie/go-tools/v3"
	"github.com/spf13/cobra"
)

// uuidCmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "uuid",
	Long:  `Multiple support<br>支持输出多个uuid`,
	Run: func(cmd *cobra.Command, args []string) {

		count, _ := cmd.Flags().GetInt(model.UUIDFlagCount.Name)
		if count == 0 {
			count = 1
		}

		version, _ := cmd.Flags().GetInt(model.UUIDFlagVersion.Name)
		if version == 0 {
			version = 4
		}

		for i := 0; i < count; i++ {
			if version == 1 {
				uuid, err := uuid.NewUUID()
				if err != nil {
					tools.Info(err)
					return
				}
				fmt.Println(uuid.String())
			} else if version == 2 {
				uuid, err := uuid.NewDCEPerson()
				if err != nil {
					tools.Info(err)
					return
				}
				fmt.Println(uuid.String())
			} else if version == 3 {
				// uuid.NewMD5()
				fmt.Println("version 3 comming soon")
			} else if version == 4 {
				uuid, err := uuid.NewRandom()
				if err != nil {
					tools.Info(err)
					return
				}
				fmt.Println(uuid.String())
			} else if version == 5 {
				// uuid.NewSHA1()
				fmt.Println("version 5 comming soon")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)
	command := new(model.Command[any])
	command.Title = uuidCmd.Short
	command.SubTitle = uuidCmd.Long
	command.Use = uuidCmd.Use
	model.AddFlag(uuidCmd, command, model.UUIDFlagCount)
	model.AddFlag(uuidCmd, command, model.UUIDFlagVersion)
	model.Commands = append(model.Commands, command)
}
