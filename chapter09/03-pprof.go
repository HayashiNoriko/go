package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func main3() {
	// 启动 loop
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:10000", nil))
	}()

	log.Println(" ===>[Start]")

	readMemStats()
	test2()
	readMemStats()

	log.Println(" ===>[force gc]")
	runtime.GC()

	log.Println(" ===>[Done]")
	readMemStats()

	go func() {
		for {
			readMemStats()
			time.Sleep(time.Second * 10)
		}
	}()

	time.Sleep(3600 * time.Second)
}
