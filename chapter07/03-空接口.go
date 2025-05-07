package main

import (
	"fmt"
)

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
	} else {
		fmt.Println("non-empty interface")
	}
}

func main3() {
	var p *int = nil
	Foo(p)
}
