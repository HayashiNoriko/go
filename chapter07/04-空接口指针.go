package main

type S struct{}

func f(x interface{}) {}

func g(x *interface{}) {}

func main4() {
	s := S{}
	p := &s

	f(s)
	// g(s) // 错误
	f(p)
	// g(p) // 错误
}
