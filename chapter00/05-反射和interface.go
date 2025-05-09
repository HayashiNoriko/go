package main

import (
	"fmt"
	"io"
	"os"
)

func main5() {
	// 以读写模式打开终端设备文件"/dev/tty"
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	// 声明一个io.Reader接口变量r
	// 将tty(实现了io.Reader接口)赋值给r
	// r 的 pair 为：(tty, *os.File)
	var r io.Reader
	r = tty

	// 声明一个io.Writer接口变量w
	// 使用类型断言将r转换为io.Writer类型
	// 这里能成功是因为tty实际上也实现了io.Writer接口
	// w 的 pair 为：(tty, *os.File)
	var w io.Writer
	w = r.(io.Writer) // 也可以写为 w = tty
	// w = r // 编译错误
	w.Write([]byte("HELLO THIS IS A TEST!!!\n"))
}
