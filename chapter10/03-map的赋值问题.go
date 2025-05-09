package main

import (
	"fmt"
)

type Student struct {
	Name string
}

var list map[string]Student

func main3() {

	list = make(map[string]Student) // 必须使用make进行初始化，否则报错：assignment to entry in nil map

	student := Student{"Tina"}

	list["0001"] = student // 值复制过程，允许
	// list["0001"].Name = "Nico" // 报错，cannot assign to struct field list["0001"].Name in map
	// 原因：list["0001"]是一个值引用，值引用的特点是只读的，不允许修改

	// 解决方法一：重新创建一个 Student（效率低）
	list["0001"] = Student{"Nico"} // 允许，因为是赋值过程

	// 解决方法二：将 map 的 value改为 *Student，即指针
	// 这样 list["0001"] 虽然是一个值引用，但是它指向的是一个 Student 指针
	// 我们不修改指针的值，而是通过指针可以修改其指向的 Student，这是允许的
	// （代码略）

	fmt.Println(list["0001"])
}
