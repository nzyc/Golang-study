package main

import "fmt"

func main() {

	/*
		if 判断
	*/
	// 第一种方式
	var ifOne int = 1
	if ifOne == 1 {
		fmt.Println("这里是1")
	} else if ifOne == 2 {
		fmt.Println("这里是2")

	} else {
		fmt.Println("这里是其他")
	}

	// 第二种方式 就是 可以直接在 if 里面声明变量
	if ifTwo := 11; ifTwo > 10 {
		fmt.Println("这里是大于十")
	} else {
		fmt.Println("这里是小于十的")
	}

	/*
		goto 跳转
	*/
	goto Here

Here:
	fmt.Println("这里是 goto 过来的")

	/*
		for 循环
	*/
	// 打印 0-10
	for index := 0; index <= 10; index++ {
		fmt.Print(index)
	}
	fmt.Println("")

	// for 配合 range 可以 遍历 map 和 slice 的值
	var mapName = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	for key, value := range mapName {
		fmt.Println(key, value) // a 1 b 2 c 3 d 4
	}

	/*
		switch
		这里的 switch 可以忽略不写
	*/
	var switchOne string = "one"
	switch switchOne {
	case "one":
		fmt.Println("one")
		break
	case "two":
		fmt.Println("two")
		break
	default:
		fmt.Println("two")
		break
	}

	/*
		函数
		格式：
			func funcName (input1 type, input2 type) (output1 type, output2 type) {
				// 这里是 逻辑代码
				return value1,value1
			}
	*/
	// 没有参数 没有返回
	funcOne()

	// 有参数 没有返回
	funcTwo("有参数", "没有返回", 123)

	// 没有参数 有返回
	var funcThree = funcThree()
	fmt.Println(funcThree)

	// 有参数 有返回
	var funcFour = funcFour(1, 2)
	fmt.Println(funcFour)

	// 变参 可以传递多个参数
	funcFIve(1, 2, 3, 213, 67, 2, 34, 52, 324, 6345, 63, 7, 6, 3, 2, 234)


	/*
		函数传递指针
		使用 &x 做实参  使用 *type 来做形参
	*/
	var solid int = 123
	var funcSix = funcSix(&solid)
	fmt.Println(solid)   // 124
	fmt.Println(funcSix) // 124


	/*
		defer 延迟加载
		defer 会在main函数执行完毕时执行 遵循一个 原则 后进先出
	*/
	for i := 0; i < 5; i++ {
		defer fmt.Printf("这块是=defer=过的%d ", i)  // 4 3 2 1 0
	}
	fmt.Println("")


	/*
		函数作为值或者类型
		格式：
			type typeName func (input1 type, input2 type) (result type)
		好处：
			这样就可以吧这个类型的函数当做值来传递
	*/
	// 例子 使用 函数的声明写一个 过滤奇数偶数的 程序
	slice := []int{1,2,3,4,5,6,7,8,9,0}
	odd := filter(slice,isOdd)
	fmt.Println(odd)
	event := filter(slice,isEvent)
	fmt.Println(event)



}

func funcOne() {
	fmt.Println("没有参数 没有返回")
}

func funcTwo(str1, str2 string, int1 int) {
	fmt.Println(str1, str2, int1)
}

func funcThree() (value float64) {
	return 3.131415926
}

func funcFour(a, b int) (int) {
	if a > b {
		return a
	} else {
		return b
	}
}

func funcFIve(args ...int) {
	fmt.Println(args)
}

func funcSix(solid *int) int {
	*solid++
	return *solid
}

/*
	例子 使用 函数的声明写一个 过滤奇数偶数的 程序
*/
type testIntType func (int) bool

// 添加同类型的函数
func isOdd (integer int) (bool) {
	if integer % 2 == 0 {
		return false
	} else {
		return true
	}
}

func isEvent (integer int) (bool) {
	if integer % 2 == 0 {
		return true
	} else {
		return false
	}
}

func filter (slice []int, fn testIntType) (result []int) {
	for _,value := range slice{
		if fn(value) {
			result = append(result,value)
		}
	}
	return
}



