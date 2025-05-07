package main

import "fmt"

func main4() {
	defer_call()
	fmt.Println("main 正常结束")
}

func defer_call() {
	defer func() {
		fmt.Println("defer panic之前，1")
	}()
	defer func() {
		fmt.Println("defer panic之前，2")
	}()

	panic("突然 panic") // 触发 defer出栈

	defer func() {
		fmt.Println("defer panic之后，永远不会执行")
	}()
}
