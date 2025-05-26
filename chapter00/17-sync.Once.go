package main

import (
	"fmt"
	"sync"
)

// 1. 基本使用
var once sync.Once

func setup() {
	fmt.Println("Initialization")
}

func main17() {
	// 多次调用 Do，但 setup 只会执行一次
	once.Do(setup)
	once.Do(setup)
	once.Do(setup)
}

// 2. 应用：单例模式
type singleton struct{}

var (
	instance      *singleton
	singletonOnce sync.Once
)

func GetInstance() *singleton {
	singletonOnce.Do(func() {
		instance = &singleton{}
	})
	return instance
}
