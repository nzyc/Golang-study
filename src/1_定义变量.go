package main

import "fmt"

// 声明全局变量
var n1 int
var n2 int
var n3 int

// 一次性声明
var (
	n11 int
	n22 int
	n33 int
)

func main () {
	fmt.Println(n1,n2,n3)
	fmt.Println(n11,n22,n33)

	// 声明 局部变量
	var one1 int = 123		// 标准版
	var one2 = 1234			// 根据值自动推导 (类型推导)	只能在 函数里边使用
	one3 := 12345			// 类型推导的简化版  			只能在 函数里边使用
	fmt.Println(one1,one2,one3)

	// 声明多个变量
	var two1, two2, two3 int
	fmt.Println(two1,two2,two3)
	var two11, two22, two33 = 111,222,333
	fmt.Println(two11,two22,two33)
	two111, two222,two333 := "aaa", "bbb","ccc"
	fmt.Println(two111,two222,two333)

}
