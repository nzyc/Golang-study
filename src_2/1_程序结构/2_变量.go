package main

import "fmt"

func main() {
	fmt.Println("hello world")
	/*
		声明变量的格式
		var 变量名 类型 = 表达式
	*/

	// 第一种
	var str1 string
	fmt.Println(str1) // ""

	// 第二种
	var a, b, c int      // int
	fmt.Println(a, b, c) // 0 0 0

	// 第三种	（自动推导）
	var d, e, f = true, 100, "str"
	fmt.Println(d, e, f)

	// 第四种 	（简短声明）
	g := "str"
	h, i, j := true, 100, "str"
	fmt.Println(g, h, i, j)

	/*
		指针
	*/
	one := 1
	one1 := &one
	fmt.Println(*one1) // 1
	*one1 = 2
	fmt.Println(*one1) // 2

	one2 := 1
	incr(&one2)
	fmt.Println(incr(&one2)) // 3


	/*
		new 函数
		相当于 testNew 函数
		new 函数是一个语法糖 不是一个新概念
	*/
	two1 := new(int)
	two2 := new(bool)
	two3 := new(string)
	fmt.Println(*two1,*two2,*two3)	// 0 false ""


	/*
		变量的声明周期

		变量的生命周期是在程序运行期间 变量存在的时间间隔
		对于包 里面 一级声明的变量 或者说是 全局变量 他们是的生命周期跟整个程序的生命周期是一样的
		局部变量的生命周期就是动态的 从每一次使用时创建，直到这个变量不在被引用时为止
		函数 里的 参数 返回值 也是局部变量
	*/
}

func incr(p *int) int {
	*p++
	return *p
}

func testNew () *int{
	var asd int
	return &asd
}
