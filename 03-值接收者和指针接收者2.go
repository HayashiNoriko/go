// 接口情况下，值接收者和指针接受者有一定限制
// 对于包含多个方法的接口，类型必须满足所有方法的接收者要求才能实现该接口
package main

import (
	"fmt"
)

func main() {
	// f := Fish{}   // 值 declared and not used
	fp := &Fish{} // 指针

	// var m Mover = f // m 为值的形式，错误，因为 Fish 没有实现 Move2，实现接口失败
	var m Mover = fp // m 为指针的形式，正确，因为 Fish* 实现了 Move 和 Move2，实现接口成功
	m.Move()
	m.Move2()

}

type Fish struct {
}

type Mover interface {
	Move()
	Move2()
}

// 使 Fish 和 Fish* 都实现 Move
func (f Fish) Move() {
	fmt.Println("实体鱼在移动")
}

// 使只有 Fish* 实现了 Move2
func (f *Fish) Move2() {
	fmt.Println("指针鱼在移动")
}
