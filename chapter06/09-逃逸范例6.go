package main

import (
	"fmt"
)

func foo9(a *int) {
	return
}

func main9() {
	data := 10
	f := foo9

	f(&data)

	fmt.Println(data)
	// println(data)
}
