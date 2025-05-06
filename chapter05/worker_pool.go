package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
Worker Manager
*/
type WorkerManager struct {

	// 用来监控 Worker 是否已经死亡的缓冲 channel
	workerChan chan *Worker

	// 一共要监控的 Worker 的数量
	nWorkers int
}

// 创建一个 WorkerManager 对象
func NewWorkerManager(nworkers int) *WorkerManager {
	return &WorkerManager{
		nWorkers:   nworkers,
		workerChan: make(chan *Worker, nworkers),
	}
}

// 启动 Worker 池，并为每个 Worker 分配一个 ID，让每个 Worker 进行工作
func (wm *WorkerManager) StartWorkerPool() {
	// 创建一定数量的 Worker
	for i := 0; i < wm.nWorkers; i++ {
		// 如果不使用 i := i，所有 goroutine 可能会看到相同的最终 i 值（通常是循环结束时的值）
		// 在 Go 1.22 及以上版本中，这种行为已经改变，循环变量在每次迭代中会自动创建新实例，不再需要这种写法
		i := i
		wk := &Worker{id: i}
		go wk.work(wm.workerChan)
	}

	// 启动保活监控
	wm.KeepLiveWorkers()
}

// 保活监控
func (wm *WorkerManager) KeepLiveWorkers() {
	// 如果有 Worker 已经死亡，则 workChan 会得到具体死亡的 Worker，然后输出异常，然后重启
	for wk := range wm.workerChan {
		// 打印错误日志
		fmt.Printf("Worker %d stopped with err:[%v]\n", wk.id, wk.err)

		// 重置错误
		wk.err = nil

		// 当前这个 wk 已经死亡了，需要重新启动它的业务
		go wk.work(wm.workerChan)
	}
}

/*
Worker
*/
type Worker struct {
	id  int
	err error
}

// 参数解释：
//
//	类型 chan<- *Worker 是一个只发送通道，表示这个函数只能向通道发送数据，不能从通道接收数据
//	而 chan *Worker 是一个双向通道
//	这样做明确函数意图、提高代码安全性
//
// 返回值解释：
//
//	err error，命名返回值。Go ​自动创建了一个名为 err 的变量，并初始化为该类型的零值​（对于 error 类型就是 nil）
//	这相当于隐式执行了 var err error（初始值为 nil）
func (wk *Worker) work(workerChan chan<- *Worker) (err error) {
	// 任何 goroutine 只要异常退出或者正常退出都会调用 defer，所以在 defer 中向 WorkerManager 的 workerChan 发送通知
	defer func() {
		// 捕获异常信息，防止 panic直接退出
		// recover() 是 Go 中用于捕获 panic 的内置函数
		// 当 goroutine 发生 panic 时，recover() 可以捕获 panic 的值
		// 只有在 defer 函数中调用 recover() 才有效
		// 这里使用它是为了防止 worker goroutine 因 panic 而完全退出，而是将错误信息记录下来并通知主 goroutine
		if r := recover(); r != nil {
			// 如果 panic 的是 error 类型，直接保存
			// 如果是其他类型，转换为 error 保存
			if err, ok := r.(error); ok {
				wk.err = err
			} else {
				wk.err = fmt.Errorf("panic happened with [%v]", r)
			}
		} else {
			wk.err = err
		}

		// 通知主 goroutine，当前子 goroutine 已经死亡
		workerChan <- wk
	}()

	// do something
	fmt.Println("Start Worker ... ID = ", wk.id)

	// 每个 Worker 睡眠一定时间之后，panic 退出或者 goexit 退出
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
	}

	// panic("worker panic...")
	runtime.Goexit()

	return err
}

/*
main
*/

func main() {
	wm := NewWorkerManager(10)
	wm.StartWorkerPool()
}
