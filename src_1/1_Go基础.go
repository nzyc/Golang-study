package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {

	/*
		定义变量
	*/
	// 定义单个变量
	var one1 int = 123            // 标准版
	var one2 = 1234               // 根据值自动推导 (类型推导)	只能在 函数里边使用
	one3 := 12345                 // 类型推导	的简化版  		只能在 函数里边使用
	fmt.Println(one1, one2, one3) // 123 1234 12345

	// 定义多个变量
	var two1, two2, two3 int
	fmt.Println(two1, two2, two3) // 0 0 0

	var two11, two22, two33 = 111, 222, 333
	fmt.Println(two11, two22, two33) // 111 222 333

	two111, two222, two333 := "aaa", "bbb", "ccc"
	fmt.Println(two111, two222, two333) // aaa bbb ccc

	/*
		定义常量
	*/
	// 跟定义变量一样 只不过 标识符不一样
	const constName int = 123
	fmt.Println(constName) // 123

	/*
		Boolean 布尔类型
	*/
	var isAction bool                              // 标准版
	var enabled, disabled = true, false            // 忽略类型声明
	flag := true                                   // 自动推导类型
	fmt.Println(isAction, enabled, disabled, flag) // false true false true

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
				complex64	 (32位实数+32位虚数
				复数的形式为RE + IMi，其中RE是实数部分，IM是虚数部分，而最后的i是虚数单位
	*/

	/*
		字符串
		格式 使用 双引号 "" 或者 反引号 ``
	*/
	var str1 string = "str1" // 标准
	var str2 = "str2"        // 简化
	str3 := "str3"           // 自动推导类型
	str4 := `str4`           // 自动推导类型
	str5 := str3 + str4      // 使用 + 拼接字符串
	fmt.Println(str1)        // str1
	fmt.Println(str2)        // str2
	fmt.Println(str3)        // str3
	fmt.Println(str4)        // str4
	fmt.Println(str5)        // str3str3

	/*
		错误类型
		内置有一个 errors的包 来处理错误
	*/
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err) // emit macho dwarf: elf header corrupted
	}

	/*
		分组声明
	*/
	var (
		fenz1     = 100
		fenz2 int = 200
		fenz3     = 300
	)
	const (
		fenz4     = 400
		fenz5 int = 500
		fenz6     = 600
	)
	fmt.Println(fenz1, fenz2, fenz3) // 100 200 300
	fmt.Println(fenz4, fenz5, fenz6) // 400 500 600

	/*
		iota枚举
		它的默认值是 0 const 没增加一行 就 加 1
		iota 在同一行时 值相同
	*/
	const (
		iota1               = iota
		iota2               = iota
		iota3               = iota
		iota4, iota5, iota6 = iota, iota, iota
		iota7               = iota
		iota8
		iota9  // 常亮 省略赋值操作的时候  默认与 前一个常量的值相同 所以 iota8 = iota iota9 = iota
	)
	fmt.Println(iota1, iota2, iota3, iota4, iota5, iota6, iota7, iota8, iota9) // 0 1 2 3 3 3 4 5 6

	/*
		数组
		不是引用类型
		定义：
			var arr [n]type
			n 是数组的类型， type 是 定义数组储存的类型
	*/
	var arr1 [10]int // 定义一个长度为10 的int类型的数组
	arr1[0] = 123    // 赋值操作 是从 0开始的
	arr1[1] = 1234
	fmt.Println(arr1) // [123 1234 0 0 0 0 0 0 0 0]  其他没有赋值的全都是 int 的默认值 0

	var arr2 = [3]int{1, 2, 3} // 标准版	定义吃长度为三的数组 他的值分别 为 1，2，3
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3} // 简化版
	fmt.Println(arr3)

	arr4 := [...]int{1, 2, 3, 4} // 忽略长度 采用 ... 定义 go会自动根据元素的长度来计算长度
	fmt.Println(arr4)

	// 定义二维数组	该数组以俩个数组为元素，其中每个数组又有四个int类型的元素
	arr5 := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	fmt.Println(arr5)

	// 二维数组 简化版
	arr6 := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(arr6) // [[1 2 3 4] [5 6 7 8]]

	/*
		数组 slice
		引用类型
		定义
			var arr []int
		ps: 和声明数组 一样 只不过 少了长度
			在初始定义数组时，我们并不知道需要多大的数组，因此我们就需要“动态数组”。在Go里面这种数据结构叫slice
			slice并不是真正意义上的动态数组，而是一个引用类型。
			slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。
		注意:
			slice 和 数组 的声明的区别
			数组的声明 [] 里必须要有 长度 或者 ...
			slice [] 里面没有任何字符
		内置函数
			len		获取slice长度
			cap		获取slice的最大容量
			append 	向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
			copy 	函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数
	*/
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1) // [1 2 3 4 5]

	// slice 也可以 从 slice 中再次声明
	var slice2 = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 声明俩个byte的 slice
	var slice3, slice4 []int
	// 进行截取
	slice3 = slice2[3:5]
	fmt.Println(slice3) // [3 4]
	// 进行截取
	slice4 = slice2[4:6]
	fmt.Println(slice4) // [4 5]

	// 简洁操作
	slice5 := slice2[2:3]
	slice6 := slice2[2:] // [2:len(slice2)]
	slice7 := slice2[:3] // [0:3]
	slice8 := slice2[:]  // [0:len(slice2)]

	fmt.Println(slice5) // [2]
	fmt.Println(slice6) // [2 3 4 5 6 7 8 9]
	fmt.Println(slice7) // [0 1 2]
	fmt.Println(slice8) // [0 1 2 3 4 5 6 7 8 9]

	// 内置函数的使用
	fmt.Println(len(slice8))         // 10
	fmt.Println(cap(slice8))         // 10
	fmt.Println(append(slice8, 999)) // [0 1 2 3 4 5 6 7 8 9 999]

	/*
		数组 map
		引用类型
		map 读取和设置 也类似 slice  通过 key 来操作  而 slice 的 index 是通过 int 类型 来操作的
	*/
	// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用 make初始化
	var numbers map[string]int
	// 使用 make 初始化
	numbers = make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	fmt.Println(numbers)        // map[one:1 two:2 three:3]
	fmt.Println(numbers["two"]) // 2

	// 第二种方式
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	fmt.Println(rating) // map[Python:4.5 C++:2 C:5 Go:4.5]

	// 第三种方式
	rating1 := make(map[string]int)
	rating1["hello"] = 1
	rating1["world"] = 2
	fmt.Println(rating1) // map[hello:1 world:2]

	/*
		小练习
	*/
	// 1
	for i := 1; i <= 10; i++ {
		fmt.Println("第一种---", i)
	}

	// 2
	ii := 1
Herr:
	fmt.Println("第二种---", ii)
	ii++
	if ii <= 10 {
		goto Herr
	}

	// 3
	var arrone = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i <= len(arrone); i++ {
		fmt.Println("第三种---", arrone[i-1])
	}

	// 4
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FizzBuzz ---", i)
		} else if i % 3 == 0 {
			fmt.Println("Fizz ---",i)
		} else if i % 5 == 0 {
			fmt.Println("Buzz ---",i)
		} else {
			fmt.Println(i)
		}
	}

	// 5
	for i:=1;i<=100;i++{
		for k:=1;k<=i;k++{
			fmt.Print("A")
		}
		fmt.Println("")
	}
	strone := "A"
	for i:=0;i<=100 ;i++  {
		fmt.Println(strone)
		strone += "A"
	}

	// 6

	str := "hello"
	b := []byte(str)
	fmt.Println(b)

	strtwo := "dsjkdshdjsdh....js"
	fmt.Printf("String %s\nLength: %d, Runes: %d\n", strtwo, len([]byte(strtwo)), utf8.RuneCount([]byte(strtwo)))


	// 7
	s := "foobar"
	a := []rune(s)
	for i,j:=0, len(a)-1; i<j; i,j=i+1,j-1 {
		a[i], a[j] = a[j], a[i]
		fmt.Printf("%s\n", string(a))
	}
}
