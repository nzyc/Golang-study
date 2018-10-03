package main

import (
	"fmt"
)

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

	// 变参 可以传递多个参数  在函数体中 args 是一个接受int 的 slice
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
		defer fmt.Printf("这块是=defer=过的%d ", i) // 4 3 2 1 0
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
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	odd := filter(slice, isOdd)
	fmt.Println(odd)
	event := filter(slice, isEvent)
	fmt.Println(event)

	/*
		main init 函数
		Go在定义这俩个函数的时候 这俩个函数没有任何参数 和返回值
		每个 package 里面可以有多个 init 函数 但是 为了容易维护 还是一个函数里面有一个init函数

		Go会自动调取 main 和 init 这俩个函数
		package 里面的init 都是可选的 但是 必须要有main 函数
	*/

	/*
		import
		导入包
	*/
	// import "fmt"

	// 相对路径 导入包
	// import "./module"	// 与当前文件 同一目录下的module 包

	// 导入绝对路径
	// import "shorturl/model"	// 加载  src_1 下面的 shorturl/model  src_1/shorturl/model

	/*
		点操作 导入包
		import (
			. "fmt"
		)
		这样我们 就可以不用
		fmt.Println()
		直接写成
		Println()
	*/

	/*
		别名操作 导入包
		import (
			f "fmt"
		)
		这样我们 就可以不用
		fmt.Println()
		直接写成
		f.Println()
	*/

	/*
		_ 操作 导入包
		import (
			_ "fmt"
		)
		这个 导入包的意思是  不使用包内 的 函数 而是使用 包里的 init函数

	*/

	/*
		恐慌(Panic)和恢复(Recover)
	*/
	// throwsPanic()

	/*
		小练习
	*/
	// 1
	floSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9}
	meanValue := average(floSlice)
	fmt.Println(meanValue)

	// 2
	aaa, bbb := compareSize(1, 2)
	fmt.Println(aaa, bbb)

	// 3
	var sss stack
	sss.push(25)
	sss.push(22)
	sss.push(21)
	fmt.Println(sss)
	fmt.Printf("stack %v\n", sss)
	sss.pop()
	fmt.Printf("stack %v\n", sss)

	// 4
	for _, term := range feibonaqi(10) {
		fmt.Printf("%v ", term)
	}

	fmt.Println("")
	// 5
	p := mkas()
	fmt.Println(p(2))

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
type testIntType func(int) bool

// 添加同类型的函数
func isOdd(integer int) (bool) {
	if integer%2 == 0 {
		return false
	} else {
		return true
	}
}

func isEvent(integer int) (bool) {
	if integer%2 == 0 {
		return true
	} else {
		return false
	}
}

func filter(slice []int, fn testIntType) (result []int) {
	for _, value := range slice {
		if fn(value) {
			result = append(result, value)
		}
	}
	return
}
func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f()
	return
}

/*
	小练习
*/

// 1、编写一个函数用于计算一个float64类型的slice的平均值。
func average(flo []float64) (float64) {
	sum := 0.0
	for _, value := range flo {
		sum += value
	}
	ave := sum / float64(len(flo))
	return ave
}

// 2、编写函数，返回其(两个)参数正确的(自然)数字顺序
// 例子
// f(7,2) → 2,7
// f(2,7) → 2,7
func compareSize(a, b int) (aa, bb int) {
	if a > b {
		aa, bb = b, a
	} else {
		aa, bb = a, b
	}
	return

}

// 3 创建一个固定大小保存整数的栈。它无须超出限制的增长。
// 定义push函数—— 将数据放入栈，和 pop 函数——从栈中取得内容。
// 栈应当是后进先出(LIFO) 的。
type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	s.data[s.i] = k
	s.i++
}
func (s *stack) pop() {
	s.i--
	//return s.data[s.i]

}

// 4、斐波那契
func feibonaqi(value int) []int {
	x := make([]int, value)
	x[0],x[1]=1,1
	for i := 2; i < value; i++ {
		x[i] = x[i-1] + x[i-2]
	}
	return x
}

// 5  编写一个函数返回另一个函数，返回的函数的作用是对一个整数 +2。
// 函数的名称叫做 plusTwo。然后可以像下面这样使用:
// p := plusTwo()
// fmt.Printf("%v\n", p(2))
// 应该打印 4。

type huidiao func(int) int 
func mkas () huidiao {
	return func(i int) int {
		return i + 2
	}
}


/*
	快速排序
*/
func qsort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	qsort(data[:head])
	qsort(data[head+1:])
}