package main

import (
	"fmt"
	"sync"
)

type user struct {
	lock sync.Mutex
	name string
	age  int
}

func main2() {
	u := new(user) // 默认给 u 分配到的内存全部置为 0

	fmt.Println(u) // &{{{} {0 0}}  0}

	u.lock.Lock() // 可以直接使用，因为 u.lock 为 0，表示开锁状态

	u.name = "张三"
	fmt.Println(u) // &{{{} {1 0}} 张三 0}

	u.lock.Unlock()

	fmt.Println(u) // &{{{} {0 0}} 张三 0}
}
