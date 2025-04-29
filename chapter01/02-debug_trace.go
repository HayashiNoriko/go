package main

// go build 02-debug_trace.go // 生成go可执行文件
// GODEBUG=schedtrace=1000 ./02-debug_trace

import (
	"fmt"
	"time"
)

func main2() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello!")
	}
}
