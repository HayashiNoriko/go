// go 基本语法
package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"time"
)

func main() {

	// 1、变量声明。必须使用，否则报错
	var a int = 666 // 显示类型
	b := 20         // 类型推断（最常用）

	fmt.Println("hello world")
	fmt.Printf("a=%d b=%d\n", a, b)

	// 2、常量声明
	const Pi = 3.14
	fmt.Printf("Pi=%f\n", Pi)

	// 3、新类型
	var bt byte = 255        // 相当于 uint8，范围 0～255，超出报错
	var rn rune = 2147483647 // 相当于 int32，范围 -2147483648 ~ 2147483647，超出报错
	var c64 complex64 = 1    // 复数类型
	var c128 complex128 = 2  // 复数类型
	fmt.Printf("bt=%d rn=%d\n c64=%f c128=%f\n", bt, rn, c64, c128)

	// 4、字符串
	s1 := "hello"                              // 直接创建，不可修改
	s2 := []byte{0x77, 0x6f, 0x72, 0x6c, 0x64} // 从字节切片创建，可以修改
	fmt.Printf("s=%s s2=%s\n", s1, s2)

	// 5、if 语句，条件不需要括号
	if a > 10 {
		fmt.Println("a>10")
	} else {
		fmt.Println("a<=10")
	}

	// 6、for 循环
	// 传统三段式
	for i := 0; i < 10; i++ {
		fmt.Printf("for i=%d\n", i)
	}
	// while替代
	x := 10
	for x > 0 {
		fmt.Printf("for x=%d\n", x)
		x--
	}
	// 无线循环
	// for {
	// }

	// 7、switch，更强大，不需要 break
	switch x {
	case 0:
		fmt.Println("x = 0")
	case 1, 2:
		fmt.Println("x = 1 or 2")
	default:
		fmt.Println("x != 0, 1 or 2")
	}

	// 8、函数特性
	fmt.Println(div(3, 2))
	fmt.Println(div(3, 0))
	xx, yy := split(10)
	fmt.Println(xx, yy)

	// 9、复合数据类型
	var array [3]int // 数组，值类型，长度固定
	array[0] = 10
	array[1] = 20
	array[2] = 30
	// var array = [3]int{10, 20, 30} // 数组的另一种初始化方式
	for i, value := range array {
		fmt.Printf("array[%d]=%d\n", i, value)
	}

	slice := []int{1, 2, 3} // 切片，动态数组
	slice = append(slice, 4)
	for i, value := range slice {
		fmt.Printf("slice[%d]=%d\n", i, value)
	}

	m := map[string]int{ // map，内置字典类型
		"Alice": 25,
		"Bob":   30,
	}
	for name, age := range m {
		fmt.Printf("m[%s]=%d\n", name, age)
	}

	// 10、结构体和方法（实现封装）
	// var p1 Point
	// p1.X = 1
	// p1.Y = 1
	p1 := Point{
		X: 1,
		Y: 1,
	}
	distance := p1.Distance()
	fmt.Printf("distance=%f\n", distance)

	// 11、接口（实现多态）
	d := Dog{}
	c := Cat{}
	MakeSound(d)
	MakeSound(c)

	// 12、组合（实现继承）
	r := Rice{
		Food:  Food{name: "五常大米"},
		color: "白色",
	}
	r.eaten()

	// 13、错误处理
	res, err := div(3, 0)
	if err != nil {
		// 处理错误
		fmt.Println("除零错误!")
	} else {
		fmt.Printf("res=%d\n", res)
	}

	// 14、并发编程
	// Goroutine：轻量级线程
	go func() {
		// 并发执行的代码
		for i := 0; i < 3; i++ {
			fmt.Println("并发执行中...")
			time.Sleep(time.Second)
		}
	}()
	for i := 0; i < 3; i++ {
		fmt.Println("main 函数中...")
		time.Sleep(time.Second)
	}
	// Channel：通信机制
	// 示例一
	ch := make(chan int)
	go func() { ch <- 123 }()
	value := <-ch
	fmt.Printf("value=%d\n", value)

	// 示例二
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 启动3个worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // 关闭channel表示没有更多任务

	// 收集结果
	for a := 1; a <= 5; a++ {
		<-results
	}
}

// 8、函数特性
// 多返回值
func div(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除零错误")
	}
	return a / b, nil
}

// 命名返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // 自动返回 x,y
}

// defer，延迟执行，常用语资源清理
func readFile() {
	f, _ := os.Open("file.txt")
	defer f.Close() // 函数返回前执行
	// 处理文件...
}

// 10、结构体和方法
// 类似 C 的结构体
type Point struct {
	X, Y int
}

// 为类型添加方法
func (p *Point) Distance() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

// 11、接口，隐式实现，不需要显示声明实现关系
type Speaker interface {
	Speak()
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

func (c Cat) Speak() {
	fmt.Println("Meow!")
}

func MakeSound(s Speaker) {
	s.Speak()
}

// 12、组合
type Food struct {
	name string
}

func (f *Food) eaten() {
	fmt.Printf("%s被吃了\n", f.name)
}

type Rice struct {
	Food
	color string
}

// 14、并发编程channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker %d 处理任务 %d\n", id, j)
		results <- j * 2
	}
}
