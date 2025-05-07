package main

// 显卡
type Card interface {
	display()
}

// 内存
type Memory interface {
	storage()
}

// 处理器
type CPU interface {
	calculate()
}

// Intel 家产品
type IntelCard struct {
	Card
}

func (c IntelCard) display() {
	println("Intel card is displaying")
}

type IntelMemory struct {
	Memory
}

func (m IntelMemory) storage() {
	println("Intel memory is storaging")
}

type IntelCPU struct {
	CPU
}

func (cpu IntelCPU) calculate() {
	println("Intel CPU is calculating")
}

// Kinston 家产品
type KinstonMemory struct {
	Memory
}

func (m KinstonMemory) storage() {
	println("Kinston memory is storaging")
}

// NVIDIA 家产品
type NVIDIACard struct {
	Card
}

func (c NVIDIACard) display() {
	println("NVIDIA card is displaying")
}

func buildComputer(card Card, memory Memory, cpu CPU) {
	card.display()
	memory.storage()
	cpu.calculate()
}

func main() {
	buildComputer(&IntelCard{}, &IntelMemory{}, &IntelCPU{})
	buildComputer(&NVIDIACard{}, &KinstonMemory{}, &IntelCPU{})
}
