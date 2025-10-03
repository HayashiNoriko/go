package main

// goroutine 协程池/线程池使用 --- 有返回值版本

import (
	"errors"
	"fmt"
	"time"
)

// 一个有返回值的任务，模拟计算平方
func square(n int) (interface{}, error) {
	time.Sleep(time.Millisecond * 500) // 模拟耗时
	return n * n, nil
}

// 一个会返回错误的任务
func failingTask() (interface{}, error) {
	time.Sleep(time.Millisecond * 500) // 模拟耗时
	return nil, errors.New("this task is designed to fail")
}

func main() {
	// 创建并启动协程池
	pool := NewPool(4, 20)
	pool.Start()

	// 用于存储所有任务的结果通道
	var resultChannels []<-chan Result

	// 提交10个计算平方的任务
	for i := 1; i <= 10; i++ {
		num := i // 正确捕获循环变量
		fmt.Printf("Submitting task: square(%d)\n", num)
		// 提交任务并保存返回的结果通道
		resultChan := pool.Submit(func() (interface{}, error) {
			return square(num)
		})
		resultChannels = append(resultChannels, resultChan)
	}

	// 提交一个会失败的任务
	fmt.Println("Submitting a failing task")
	resultChannels = append(resultChannels, pool.Submit(failingTask))

	fmt.Println("\n--- All tasks submitted. Waiting for results... ---")

	// 遍历结果通道，获取并打印每个任务的结果
	// 从通道接收结果的操作 (<-resChan) 是阻塞的，
	// 这会一直等到对应的任务完成。
	for i, resChan := range resultChannels {
		// 从结果通道中接收结果
		res := <-resChan
		if res.Err != nil {
			fmt.Printf("Result for task %d: Failed with error -> %v\n", i, res.Err)
		} else {
			// 使用类型断言来获取具体类型的值
			if val, ok := res.Value.(int); ok {
				fmt.Printf("Result for task %d: Success -> %d\n", i, val)
			}
		}
	}

	fmt.Println("\n--- All results received. ---")

	// 关闭协程池
	pool.Shutdown()
}
