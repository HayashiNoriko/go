package main

import "fmt"

func main6() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常：", err)
		} else {
			fmt.Println("没有异常")
		}
	}()

	defer func() {
		panic("defer 中的 panic")
	}()

	panic("main 中的 panic")
}
