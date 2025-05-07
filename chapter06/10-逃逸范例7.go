package main

import (
	"fmt"
)

func foo10(a []string) {
	return
}

func main10() {
	s := []string{"hello"}
	foo10(s)
	fmt.Println(s)
}
