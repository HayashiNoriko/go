package main

import (
	"fmt"
	// "time"
)

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			fmt.Println("c <- x", x)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
			// default:
			// fmt.Println("default")
		}
		// fmt.Println("fibonacci return")
		// return
	}
}

func main13() {
	c := make(chan int)
	quit := make(chan int)

	// 需要确保先执行<-c，打开接收端，c <- x才不会阻塞
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	// time.Sleep(time.Second)

	fibonacci(c, quit)

}
