package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("defer 1")
		func() {
			defer fmt.Println("defer 2")
			runtime.Goexit()
			// return
		}()
		fmt.Println("外层打印")
	}()

	time.Sleep(time.Second)
}
