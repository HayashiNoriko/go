package main

import "fmt"

func main1() {
	var p *int

	p = new(int) // 必须分配内存，否则报错

	fmt.Println(*p) // 0

	*p = 10
	fmt.Println(*p) // 10
}
