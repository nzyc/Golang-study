package main

import "fmt"

func main() {
	/*
		Go语言中有25个关键词
		package 	import 		func 		var 	const
		if 			else 		switch 		case 	break
		continue 	default 	for 		range 	return
		type 		struct 		interface 	map 	defer
		chan 		fallthrough select 		go 		goto

		内建常量
		true		false		iota		nil

		内建类型
		int		int8		int16		int32		int64
		uint	uint8		uint16		uint32		uint64
		float32	float64		complex64	complex128	error
		byte	rune		string		bool		uintptr

		内建函数
		close					用于channel通讯,使用它来关闭channel
		delete					用于删除map中的实例
		len和cap					可用于不同类型 len 返回长度 cap 返回详细信息
		new						用于不同类型的内存分配
		make					用于内建模型 map slice channel 的内存分配
		copy					用于复制slice
		panic 和 recover			用于异常处理机制
		print 和 println			打印函数，用于调试
		complex、real 和 imag	全部用于处理 复数
	*/

	/*
		go语言推荐使用 驼峰式命名
	*/
	fmt.Println("hello world")
}
