## workwork
开发/测试常用工具的命令行实现
# base64
base64编码/解码
支持用空格隔开, 一次进行多个编码/解码

|params(参数)|shorthand(缩写)|default(默认值)|usage(说明)|
|---|---|---|---|
|--file|-f||指定文件的全路径|
|--decode|-d|false|是否是要解码|
|--image|-i|false|是否是对图片的编解码, 对图片编码会增加"data:image/png;base64,"前缀, 对base64进行图片解码会保存png文件到当前目录|
# cc
计算器
使用Go解释器实现,支持各种运算符,逻辑运算以及位运算
# go
使用字符串, 像脚本一样直接运行Go代码
fmt,time,os,math包默认引入, 其他包需单独import
# md5
md5
支持用空格隔开,传入多个字符串md5,支持文件md5

|params(参数)|shorthand(缩写)|default(默认值)|usage(说明)|
|---|---|---|---|
|--file|-f||指定文件的全路径|
# regex
正则表达式测试工具
正则表达式测试工具

|params(参数)|shorthand(缩写)|default(默认值)|usage(说明)|
|---|---|---|---|
|--count|-c|1|使用查找匹配文本内容时, 用该int值指定返回的最大匹配数量, 默认只返回第一个匹配项|
|--file|-f||指定文件的全路径|
|--wanna|-w||查找常用的正则表达式|
|--match|-m|false|能否找到正则匹配的项, 返回true或者false|
# time
时间转换工具
支持用空格隔开, 一次进行多个时间转换

|params(参数)|shorthand(缩写)|default(默认值)|usage(说明)|
|---|---|---|---|
|--unix|-u|false|指定时间戳, 支持秒级和毫秒级时间戳|
# trans
汉英/英汉翻译
更详细的单词释义音标等, 后续考虑加入
# url
url编码/解码
url编码/解码

|params(参数)|shorthand(缩写)|default(默认值)|usage(说明)|
|---|---|---|---|
|--decode|-d|false|是否是要解码|