package main

// goroutine 协程池/线程池定义 --- 有返回值版本

import (
	"fmt"
	"sync"
)

// Result 封装了任务的执行结果和可能发生的错误
type Result struct {
	Value interface{} // 任务的返回值
	Err   error       // 任务执行过程中遇到的错误
}

// Job 封装了要执行的任务以及其结果通道
type Job struct {
	task       func() (interface{}, error) // 要执行的任务，有返回值和错误
	resultChan chan Result                 // 用于接收结果的专属通道
}

// Pool 是我们的协程池结构体
type Pool struct {
	Jobs    chan Job       // 任务通道，现在传递的是 Job 结构体
	wg      sync.WaitGroup // 用于在关闭时等待所有协程完成任务
	workers int            // 工作协程的数量
}

// NewPool 创建一个新的协程池
func NewPool(workers int, queueSize int) *Pool {
	if workers <= 0 {
		workers = 1
	}
	if queueSize < 0 {
		queueSize = 0
	}

	p := &Pool{
		// 任务通道现在持有 Job 类型
		Jobs:    make(chan Job, queueSize),
		workers: workers,
	}

	return p
}

// Start 启动协程池的所有工作协程
func (p *Pool) Start() {
	for i := 0; i < p.workers; i++ {
		p.wg.Add(1)
		go p.worker(i + 1)
	}
}

// worker 是实际执行任务的工作协程
func (p *Pool) worker(id int) {
	defer p.wg.Done()
	fmt.Printf("Worker %d starting\n", id)

	// 从 Jobs channel 中接收 Job
	for job := range p.Jobs {
		// 执行任务并获取结果
		value, err := job.task()

		// 将结果发送到 Job 专属的结果通道中
		job.resultChan <- Result{Value: value, Err: err}

		// 关闭结果通道，向接收者表明结果已经发送完毕
		close(job.resultChan)
	}

	fmt.Printf("Worker %d stopping\n", id)
}

// Submit 允许用户向协程池提交任务，并返回一个用于接收结果的通道
func (p *Pool) Submit(task func() (interface{}, error)) <-chan Result {
	// 创建一个 Job
	job := Job{
		task: task,
		// 创建一个带缓冲的结果通道。缓冲大小为1，
		// 这样 worker 在发送结果时，即使接收者还没准备好，也不会被阻塞。
		resultChan: make(chan Result, 1),
	}

	// 将 Job 发送到公共任务队列
	// 如果队列满了，这里会阻塞
	p.Jobs <- job

	// 将结果通道返回给调用者
	return job.resultChan
}

// Shutdown 优雅地关闭协程池
func (p *Pool) Shutdown() {
	fmt.Println("Shutting down the pool...")
	close(p.Jobs) // 关闭任务通道
	p.wg.Wait()   // 等待所有 worker 退出
	fmt.Println("Pool shut down.")
}
