package mem_test

import (
	"fmt"
	"hello/chapter03/zmem/mem"
	"testing"
)

func TestBufPoolSetGet(t *testing.T) {
	pool := mem.MemPool()

	buffer, err := pool.Alloc(1)
	if err != nil {
		fmt.Println("pool Alloc Error", err)
		return
	}

	buffer.SetBytes([]byte("Hello world"))

	fmt.Printf("GetBytes = %+v, ToString = %s\n", buffer.GetBytes(), string(buffer.GetBytes()))

	buffer.Pop(4)

	fmt.Printf("GetBytes = %+v, ToString = %s\n", buffer.GetBytes(), string(buffer.GetBytes()))
}

func TestBufPoolCopy(t *testing.T) {
	pool := mem.MemPool()

	buffer, err := pool.Alloc(1)
	if err != nil {
		fmt.Println("pool Alloc Error", err)
		return
	}

	buffer.SetBytes([]byte("Hello world"))
	fmt.Printf("buffer1 = %s\n", string(buffer.GetBytes()))

	buffer2, err := pool.Alloc(1)
	if err != nil {
		fmt.Println("pool Alloc Error", err)
		return
	}
	buffer2.Copy(buffer)
	fmt.Printf("buffer2 = %s\n", string(buffer2.GetBytes()))

}

func TestBufPoolAdjust(t *testing.T) {
	pool := mem.MemPool()

	buffer, err := pool.Alloc(1)
	if err != nil {
		fmt.Println("pool Alloc Error", err)
		return
	}

	buffer.SetBytes([]byte("Hello world"))
	fmt.Printf("GetBytes = %+v, Head = %d, Length = %d\n", buffer.GetBytes(), buffer.Head(), buffer.Length())

	buffer.Pop(4)
	fmt.Printf("GetBytes = %+v, Head = %d, Length = %d\n", buffer.GetBytes(), buffer.Head(), buffer.Length())

	buffer.Adjust()
	fmt.Printf("GetBytes = %+v, Head = %d, Length = %d\n", buffer.GetBytes(), buffer.Head(), buffer.Length())
}
