package c_test

import (
	// "zmem/c"
	"bytes"
	"encoding/binary"
	"fmt"
	"hello/chapter03/zmem/c"
	"testing"
	"unsafe"
)

func IsLittleEndian() bool {
	var n int32 = 0x01020304

	// 将 int32 指针转换为 Byte 类型的指针
	u := unsafe.Pointer(&n)
	pb := (*byte)(u)

	// 取得pb 位置对应的值
	b := *pb

	// 由于 b 是 byte 类型，最多保存 8 位，所以只能取得开始的 8 位
	// 小端：04
	// 大端：01
	return (b == 0x04)
}

func IntToBytes(n uint32) []byte {
	x := int32(n)
	BytesBuffer := bytes.NewBuffer([]byte{})

	var order binary.ByteOrder
	if IsLittleEndian() {
		order = binary.LittleEndian
	} else {
		order = binary.BigEndian
	}

	binary.Write(BytesBuffer, order, x)

	return BytesBuffer.Bytes()
}

func TestMemoryC(t *testing.T) {
	// 1、开辟 4 字节内存
	data := c.Malloc(4) // data 是 unsafe.Pointer
	fmt.Printf(" data %+v, %T\n", data, data)

	// 2、将这 4 字节赋值为 5
	myData := (*uint32)(data) // myData 是 *uint32，这才可以在 go 程序里赋值
	*myData = 5
	fmt.Printf(" data %+v, %T\n", *myData, *myData)

	// 3、将 100 通过 Memcpy 复制给这 4 字节
	var a uint32 = 100
	c.Memcpy(data, IntToBytes(a), 4)
	fmt.Printf(" data %+v, %T\n", *myData, *myData)

	c.Free(data)
}
