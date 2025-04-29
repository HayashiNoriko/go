package main

// go run 01-go\ tool\ trace.go
// go tool trace trace.out

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	// 创建 trace 文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 启动 trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	// main
	fmt.Println("hello world")
}
