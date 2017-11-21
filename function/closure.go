package main

import "fmt"

// 匿名函数(闭包 function literal), 函数式编程
// 匿名函数 携带了定义这个函数的环境，也就是所被定义的这个函数可以完全访问携带的函数环境
// 匿名函数，也就是这个函数值，它其实是一个新对象(包含匿名函数地址和环境变量地址)的地址，定义如下
// type funcValue {
//     f uintptr //函数地址
//     x *int  // 环境
// }

// 这个函数返回一个func，这个func无参数，有一个int返回值
func squares() func() int {
	var x int // 初始值默认为零值，int的零值为0
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares() // 得到一个函数值，类似于C语言的函数指针
	// 这里函数值一旦确定，那么函数的环境变量就是唯一的，因此每次执行这个函数f时，它所操作的x是同一个。
	//f2 := squares()
	//fmt.Println(f == f2) 编译错误，函数值属于引用类型，且函数值不可比较。
	fmt.Println(f()) // 1
	fmt.Println(f()) // 4
	fmt.Println(f()) // 9
	fmt.Println(f()) // 16
}
