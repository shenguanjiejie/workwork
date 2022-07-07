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

type FlagT interface {
	~int | ~string | ~bool
}

type Flag[T FlagT] struct {
	Name      string
	Shorthand string
	Value     T
	Usage     string
}

var FileFlag = &Flag[string]{
	Name:      "file",
	Shorthand: "f",
	Value:     "",
	Usage:     "指定文件的全路径",
}

var DecodeFlag = &Flag[bool]{
	Name:      "decode",
	Shorthand: "d",
	Value:     false,
	Usage:     "是否是要解码",
}

var ImageFlagBase64 = &Flag[bool]{
	Name:      "image",
	Shorthand: "i",
	Value:     false,
	Usage:     "是否是对图片的编解码, 对图片编码会增加\"data:image/png;base64,\"前缀, 对base64进行图片解码会保存png文件到当前目录",
}

var UnixFlag = &Flag[bool]{
	Name:      "unix",
	Shorthand: "u",
	Value:     false,
	Usage:     "指定时间戳, 支持秒级和毫秒级时间戳",
}

var WannaFlag = &Flag[string]{
	Name:      "wanna",
	Shorthand: "w",
	Value:     "",
	Usage:     "查找常用的正则表达式",
}

var MatchFlag = &Flag[bool]{
	Name:      "match",
	Shorthand: "m",
	Value:     false,
	Usage:     "能否找到正则匹配的项, 返回true或者false",
}

// 默认使用find策略
// var FindFlagRegex = &Flag[bool]{
// 	Name:      "find",
// 	Shorthand: "",
// 	Value:     false,
// 	Usage:     "是否要用正则表达式查找文本内匹配的内容, 默认返回第一个匹配项, 如果要返回多个, 使用--count(缩写-c)指定最大的返回数量",
// }

var FindFlagCount = &Flag[int]{
	Name:      "count",
	Shorthand: "c",
	Value:     1,
	Usage:     "使用查找匹配文本内容时, 用该int值指定返回的最大匹配数量, 默认只返回第一个匹配项",
}
