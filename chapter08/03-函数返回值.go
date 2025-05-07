package main

import "fmt"

func myfun() (t int) {
	fmt.Println("t =", t)
	defer func() {
		t = 999
	}()
	return 666
}

func main3() {
	res := myfun()
	fmt.Println("res =", res)
	/*
		输出：
		t = 0
		res = 999
	*/
}
