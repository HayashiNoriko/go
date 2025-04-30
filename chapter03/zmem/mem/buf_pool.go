package mem

import (
	"errors"
	"fmt"
	"sync"
)

// 内存管理池类型
// key 表示该组 buffer 的 Capacity
// value 是一个 Buf 链表，每个 Buf 的 Capacity 都相同
type Pool map[int]*Buf

// Buf 内存池
type BufPool struct {
	// 所有 buffer 的一个 map 集合句柄
	Pool     Pool
	PoolLock sync.RWMutex

	// 总 buffer 池的内存大小（空闲的，未分配出去的），单位为 KB
	TotalMem uint64
}

// 单例对象
var bufPoolInstance *BufPool
var once sync.Once

// 获取 BufPool 对象（单利模式）
func MemPool() *BufPool {
	once.Do(func() {
		bufPoolInstance = new(BufPool)
		bufPoolInstance.Pool = make(map[int]*Buf)
		bufPoolInstance.TotalMem = 0
		//bufPoolInstance.prev=nil
		bufPoolInstance.initPool()
	})

	return bufPoolInstance
}

// 初始化 Pool
const (
	m4K   int = 4096
	m16K  int = 16384
	m64K  int = 655535
	m256K int = 262144
	m1M   int = 1048576
	m4M   int = 4194304
	m8M   int = 8388608
)

func (bp *BufPool) initPool() {
	// 开辟 4KB buffer 的内存池
	// 预先开辟 5000 个，约 20MB 供开发者使用
	bp.makeBufList(m4K, 5000)

	// 开辟 16KB buffer 的内存池
	// 预先开辟 1000 个，约 16MB 供开发者使用
	bp.makeBufList(m16K, 1000)

	// 开辟 64KB buffer 的内存池
	// 预先开辟 500 个，约 32MB 供开发者使用
	bp.makeBufList(m64K, 500)

	// 开辟 256KB buffer 的内存池
	// 预先开辟 200 个，约 50MB 供开发者使用
	bp.makeBufList(m256K, 200)

	// 开辟 1MB buffer 的内存池
	// 预先开辟 50 个，约 50MB 供开发者使用
	bp.makeBufList(m1M, 50)

	// 开辟 4MB buffer 的内存池
	// 预先开辟 20 个，约 80MB 供开发者使用
	bp.makeBufList(m4M, 20)

	// 开辟 8MB buffer 的内存池
	// 预先开辟 10 个，约 80MB 供开发者使用【io_buf】
	bp.makeBufList(m8M, 10)
}

// 创建一串 Buf 链表
func (bp *BufPool) makeBufList(cap int, num int) {
	bp.Pool[cap] = NewBuf(cap)

	var prev *Buf
	prev = bp.Pool[cap]
	for i := 1; i < num; i++ {
		prev.Next = NewBuf(cap)
		prev = prev.Next
	}

	// 内存池更新总容量
	bp.TotalMem += (uint64(cap) / 1024) * uint64(num)
}

// 用户从内存池中申请 buffer
const (
	// 总内存池的最大限制（KB）
	EXTRA_MEM_LIMIT int = 5 * 1024 * 1024
)

func (bp *BufPool) Alloc(N int) (*Buf, error) {
	// 1. 找到 N 最接近哪个 hash 组
	var index int
	if N <= m4K {
		index = m4K
	} else if N <= m16K {
		index = m16K
	} else if N <= m64K {
		index = m64K
	} else if N <= m256K {
		index = m256K
	} else if N <= m1M {
		index = m1M
	} else if N <= m4M {
		index = m4M
	} else if N <= m8M {
		index = m8M
	} else {
		return nil, errors.New("alloc size Too Large")
	}

	bp.PoolLock.Lock() // 要申请的 index已经确定，可以开始准备申请了，上锁

	var resBuf *Buf
	if bp.Pool[index] == nil {
		// 2. 如果这个 buf 链还没有，则需要新建
		// 检查如果申请的话是否会超出最大容量
		if bp.TotalMem+uint64(index/1024) > uint64(EXTRA_MEM_LIMIT) {
			return nil, errors.New("already use too many memory")
		}

		resBuf = NewBuf(index)
		// bp.TotalMem += uint64(index / 1024) // 这里不应该加啊
	} else {
		// 3.如果这个 buf 链已经存在，则返回其中的一个 buffer 块，并从pool 中移除它
		resBuf = bp.Pool[index]
		bp.Pool[index] = resBuf.Next
		bp.TotalMem -= uint64(index / 1024)
		resBuf.Next = nil
	}

	bp.PoolLock.Unlock() // 已申请好，可以解锁

	fmt.Printf("Alloc Mem Size: %d KB\n", resBuf.Capacity/1024)
	return resBuf, nil

}

// 用户退还 buffer
func (bp *BufPool) Revert(buf *Buf) error {
	index := buf.Capacity

	bp.PoolLock.Lock()

	// 检查 1，是否有相应的 index位
	if _, ok := bp.Pool[index]; !ok {
		errStr := fmt.Sprintf("index %d not in the BufPool", index)
		return errors.New(errStr)
	}

	// 检查 2，是否有足够空间可以归还
	if bp.TotalMem+uint64(index/1024) > uint64(EXTRA_MEM_LIMIT) {
		return errors.New("can not revert, too much memory")
	}

	// 重置 buf 中的内置位置指针
	buf.Clear()

	// 将 buffer 插回链表头部
	buf.Next = bp.Pool[index]
	bp.Pool[index] = buf
	bp.TotalMem += uint64(index / 1024)

	bp.PoolLock.Unlock()

	fmt.Printf("Revert Mem Size: %d KB\n", index/1024)

	return nil
}
