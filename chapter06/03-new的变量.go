package main

//go:noinline
func foo_new(arg_val int) *int {
	// 和原 foo 函数不同的是，现在这些变量都是用 new 创建的了
	var foo_val1 *int = new(int)
	var foo_val2 *int = new(int)
	var foo_val3 *int = new(int)
	var foo_val4 *int = new(int)
	var foo_val5 *int = new(int)

	for i := 0; i < 5; i++ {
		println(arg_val, foo_val1, foo_val2, foo_val3, foo_val4, foo_val5)
	}

	// 将 foo_val3 返给 main 函数
	return foo_val3
}

func main3() {
	main_val := foo_new(666)
	println(*main_val, main_val)
}
