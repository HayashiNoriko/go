package main

import (
	"encoding/json"
	"fmt"
)

type Man struct {
	Name     string
	Age      int
	Province string
}

func main11() {
	// 1.反序列化（json => 结构体）
	// json字符串
	jsonStr := `{"name":"zhangsan","age":18,"province":"山东"}`
	// 声明结构体变量
	var man Man

	err := json.Unmarshal([]byte(jsonStr), &man)
	if err != nil {
		// 处理错误
		fmt.Println("err = ", err)
	}
	fmt.Println("man = ", man) // man =  {zhangsan 18 山东}

	// 2. 序列化（结构体 => json）
	// 结构体变量
	man2 := Man{Name: "lisi", Age: 20, Province: "北京"}
	// 序列化为 json 字符串
	jsonStr2, err := json.Marshal(man2)
	if err != nil {
		// 处理错误
		fmt.Println("err = ", err)
	}
	fmt.Println("jsonStr2 = ", string(jsonStr2)) // jsonStr2 =  {"Name":"lisi","Age":20,"Province":"北京"}
}
