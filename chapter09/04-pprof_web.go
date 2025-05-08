package main

import (
	"bytes"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func init() {
	rand.Seed(1) // 固定种子，禁用自动种子生成
}

func test4() {
	for {
		println(genSomeBytes())
	}
}

func genSomeBytes() *bytes.Buffer {
	buff := make([]byte, 20000)
	for i := range buff {
		buff[i] = byte(rand.Intn(10)) + '0' // 使用 math/rand
	}
	return bytes.NewBuffer(buff)
}

func main() {
	// 禁用 HTTP 服务，仅测试随机数生成
	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		go test4()
	}
	http.ListenAndServe("0.0.0.0:10000", nil)
	time.Sleep(time.Second * 3600)
}
