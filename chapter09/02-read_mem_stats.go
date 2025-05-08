package main

import (
	"log"
	"runtime"
	"time"
)

func readMemStats() {
	var ms runtime.MemStats

	runtime.ReadMemStats(&ms)

	log.Printf(" ===> Alloc: %d(Bytes) HeadIdle:%d(Bytes) HeadReleased:%d(Bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

func test2() {
	// slice 会动态扩容，用 slice 来做堆内存申请
	container := make([]int, 8)

	log.Println(" ===> loop begin.")

	for i := 0; i < 32000000; i++ {
		container = append(container, i)
		if i == 16000000 {
			readMemStats()
		}
	}

	log.Println(" ===> loop end.")
}

func main2() {
	log.Println(" ===> [Start]")

	readMemStats()
	test2()
	readMemStats()

	log.Println(" ===> [force gc]")

	// 强制调用 gc 回收
	runtime.GC()
	log.Println(" ===> [Done]")

	readMemStats()

	go func() {
		for {
			readMemStats()
			time.Sleep(time.Second * 10)
		}
	}()

	// 睡眠，保持程序不退出
	time.Sleep(time.Second * 3600)
}
