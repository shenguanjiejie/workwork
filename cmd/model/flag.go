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

type Flag[T any] struct {
	Name      string
	Shorthand string
	Value     T
	Usage     string
}

var FileFlag = &Flag[any]{
	Name:      "file",
	Shorthand: "f",
	Value:     "",
	Usage:     "Specify full file path.<br>指定文件的全路径",
}

var DecodeFlag = &Flag[any]{
	Name:      "decode",
	Shorthand: "d",
	Value:     false,
	Usage:     "Decode flag. <br>是否是要解码",
}

var ImageFlagBase64 = &Flag[any]{
	Name:      "image",
	Shorthand: "i",
	Value:     false,
	Usage:     "Image flag will append \"data:image/png;base64,\" to header of decode result, and save image to current path after encode.<br>是否是对图片的编解码, 对图片编码会增加\"data:image/png;base64,\"前缀, 对base64进行图片解码会保存png文件到当前目录",
}

var UnixFlag = &Flag[any]{
	Name:      "unix",
	Shorthand: "u",
	Value:     false,
	Usage:     "Input a Unix time or millisecond Unix time. <br>指定时间戳, 支持秒级和毫秒级时间戳",
}

var WannaFlag = &Flag[any]{
	Name:      "wanna",
	Shorthand: "w",
	Value:     "",
	Usage:     "Find common regex. <br>查找常用的正则表达式",
}

var MatchFlag = &Flag[any]{
	Name:      "match",
	Shorthand: "m",
	Value:     false,
	Usage:     "Return true or false by match result. <br>能否找到正则匹配的项, 返回true或者false",
}

// 默认使用find策略
// var FindFlagRegex = &Flag[bool]{
// 	Name:      "find",
// 	Shorthand: "",
// 	Value:     false,
// 	Usage:     "是否要用正则表达式查找文本内匹配的内容, 默认返回第一个匹配项, 如果要返回多个, 使用--count(缩写-c)指定最大的返回数量",
// }

var FindFlagCount = &Flag[any]{
	Name:      "count",
	Shorthand: "c",
	Value:     1,
	Usage:     "Specify the maximum number of matches. return the first one by default.<br> 使用查找匹配文本内容时, 用该int值指定返回的最大匹配数量, 默认只返回第一个匹配项",
}

var UUIDFlagCount = &Flag[any]{
	Name:      "count",
	Shorthand: "c",
	Value:     1,
	Usage:     "UUID count.<br> UUID 个数",
}

var UUIDFlagVersion = &Flag[any]{
	Name:      "version",
	Shorthand: "v",
	Value:     4,
	Usage:     "UUID version, default: 4.<br> UUID 版本, 默认: 4",
}
