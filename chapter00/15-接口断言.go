package main

import "fmt"

type IBook interface {
	GetName() string
}

type Book struct {
	Name string
}

func (b *Book) GetName() string {
	return b.Name
}

func main15() {
	var mybook IBook = &Book{Name: "三国演义"}

	// 两种获取书名的办法

	// 1. 使用接口给我们提供的 GetName 方法
	name := mybook.GetName()
	fmt.Println(name)

	// 2. 使用接口断言
	if mb, ok := mybook.(*Book); ok {
		fmt.Println(mb.Name)
	}
}
