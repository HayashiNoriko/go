package main

import (
	"fmt"
	"sync"
	"time"
)

type httpPkg struct{}

func (httpPkg) Get(url string) {
	// 模拟网络请求
	time.Sleep(1 * time.Second)
	fmt.Println("Get", url)
}

var hp httpPkg

func main16() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.example.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			hp.Get(url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()

	fmt.Println("Finished all HTTP requests.")
}
