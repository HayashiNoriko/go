package main

import (
	"log"
	"runtime"
	"time"
)

func test() {
	// slice 会动态扩容，用 slice 来做堆内存申请
	container := make([]int, 8)

	log.Println(" ===> loop begin.")

	for i := 0; i < 32000000; i++ {
		container = append(container, i)
	}

	log.Println(" ===> loop end.")
}

func main1() {
	// time.Sleep(time.Second * 10)
	log.Println("Start.")
	test()
	log.Println("Force gc.")

	// 强制调用GC 回收
	runtime.GC()

	log.Println("Done.")

	// 睡眠，保证程序不退出
	time.Sleep(time.Second * 3600)
}
