package main

import "fmt"

func main() {
	// 1、每个原文件必须是 go 为后缀
	// 2、go语言严格区分大小写
	// 3、go 语句没有标点符号
	// 4、go 语言是逐行编译的所以一行只能写一个语句
	// 5、定义的变量 或者 import 引入的包 没有使用 go 的编译不能通过、
	// 6、大括号必须成对出现，否则编译不通过
	// 7、大写字母的变量是可以导出的，也就四其他的包可以读取，是公有的变量，小写字母开头的是 私有变量，是不可以导出的
	// 8、大写字母的函数也是一样，大写的相等于 class 中带有 public 关键词的公有函数，小写字母开头的就是private 关键词的私有函数

	fmt.Println("hello Golang")
}
