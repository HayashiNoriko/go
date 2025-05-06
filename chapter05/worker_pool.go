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

func (wk *Worker) work(workerChan chan<- *Worker) (err error) {
	// 任何 goroutine 只要异常退出或者正常退出都会调用 defer，所以在 defer 中向 WorkerManager 的 workerChan 发送通知
	defer func() {
		// 捕获异常信息，防止 panic直接退出
		if r := recover(); r != nil {
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
