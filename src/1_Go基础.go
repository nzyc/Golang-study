package main

import (
	"errors"
	"fmt"
)

func main () {


	/*
		定义变量
	*/
	// 定义单个变量
	var one1 int = 123				// 标准版
	var one2 = 1234					// 根据值自动推导 (类型推导)	只能在 函数里边使用
	one3 := 12345					// 类型推导	的简化版  		只能在 函数里边使用
	fmt.Println(one1,one2,one3)

	// 定义多个变量
	var two1, two2, two3 int
	fmt.Println(two1,two2,two3)

	var two11, two22, two33 = 111,222,333
	fmt.Println(two11,two22,two33)

	two111, two222,two333 := "aaa", "bbb","ccc"
	fmt.Println(two111,two222,two333)


	/*
		定义常量
	*/
	// 跟定义变量一样 只不过 标识符不一样
	const constName int =  123
	fmt.Println(constName)


	/*
		Boolean 布尔类型
	*/
	var isAction bool						// 标准版
	var enabled, disabled = true, false		// 忽略类型声明
	flag := true							// 自动推导类型
	fmt.Println(isAction, enabled, disabled, flag)


	/*
		数值类型

			整型
				rune, int8, int16, int32, int64
				byte, uint8,  uint16,  uint32,  uint64
				rune 是 int32 的别称
				byte 是 uint8 的别称
				int 与 uint 这俩种的类型长度一样 具体长度取决于 不同编译器的实现

			浮点型
				float32
				float64	(默认)

			复数
				complex128	（64位实数+64位虚数）
				complex63	 (32位实数+32位虚数
				复数的形式为RE + IMi，其中RE是实数部分，IM是虚数部分，而最后的i是虚数单位
	*/


	/*
		字符串
		格式 使用 双引号 "" 或者 反引号 ``
	*/
	var str1 string	= "str1"		// 标准
	var str2 = "str2"				// 简化
	str3 := "str3"					// 自动推导类型
	str4 := `str3`					// 自动推导类型
	str5 := str3 + str4				// 使用 + 拼接字符串
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(str4)
	fmt.Println(str5)


	/*
		错误类型
		内置有一个 errors的包 来处理错误
	*/
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err)
	}


	/*
		分组声明
	*/
	var (
		fenz1 = 100
		fenz2 = 200
		fenz3 = 300
	)
	const (
		fenz4 = 400
		fenz5 = 500
		fenz6 = 600
	)
	fmt.Println(fenz1,fenz2,fenz3)
	fmt.Println(fenz4,fenz5,fenz6)


	/*
		iota枚举
		它的默认值是 0 const 没增加一行 就 加 1
	*/


}
