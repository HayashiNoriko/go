package mem

import (
	"C"
	"hello/chapter03/zmem/c"
	"unsafe"
)
import "fmt"

type Buf struct {
	// 如果存在多个buffer，则采用链表的形式连接起来
	Next *Buf
	// 当前 buffer 的缓存容量大小
	Capacity int
	// 当前 buffer 的有效数据长度
	length int
	// 未处理数据的头部位置索引
	head int
	// 当前 buf 所保存的数据地址
	data unsafe.Pointer
}

// 创建一个buffer
func NewBuf(size int) *Buf {
	return &Buf{
		Capacity: size,
		length:   0,
		head:     0,
		Next:     nil,
		data:     c.Malloc(size),
	}
}

// 给一个buffer 填充[]byte 数据
func (b *Buf) SetBytes(src []byte) {
	c.Memcpy(unsafe.Pointer(uintptr(b.data)+uintptr(b.head)), src, len(src))
	b.length += len(src)
}

// 获取一个buffer 的数据，以[]Byte 形式展现
func (b *Buf) GetBytes() []byte {
	data := C.GoBytes(unsafe.Pointer(uintptr(b.data)+uintptr(b.head)), C.int(b.length))
	return data
}

// 将其他 buffer 对象复制到自己中
func (b *Buf) Copy(other *Buf) {
	c.Memcpy(b.data, other.GetBytes(), other.length)
	b.head = 0
	b.length = other.length
}

// 处理长度为 len 的数据，移动 head 和修正 length
func (b *Buf) Pop(len int) {
	if b.data == nil {
		fmt.Printf("pop data is nil")
		return
	}
	if len > b.length {
		fmt.Printf("pop len > length")
		return
	}
	b.length -= len
	b.head += len
}

// 将已经处理过的数据清空，将未处理的数据提至数据首地址
func (b *Buf) Adjust() {
	if b.head != 0 {
		if b.length != 0 {
			c.Memmove(b.data, unsafe.Pointer(uintptr(b.data)+uintptr(b.head)), b.length)
		}
		b.head = 0
	}
}

// 清空数据
func (b *Buf) Clear() {
	b.length = 0
	b.head = 0
}

// 其他的提供的访问 head 和 length 的方法
func (b *Buf) Head() int {
	return b.head
}

func (b *Buf) Length() int {
	return b.length
}
