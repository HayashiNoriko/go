package main

import "fmt"

type Commen1 struct {
	name string
	age  int
	sex  string
}

type Commen2 struct {
	name string
	age  int
	sex  string
}

type Special struct {
	Commen1 // 这是组合，或者叫继承，不是字段，所以不用指定字段名
	// commen  Commen // 这是字段

	Commen2 // 测试：同时继承两个有相同字段的结构体
	special string
}

func main19() {
	s := Special{
		// 继承的结构体，没有字段名，这里冒号前要写上类名 Commen
		Commen1: Commen1{
			name: "zhangsan",
			age:  18,
			sex:  "男",
		},
		Commen2: Commen2{
			name: "lisi",
			age:  20,
			sex:  "女",
		},
		special: "special zhangsan",
	}

	// 由于字段名冲突，必须明确指定是哪个组合结构体的字段
	fmt.Println("Commen1的字段:")
	fmt.Println(s.Commen1.name)
	fmt.Println(s.Commen1.age)
	fmt.Println(s.Commen1.sex)

	fmt.Println("Commen2的字段:")
	fmt.Println(s.Commen2.name)
	fmt.Println(s.Commen2.age)
	fmt.Println(s.Commen2.sex)

	fmt.Println("Special自己的字段:")
	fmt.Println(s.special)

	// 这样会报错：ambiguous selector
	// fmt.Println(s.name)  // 编译器不知道你要的是Commen1.name还是Commen2.name

	// 但如果只继承 Commen1，那其实可以直接调用 s.name，等价于调用 s.Commen1.name
}
