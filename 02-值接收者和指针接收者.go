// 普通情况下，值接收者和指针接受者的灵活转换
package main

import (
	"fmt"
)

func main2() {
	s := Sheep{name: "喜羊羊"}
	// s := &Sheep{name: "喜羊羊"} // 这两种没有区别
	s.Bark()
	s.Rename("懒羊羊")
	s.Bark()
}

type Sheep struct {
	name string
}

// 值接收者 ===> 语义是 const 的，不修改 s 内部（name 等）
func (s Sheep) Bark() {
	fmt.Printf("%s 咩咩咩\n", s.name)
}

// 指针接收者 ===> 语义是 mutable 的，可以修改 s 内部
func (s *Sheep) Rename(newName string) {
	s.name = newName
}

// 至于 s:= 后面加不加 &？目前看来都是一样的
// 加 &，意思是指针，但依然可以调用值接收者的函数，会灵活转换
// 不加 &，意思是值，但依然可以调用指针接收者的函数（且可以修改内部），也会灵活转换
