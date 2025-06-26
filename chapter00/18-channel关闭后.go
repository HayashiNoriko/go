// channel 关闭后:
// 接收方可以继续接收，不会报错，只会一直收到零值；可以用 ok 变量来判断是否关闭
// 发送方不可以继续发送数据，会 panic 的
package main

import "time"

func main18() {
	ch := make(chan int, 10)

	go func() {
		// 1. 发送方关闭 channel
		time.Sleep(1 * time.Second)
		close(ch)
		// ch <- 666 // 会 panic
	}()

	// 2. 接收方1
	go func() {
		println("receiver 1: waiting for close")
		num, ok := <-ch
		println("receiver 1: channel is closed, num is ", num, ", ok is", ok)
		num, ok = <-ch
		println("receiver 1: receive again, num is ", num, ", ok is", ok)
	}()

	// 3. 接收方 2
	go func() {
		println("receiver 2: waiting for close")
		num, ok := <-ch
		println("receiver 2: channel is closed, num is ", num, ", ok is", ok)
		num, ok = <-ch
		println("receiver 2: receive again, num is ", num, ", ok is", ok)
	}()

	time.Sleep(5 * time.Second)

	// 子 goroutine 都退出的情况下，主 goroutine 不可再使用 select{}
	// 否则 Go 会认为程序死锁，报 fatal error
	// select {}
}
