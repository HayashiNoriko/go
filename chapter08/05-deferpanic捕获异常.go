package main

import "fmt"

func main5() {
	defer_call5()
	fmt.Println("main 正常结束")
}

func defer_call5() {
	defer func() {
		fmt.Println("defer panic之前，1，并捕获异常")
		if err := recover(); err != nil {
			fmt.Println("捕获异常：", err)
		}
	}()

	defer func() {
		fmt.Println("defer panic之前，2，不捕获")
	}()

	panic("突然 panic") // 触发 defer出栈

	defer func() {
		fmt.Println("defer panic之后，永远不会执行")
	}()
}
