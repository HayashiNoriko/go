package main

import "fmt"

type People interface {
	Say(string) string
}

type Student struct{}

func (s *Student) Say(think string) (talk string) {
	if think == "OK" {
		talk = "Hi"
	} else {
		talk = "Bye"
	}
	return
}

func main2() {
	var p People = &Student{} // 复习一下，Say 使用指针接收者，所以这里必须使用指针。这是大多数的业务场景

	think := "OK"
	fmt.Println(p.Say(think))

	think = "Love"
	fmt.Println(p.Say(think))
}
