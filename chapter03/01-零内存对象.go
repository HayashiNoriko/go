package main

import (
	"fmt"
)

func main() {
	var s1 struct{}
	var s2 [0]int
	var s3 [100]struct{}
	var s4 = make([]struct{}, 100)

	// 以下 4 行都是零内存对象【指针地址相同】
	fmt.Printf("%p\n", &s1)
	fmt.Printf("%p\n", &s2)
	fmt.Printf("%p\n", &s3[66]) // 取任意元素
	fmt.Printf("%p\n", &s4[66]) // 取任意元素

	// 空切片 => 已初始化，但为空。语义上表示“有，但空”，行为更接近常规切片
	// 【有自己独立的指针地址】
	s5 := []int{}
	fmt.Printf("%p\n", &s5)
	fmt.Printf("%t\n", s5 == nil) // false

	// nil 切片 => 未初始化，底层指针为 0。语义上类似“无”的，但可安全调用部分操作（如 append）
	// 【有自己独立的指针地址】
	// append 后，就不为 nil 了，它被升级为一个真正的切片
	var s6 []int
	fmt.Printf("%p\n", &s6)
	fmt.Printf("%t\n", s6 == nil) // true
	s6 = append(s6, 1)
	fmt.Printf("%t\n", s6 == nil) // false

	// 空接口/任意类型
	// 【有自己独立的指针地址】
	var s7 any
	fmt.Printf("%p\n", &s7)
	fmt.Printf("%t\n", s6 == nil) // false

}
