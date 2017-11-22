package main

import "fmt"

// panic(恐慌) and recover 异常和恢复 Panic抛出异常,Recover捕获异常
// Panic 内建函数，中断原有控制流程，不会影响defer中函数列表的正常执行, 恐慌可以由Panic产生也可以由运行错误产生。
// Recover 内建函数，他可以铺货到panic的输入值，使得处于panic中goroutine恢复过来, **recover仅仅在defer中有效**，
// 在正常执行过程中recover只会返回nil.
//
// 由于panic会引发程序的崩溃，因此panic一般用于严重错误，在程序应当尽量避免使用panic
// 使用recover可以帮助我们在程序崩溃前做一些必要的处理，典型的例子是web服务器, web服务器遇到Panic会调用Recover，
// 输出堆栈信息并继续运行。
// 但是使用Recover无法保证包级变量的状态任然和预期一致，因此我们不应该去Recover其他包的panic，而是应该使用error
// 返回错误。

// 通常情况下我们会将Panic Value的一些类型进行标识，当调用recover时检查它并判断是否恢复，以此来恢复一些应该恢复的
// panic。

func throwPanicChoose(f func(int), arg int) (b bool) {
	defer func() {
		switch x := recover(); x {
		case nil:
		case "index out":
			b = true
			fmt.Println("index out")
		case "index error":
			b = true
			fmt.Println("index error")
		default:
			panic(x)
		}
	}()
	f(arg)
	return
}

func throwPanic(f func(int), arg int) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
			b = true
		}
	}()
	f(arg)
	return
}

func err_panic(index int) {
	var slice = []int{1, 2, 3, 4}
	fmt.Println(slice[index])
}

func panic_panic(index int) {
	var slice = []int{1, 2, 3, 4}
	if index > len(slice) {
		panic("index out")
	}
	if index < 0 {
		panic("index error")
	}
}

// 没有return语句但能返回一个非0值的函数
func noReturn() (result int) {
	defer func() {
		p := recover()
		result = p.(int)
	}()
	panic(23)
}

func main() {
	fmt.Println(throwPanic(err_panic, 5))   // true
	fmt.Println(throwPanic(err_panic, 3))   // 4 false
	fmt.Println(throwPanic(panic_panic, 5)) // true
	fmt.Println(throwPanic(panic_panic, 3)) // false

	fmt.Println(throwPanicChoose(panic_panic, 5))  // true
	fmt.Println(throwPanicChoose(panic_panic, 3))  // fale
	fmt.Println(throwPanicChoose(panic_panic, -1)) // fale

	fmt.Println(noReturn())
}
