package main

// goroutine 协程池/线程池定义 --- 无返回值版本

import (
	"fmt"
	"sync"
)

// Task 定义了我们协程池中执行的任务类型，它是一个无参数也无返回值的函数。
// 我们可以通过闭包来传递参数。
type Task func()

// Pool 是我们的协程池结构体
type Pool struct {
	Tasks   chan Task      // 任务通道，用于接收任务
	wg      sync.WaitGroup // 用于在关闭时等待所有协程完成任务
	workers int            // 工作协程的数量
}

// NewPool 创建一个新的协程池
// - workers: 池中工作协程的数量
// - queueSize: 任务队列（channel）的缓冲大小
func NewPool(workers int, queueSize int) *Pool {
	if workers <= 0 {
		workers = 1 // 保证至少有一个 worker
	}
	if queueSize < 0 {
		queueSize = 0 // 无缓冲 channel
	}

	p := &Pool{
		Tasks:   make(chan Task, queueSize),
		workers: workers,
	}

	return p
}

// start 启动协程池的所有工作协程
func (p *Pool) start() {
	// 启动指定数量的 worker 协程
	for i := 0; i < p.workers; i++ {
		// 为每个 worker 协程增加一个计数
		p.wg.Add(1)
		go p.worker(i + 1)
	}
}

// worker 是实际执行任务的工作协程
func (p *Pool) worker(id int) {
	// 在协程退出时，通知 WaitGroup 任务已完成
	defer p.wg.Done()
	fmt.Printf("Worker %d starting\n", id)

	// 使用 for-range 循环不断地从 Tasks channel 中接收任务
	// 当 channel 被关闭且里面的数据都被消费完后，循环会自动退出
	for task := range p.Tasks {
		// 执行任务
		task()
	}

	fmt.Printf("Worker %d stopping\n", id)
}

// Submit 允许用户向协程池提交任务
func (p *Pool) Submit(task Task) {
	// 将任务发送到任务通道
	// 如果通道已满（对于有缓冲的 channel），这里会阻塞
	p.Tasks <- task
}

// Shutdown 优雅地关闭协程池
func (p *Pool) Shutdown() {
	fmt.Println("Shutting down the pool...")

	// 关闭 Tasks channel。这会向所有正在等待的 worker 发出信号，
	// 告诉它们不会再有新的任务了。
	close(p.Tasks)

	// 等待所有 worker 协程执行完毕。
	// 当所有 worker 的 Done() 都被调用后，Wait() 才会返回。
	p.wg.Wait()

	fmt.Println("Pool shut down.")
}
