package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

var wg4 = sync.WaitGroup{}

func busi4(ch chan bool, i int) {
	fmt.Println("go func", i, "goroutine count=", runtime.NumGoroutine())
	<-ch
	wg4.Done()
}

func main4() {
	// 模拟用户需求业务的数量
	task_cnt := math.MaxInt64

	ch := make(chan bool, 3)

	for i := 0; i < task_cnt; i++ {
		wg4.Add(1)
		ch <- true
		go busi4(ch, i)
	}

	wg4.Wait()
}
