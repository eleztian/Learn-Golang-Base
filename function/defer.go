package main

import (
	"fmt"
	"os"
)

// 延迟函数 延迟函数不会立即执行，而是在return / panic后按照先进后出(FILO)依次执行。
// 因此它的实现方式其实就是一个栈
// defer 可以用于 观察返回值， 成对操作的后续处理比如关闭文件等

// 可以得到最后返回的是1而不是0
func fun() (ret int) {
	defer func() {
		ret++
	}()
	return 0
}

// 将slice反序输出
func reverse(slice []int) {
	for i := range slice {
		defer func(i int) { // 匿名函数， 也称闭包
			fmt.Printf("%d ", i)
		}(i)
	}
	fmt.Println("reverse slice :")
}

// 输出结果 全部为第一个。为什么呢？
// 可以看匿名函数的实现得到原因
// 对于defer这行的理解是在声明一个函数后并得到了一个函数值，这个函数值包含了当前的环境变量地址，
// 即是slice,i的地址,也就是说他们的环境变量是一样的
// 而当rang执行到最后时&i这个地址下的值已经变成了最后一次的值。所以在函数执行完执行defer栈中的函数是i的值是一样的
// 都为最后一次的值
func reverse2(slice []int) {
	for i := range slice {
		defer func() { // 匿名函数
			fmt.Printf("%d ", slice[i])
		}() // 括号不能少
	}
	fmt.Println("reverse slice :")
}

// 用于关闭文件
func ReadWrite() bool {
	f, ok := os.Open("testfile")
	if ok != nil {
		return false
	}
	defer f.Close()
	// ...
	return true

}

func main() {
	fmt.Println(fun()) // 1
	slice := []int{1, 2, 3, 4, 5}
	reverse(slice) // 5 4 3 2 1
	fmt.Println("")
	reverse2(slice) // 4 4 4 4 4
}
