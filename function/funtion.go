package main

import "fmt"

// 函数的可变参数，函数的多值返回

func AnyLength(args ...int) {
	for _, n := range args {
		fmt.Println(n)
	}
}

func Hello(to string) (int, error) {
	if to == "" {
		return 0, fmt.Errorf("no name input")
	}

	return fmt.Println("Hello, to", to)

}

func Hello2(to string) (len int, err error) {
	if to == "" {
		//	return 0, fmt.Errorf("no name input")
		len = 0
		err = fmt.Errorf("no name input")
		return
	}

	len, err = fmt.Println("Hello, to", to)
	return
}

func main() {

	AnyLength(1, 2, 3)
	AnyLength()
	var args = []int{1, 2, 3, 4, 5}
	//AnyLength(args) //错误
	AnyLength(args...) // args... 表示将slice里的内容一个一个取出
	len, ok := Hello("")
	if ok == nil {
		fmt.Printf("say %d chars\n", len)
	} else {
		fmt.Println(ok)
	}
	len, ok = Hello("Jok")
	if ok == nil {
		fmt.Printf("say %d chars\n", len)
	} else {
		fmt.Println(ok)
	}

}
