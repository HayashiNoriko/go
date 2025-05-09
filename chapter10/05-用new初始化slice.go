package main

import "fmt"

func main() {
	list := new([]int)
	// list = append(list, 1) // 报错
	*list = append(*list, 1) // 正确
	fmt.Println(*list)       // [1]
}
