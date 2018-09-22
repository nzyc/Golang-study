package main

import "fmt"

func main () {
	// 1、每个原文件必须是 go 为后缀
	// 2、go语言严格区分大小写
	// 3、go 语句没有标点符号
	// 4、go 语言是逐行编译的所以一行只能写一个语句
	// 5、定义的变量 或者 import 引入的包 没有使用 go 的编译不能通过、
	// 6、大括号必须成对出现，否则编译不通过

	fmt.Println("hello Golang")
}
