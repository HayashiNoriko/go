// 1、一个类型可以同时满足多个接口
// 2、接口可以组合其他接口（就像 struct 的组合那样，而一个类只要实现了一个接口内部的子接口和方法，那么这个类也就实现了这个接口）例略
// 3、可以轻松地创建适配器模式
// 4、空接口
package main

import (
	"fmt"
)

func main4() {
	// 1、一个类型可以同时满足多个接口
	var m1d Man1 = Doctor{}
	m1d.work()

	var m2d Man2 = Doctor{}
	m2d.work()

	var m1t Man1 = Teacher{}
	m1t.work()

	var m2t Man2 = Teacher{}
	m2t.work()

	// 3、可以轻松地创建适配器模式
	oldPrinter := &MyOldPrinter{}
	adapter := &PrinterAdapter{
		oldPrinter: oldPrinter,
		msg:        "Hello world",
	}
	// 现在可以用 NewPrinter 调用旧系统的功能
	var newPrinter NewPrinter = adapter
	fmt.Println(newPrinter.PrintStored())

	// 4、空接口
	var anything1 interface{} // 传统写法
	var anything2 any         // go 1.18 引入的别名（推荐）
	// 可以存储任意类型的值
	anything1 = 42
	anything2 = "hello"
	anything2 = struct{ Name string }{"Alice"}
	anything2 = []int{1, 2, 3}
	// 作为函数参数
	PrintAnything(anything1)
	PrintAnything(anything2)
	PrintAnything("anything...")
	// 空接口无法直接使用，必须通过类型断言或类型判断来恢复原始类型
	// anything1+=10 // 错误，不能直接操作
	if num, ok := anything1.(int); ok {
		// num 是 anything1 容器的值的副本，这里被提取出来，存储到 num 变量中
		// 需要重新给 anything1 赋值，才能真正修改 anything1 容器内部的值
		anything1 = num + 10
	}
	fmt.Println(anything1)
	var anything3 any = "hellohello"
	str := anything3.(string) // 类型断言，如果失败会 panic
	fmt.Println(str)
	if num, ok := anything3.(int); ok {
		fmt.Println("anything3 的类型是数字，为", num)
	} else {
		fmt.Println("anything3 的类型不是数字")
	}
}

// 1、一个类型可以同时满足多个接口

type Man1 interface {
	work()
}

type Man2 interface {
	work()
}

type Doctor struct{}

func (d Doctor) work() {
	fmt.Println("医生在做手术")
}

type Teacher struct{}

func (t Teacher) work() {
	fmt.Println("老师在讲课")
}

// 3、可以轻松地创建适配器模式
// 假设有一个旧系统的 OldPrinter 接口
type OldPrinter interface {
	Print(s string) string
}

// 旧系统的实现
type MyOldPrinter struct{}

func (p *MyOldPrinter) Print(s string) string {
	return fmt.Sprintf("Old: %s", s)
}

// 现在有一个新系统，它期望使用 NewPrinter 接口
type NewPrinter interface {
	PrintStored() string
}

// 为了让旧系统兼容新系统，我们可以创建一个适配器
// 让 OldPrinter 适配 NewPrinter
type PrinterAdapter struct {
	oldPrinter OldPrinter
	msg        string
}

func (p *PrinterAdapter) PrintStored() string {
	if p.oldPrinter != nil {
		return p.oldPrinter.Print(p.msg)
	}
	return p.msg
}

// 4、空接口
// 作为函数参数
func PrintAnything(v any) {
	fmt.Printf("值：%v，类型：%T\n", v, v)
}
