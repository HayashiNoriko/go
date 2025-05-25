package main

import "fmt"

// 接口
type IBroker interface {
	GetPendingTasks()
}

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

func (rb *RedisBroker) GetPendingTasks() {
	fmt.Println("queue=", rb.queue)
	fmt.Println("Broker.queue=", rb.Broker.queue)
	fmt.Println("Get pending tasks from Redis")
	fmt.Println(rb.Broker.pendingTasks)
}

func main14() {
	var broker IBroker = &RedisBroker{
		Broker: Broker{
			pendingTasks: []string{"task1", "task2"},
			queue:        "base tasks",
		},
		//queue: "derived tasks",
	}

	// 调用子类方法，成功
	broker.GetPendingTasks()

	// 直接调用父类方法和字段，失败
	// 因为 broker 是 IBrouker 接口类型，不是 RedisBroker 对象，还不具有 Broker 字段和 queue 字段
	// broker.Broker.GetPendingTasks()    // 报错
	// fmt.Println("queue=",broker.queue) // 报错

	// 使用类型断言，把接口转换成子类对象，再调用父类方法和字段
	if rb, ok := broker.(*RedisBroker); ok {
		rb.Broker.GetPendingTasks()
		fmt.Println("Broker.queue=", rb.queue)
	}

}
