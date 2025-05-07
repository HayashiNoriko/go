package main

// 防止内联！
// 如果是内联函数，则 main 调用 foo 将是原地展开，所以 foo_val 1～5 相当于 main作用域的变量
// 即使 foo_val3 发生逃逸，地址与其他也是连续的
//
//go:noinline
func foo(arg_val int) *int {
	var foo_val1 int = 11
	var foo_val2 int = 22
	var foo_val3 int = 33
	var foo_val4 int = 44
	var foo_val5 int = 55

	// 此处循环是书上给的一种防止内联的方法，但实测并不能防止内联
	// 还是需要在函数前面加上 //go:noinline 标记
	for i := 0; i < 5; i++ {
		println(&arg_val, &foo_val1, &foo_val2, &foo_val3, &foo_val4, &foo_val5)
	}

	// 将 foo_val3 返给 main 函数
	return &foo_val3
}

func main2() {
	main_val := foo(666)
	println(*main_val, main_val)
}
