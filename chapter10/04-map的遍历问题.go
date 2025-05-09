package main

import "fmt"

type Teacher struct {
	Name string
	Age  int
}

func main4() {
	m := make(map[string]*Teacher)

	// 创建一个 Teacher 数组
	teachers := []Teacher{
		{Name: "张老师", Age: 20},
		{Name: "李老师", Age: 22},
		{Name: "王老师", Age: 24},
	}

	// 将数组依次添加到 map 中
	for _, t := range teachers {
		m[t.Name] = &t
	}

	// 打印 map
	for k, v := range m {
		fmt.Println(k, "=>", v)
	}
}
