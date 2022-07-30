package vector

import (
	"fmt"
	"sync"

	num "gitee.com/ryjer/go-generic/number"
)

const (
	defaultCapacity = 8 // 最小容量
)

// 数组列表数据结构，默认采用切片slice，使用数组需使用特定的方法
type Vector[T num.Q] struct {
	mutex    sync.Mutex // 并发控制锁
	size     int        // 已用容量
	capacity int        // 总容量
	data     []T        // 使用内置切片作为动态数组，但总是满的(len==cap)
}

// New 构造
func New[T num.Q]() *Vector[T] {
	newVector := Vector[T]{
		size:     0,
		capacity: defaultCapacity,
		data:     make([]T, defaultCapacity),
	}
	return &newVector
}

// NewFromSlice 从已有切片构造
func NewFromSlice[T num.Q](elements []T) *Vector[T] {
	var newSlice []T
	newCapacity := defaultCapacity
	if newCapacity < len(elements) { // 当传入元素数小于最低容量时，使用最低容量
		for newCapacity < len(elements) {
			newCapacity *= 2
		}
	}
	// 创建新切片作为数据空间
	newSlice = make([]T, newCapacity)
	copy(newSlice, elements)
	// 构造新列表/向量
	newVector := Vector[T]{
		mutex:    sync.Mutex{},
		size:     len(elements),
		capacity: cap(newSlice),
		data:     newSlice,
	}
	return &newVector
}

// NewFromVector 从已有列表/向量子区间构造
func NewFromVector[T num.Q](anotherVector *Vector[T], lo, hi int) *Vector[T] {
	newVector := Vector[T]{
		size:     0,
		capacity: defaultCapacity,
		data:     make([]T, defaultCapacity),
	}
	newVector.copyFrom(anotherVector, lo, hi)
	return &newVector
}

// copyFrom 列表/向量间复制方法
func (this *Vector[T]) copyFrom(another *Vector[T], lo, hi int) {
	// 调整空间，预留2倍空间
	newCapacity := 2 * (hi - lo)
	newData := make([]T, newCapacity)
	// 复制数据到新空间
	copy(newData, another.data[lo:hi])
	// 更换新空间
	this.data = newData
}

// expand 空间拓展方法
func (this *Vector[T]) expand() {

}

// Size 接口，返回已用空间
func (this *Vector[T]) Size() (usedSize int) {
	return this.size
}

// 序列化方法
func (this *Vector[T]) String() string {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	return fmt.Sprintf("{%v %v %v}", this.size, this.capacity, this.data[:this.size])
}

// 协助空间计算函数，计算能够覆盖对应
func enoughSize(n int) int {
	var newSize int = defaultCapacity
	// 扩容到一个最小的2^n，未来将借助计算机的二进制位运算机制优化
	for newSize < n {
		newSize *= 2
	}
	return newSize
}
