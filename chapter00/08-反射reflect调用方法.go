package main

import (
	"fmt"
	"reflect"
)

type User2 struct {
	Id   int
	Name string
	Age  int
}

func (u User2) ReflectCallFuncHasArgs(name string, age int) {
	fmt.Println("ReflectCallFuncHasArgs is called. name: ", name, ", age:", age, "and origal User.Name:", u.Name)
}

func (u User2) ReflectCallFuncNoArgs() {
	fmt.Println("ReflectCallFuncNoArgs is called.")
}

// 如何通过反射来进行方法的调用？
// 首先要将方法注册，也就是MethodByName，然后通过反射调动 m.Call

func main() {
	unknown := User2{1001, "Tom", 23}

	// 假设unknown 是外部接受的一个未知类型的对象
	// 1. 获取反射值
	val := reflect.ValueOf(unknown)

	// 2. 获取类型信息
	typ := val.Type()
	fmt.Printf("对象类型：%s\n", typ.Name())

	// 3. 遍历所有方法
	for i := 0; i < val.NumMethod(); i++ {
		method := typ.Method(i)
		fmt.Printf("第%d个方法：%s\n", i, method.Name)

		// 根据方法参数数量决定如何调用
		m := val.Method(i)
		if m.Type().NumIn() == 0 {
			// 无参数方法
			m.Call([]reflect.Value{})
		} else {
			// 有参数方法
			// 这里简化处理，实际应该根据参数类型动态构造参数
			m.Call([]reflect.Value{reflect.ValueOf("Tina"), reflect.ValueOf(23)})
		}

	}
}
