package main

import "fmt"

func main1() {
	// 依次输出 C、B、A
	defer func1()
	defer func2()
	defer func3()
}

func func1() {
	fmt.Println("A")
}
func func2() {
	fmt.Println("B")
}
func func3() {
	fmt.Println("C")
}
