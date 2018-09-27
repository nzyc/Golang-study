package main

import "fmt"

/*
	struct 声明一个新属相
	来储存 类型 的属相或字段的容器
*/
type person struct {
	name string
	age  int
}


func main() {


	// 简单使用
	var P person
	P.name = "shabi"
	P.age = 20
	fmt.Println(P) // {shabi 20}

	/*
		小案列
		比较俩个人的年龄 返回年龄大的 并返回这俩个人的年龄差
	*/
	// 第一种定义的方式
	var structOne person
	structOne.name = "Andy"
	structOne.age = 30

	// 第二种定义的方式 (顺序可以颠倒)
	structTwo := person{name: "beat", age: 29}

	// 第三种方式
	structThree := person{"miya", 24}

	// 比大小
	obj1, num1:= older(structOne, structTwo)
	obj2, num2:= older(structThree, structTwo)
	fmt.Println(obj1,num1)
	fmt.Println(obj2,num2)



	/*
		struct 匿名字段
	*/



}

func older(one, two person) (person,int) {
	if one.age > two.age {
		var diffNum = one.age - two.age
		return one, diffNum
	} else if one.age < two.age  {
		var diffNum = two.age - one.age
		return two, diffNum
	} else {
		return one,123
	}
}
