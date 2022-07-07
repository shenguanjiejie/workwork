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
package model

import (
	"reflect"

	"github.com/spf13/cobra"
)

type Command struct {
	Title         string          `json:"title" bson:"title"`
	SubTitle      string          `json:"sub_title" bson:"sub_title"`
	Use           string          `json:"use" bson:"use"`
	FlagIntArr    []*Flag[int]    `json:"flag_int_arr" bson:"flag_int_arr"`
	FlagStringArr []*Flag[string] `json:"flag_string_arr" bson:"flag_string_arr"`
	FlagBoolArr   []*Flag[bool]   `json:"flag_bool_arr" bson:"flag_bool_arr"`
}

var Commands = []*Command{}

func AddFlag[T FlagT](cmd *cobra.Command, command *Command, flag *Flag[T]) {
	v := reflect.ValueOf(flag.Value)
	flagV := reflect.ValueOf(flag)

	switch v.Kind() {
	case reflect.Int:
		cmd.PersistentFlags().IntP(flag.Name, flag.Shorthand, v.Interface().(int), flag.Usage)
		if command.FlagIntArr == nil {
			command.FlagIntArr = make([]*Flag[int], 0)
		}
		command.FlagIntArr = append(command.FlagIntArr, flagV.Interface().(*Flag[int]))
	case reflect.String:
		cmd.PersistentFlags().StringP(flag.Name, flag.Shorthand, v.Interface().(string), flag.Usage)
		if command.FlagStringArr == nil {
			command.FlagStringArr = make([]*Flag[string], 0)
		}
		command.FlagStringArr = append(command.FlagStringArr, flagV.Interface().(*Flag[string]))
	case reflect.Bool:
		cmd.PersistentFlags().BoolP(flag.Name, flag.Shorthand, v.Interface().(bool), flag.Usage)
		if command.FlagBoolArr == nil {
			command.FlagBoolArr = make([]*Flag[bool], 0)
		}
		command.FlagBoolArr = append(command.FlagBoolArr, flagV.Interface().(*Flag[bool]))
	}

}
