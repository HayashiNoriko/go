package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc() {
	fmt.Println("ReflectCallFunc")
}

func main7() {
	user := User{
		Id:   1,
		Name: "Alice",
		Age:  18,
	}
	DoFieldAndMethod(user)
}

// 通过接口来获取任意参数，然后一一揭晓
func DoFieldAndMethod(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println("get type is:", getType.Name()) // 不加 .Name() 就是 main.User

	getValue := reflect.ValueOf(input)
	fmt.Println("get value is:", getValue) // 字段值的列表

	// 获取字段
	// 已有 reflect.Type，通过 NumField(字段数) 进行遍历
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)              // 获取字段（名、类型）
		value := getValue.Field(i).Interface() // 获取字段值
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 已有 reflect.Type，然后通过 NumMethod(方法数) 进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
