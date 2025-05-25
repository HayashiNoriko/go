package main

import "fmt"

// 父类
type Broker struct {
	pendingTasks []string
	queue        string
}

func (b *Broker) GetPendingTasks() {
	fmt.Println("Not implement")
}

// 子类
type RedisBroker struct {
	Broker
	// queue string // 如果有，会覆盖父类字段；如果没有，会继承父类字段（rb.queue == rb.Broker.queue）
}

// 子类方法，如果有，会覆盖父类方法；如果没有，会继承父类方法（rb.GetPendingTasks == rb.Broker.GetPendingTasks）
func (rb *RedisBroker) GetPendingTasks() {
	fmt.Println("Redis broker pending tasks:", rb.pendingTasks)
}

func main14() {
	broker := &RedisBroker{
		Broker: Broker{
			pendingTasks: []string{"task1", "task2"},
			queue:        "base tasks",
		},
		//queue: "derived tasks",
	}

	// 调用子类方法（如果子类实现了这个方法，那么就会调用子类的版本；如果没有实现，那么就会调用父类的版本）
	broker.GetPendingTasks()

	// 显示调用父类方法
	broker.Broker.GetPendingTasks()

}
