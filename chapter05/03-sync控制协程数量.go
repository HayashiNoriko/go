package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

var wg3 = sync.WaitGroup{}

func busi3(i int) {
	fmt.Println("go func", i, "goroutine count=", runtime.NumGoroutine())
	wg3.Done()
}

func main3() {
	// 模拟用户需求业务的数量
	task_cnt := math.MaxInt64

	for i := 0; i < task_cnt; i++ {
		wg3.Add(1)
		go busi3(i)
	}

	wg3.Wait()
}
