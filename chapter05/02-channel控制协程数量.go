package main

import (
	"fmt"
	"math"
	"runtime"
)

func busi2(ch chan bool, i int) {
	fmt.Println("go func", i, "goroutine count=", runtime.NumGoroutine())
	<-ch
}

func main2() {
	// 模拟用户需求业务的数量
	task_cnt := math.MaxInt64
	// task_cnt = 10

	ch := make(chan bool, 3)

	for i := 0; i < task_cnt; i++ {
		ch <- true
		go busi2(ch, i)
	}
}
