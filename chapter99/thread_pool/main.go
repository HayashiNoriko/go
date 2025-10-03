package main

// goroutine 协程池/线程池使用 --- 无返回值版本

import (
	"fmt"
	"time"
)

// 这是一个需要参数的函数
func myTask(id int, duration time.Duration) {
	fmt.Printf("Task %d starting, will take %s\n", id, duration)
	time.Sleep(duration)
	fmt.Printf("Task %d finished\n", id)
}

func main() {
	// 创建一个包含3个工作协程、任务队列大小为10的协程池
	pool := NewPool(3, 10)

	// 启动协程池
	pool.start()

	// 提交10个任务
	for i := 1; i <= 10; i++ {
		// 为了在循环中正确地捕获变量 i，我们需要将它作为参数传递给闭包
		taskID := i
		// 使用闭包来封装需要参数的函数调用
		pool.Submit(func() {
			myTask(taskID, time.Second*1)
		})
	}

	fmt.Println("All tasks submitted.")

	// 等待所有任务完成后，关闭协程池
	// 在实际应用中，你可能在程序的其他地方（如收到退出信号时）调用 Shutdown
	// 这里为了演示，我们在提交完任务后等待一会再关闭
	time.Sleep(5 * time.Second) // 等待一部分任务完成
	pool.Shutdown()

	// 尝试在关闭后提交任务会导致 panic，因为 channel 已经关闭
	// pool.Submit(func() { fmt.Println("This will panic!") })
}
