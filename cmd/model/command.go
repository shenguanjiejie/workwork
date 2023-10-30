package model

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
	"reflect"

	"github.com/spf13/cobra"
)

type Command[T any] struct {
	Title    string     `json:"title" bson:"title"`
	SubTitle string     `json:"sub_title" bson:"sub_title"`
	Use      string     `json:"use" bson:"use"`
	Flags    []*Flag[T] `json:"flags" bson:"flags"`
}

var Commands = []*Command[any]{}

func AddFlag(cmd *cobra.Command, command *Command[any], flag *Flag[any]) {
	v := reflect.ValueOf(flag.Value)

	switch v.Kind() {
	case reflect.Int:
		cmd.PersistentFlags().IntP(flag.Name, flag.Shorthand, v.Interface().(int), flag.Usage)
	case reflect.String:
		cmd.PersistentFlags().StringP(flag.Name, flag.Shorthand, v.Interface().(string), flag.Usage)
	case reflect.Bool:
		cmd.PersistentFlags().BoolP(flag.Name, flag.Shorthand, v.Interface().(bool), flag.Usage)
	}

	if command.Flags == nil {
		command.Flags = make([]*Flag[any], 0)
	}
	command.Flags = append(command.Flags, flag)

}
