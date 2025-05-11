package main

import (
	"fmt"
	"reflect"
)

type Duck struct {
	Name string
	Age  int
}

func main9() {
	obj := Duck{"Jenny", 10}

	typ := reflect.TypeOf(obj)
	field, _ := typ.FieldByName("Name")
	fmt.Println(field) // {Name  string  0 [0] false}

	val := reflect.ValueOf(obj)
	fieldValue := val.FieldByName("Name")
	fmt.Println(fieldValue) // Jenny
}
