package main

import (
	"fmt"
	"reflect"
)

func main6() {
	var num float64 = 1.2345

	// 1. reflect.ValueOf() 和 reflect.TypeOf() 的使用
	fmt.Println("Type:", reflect.TypeOf(num))   // Type: float64
	fmt.Println("Value:", reflect.ValueOf(num)) // Value: 1.2345

	// 2. 从 relfect.Value 中获取接口 interface 的信息
	num_pointer_value := reflect.ValueOf(&num)
	num_value := reflect.ValueOf(num)
	// 2.1 已知原类型的情况
	// 先转换为接口类型，再进行类型断言转换为具体类型
	// 类型一定要完全符合，否则会 panic
	convertPointer := num_pointer_value.Interface().(*float64)
	convertValue := num_value.Interface().(float64)
	fmt.Println("convertPointer:", convertPointer) // convertPointer: 0xc000104020
	fmt.Println("convertValue:", convertValue)     // convertValue: 1.2345
	// 2.2. 未知原类型，且原类型为 struct 的情况
	// 见下节

	// 3. 通过 reflect.Value 设置实际变量的值（适用于不知道原类型的情况，因为如果知道原类型，那就用 2.1的方法，类型断言，就可以了）
	// Elem() 的调用者……？

	// 3.1 为值
	fmt.Println("num_value.Kind():", num_value.Kind())     // num_value.Kind(): float64
	fmt.Println("num_value.Type():", num_value.Type())     // num_value.Type(): float64
	fmt.Println("num_value.CanSet():", num_value.CanSet()) // num_value.CanSet(): false
	// elem1 := num_value.Elem() // panic: reflect: call of reflect.Value.Elem on float64 Value

	// 3.2 为指针
	fmt.Println("num_pointer_value.Kind():", num_pointer_value.Kind())     // num_pointer_value.Kind(): ptr
	fmt.Println("num_pointer_value.Type():", num_pointer_value.Type())     // num_pointer_value.Kind(): *float64
	fmt.Println("num_pointer_value.CanSet():", num_pointer_value.CanSet()) // num_pointer_value.Kind(): false
	elem2 := num_pointer_value.Elem()
	fmt.Println("elem2.Kind():", elem2.Kind())     // elem2.Kind(): float64
	fmt.Println("elem2.Type():", elem2.Type())     // elem2.Type(): float64
	fmt.Println("elem2.CanSet():", elem2.CanSet()) // elem2.CanSet(): true
	elem2.SetFloat(5.6789)
	fmt.Println("new value of num:", num) // new value of num: 5.6789

	// 3.3 为接口，接口实际存储值
	var inter1 interface{} = num
	inter1_value := reflect.ValueOf(inter1)
	fmt.Println("inter1_value.Kind():", inter1_value.Kind())     // inter1_value.Kind(): float64
	fmt.Println("inter1_value.Type():", inter1_value.Type())     // inter1_value.Type(): float64
	fmt.Println("inter1_value.CanSet():", inter1_value.CanSet()) // inter1_value.CanSet(): false

	// elem3 := inter1_value.Elem() // panic reflect: call of reflect.Value.Elem on float64 Value

	// 3.4 为接口，接口实际存储指针
	var inter2 interface{} = &num
	inter2_value := reflect.ValueOf(inter2)
	fmt.Println("inter2_value.Kind():", inter2_value.Kind())     // inter1_value.Kind(): ptr
	fmt.Println("inter2_value.Type():", inter2_value.Type())     // inter1_value.Type(): *float64
	fmt.Println("inter2_value.CanSet():", inter2_value.CanSet()) // inter1_value.CanSet(): false
	elem4 := inter2_value.Elem()
	fmt.Println("elem4.Kind():", elem4.Kind())     // inter1_value.Kind(): float64
	fmt.Println("elem4.Type():", elem4.Type())     // inter1_value.Type(): float64
	fmt.Println("elem4.CanSet():", elem4.CanSet()) // inter1_value.CanSet(): true
	elem4.SetFloat(6.7899)
	fmt.Println("new value of num:", num) // new value of num: 6.7899

}
