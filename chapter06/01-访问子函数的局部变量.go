package main

// 这在 C/C++ 中是绝对不允许的

func foo1() *int {
	var foo_val int = 666
	return &foo_val
}

func main1() {
	res := foo1()
	println(*res)
}
