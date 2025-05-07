package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer func called")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called")
	return 0
}

func twoFunc() int {
	defer deferFunc()
	return returnFunc()
}

func main2() {
	twoFunc()
	/*
		输出：
		return func called
		defer func called
	*/
}
