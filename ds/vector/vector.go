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

// New 空构造
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
	newCapacity := enoughCapacity(len(elements)) // 计算容量
	newSlice := make([]T, newCapacity)           // 创建新切片作为数据空间
	copy(newSlice, elements)                     // 复制切片数据
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
	// 错误检查，选定长度非法或为0时返回0长度新向量
	if newCapacity := hi - lo; newCapacity <= 0 {
		return New[T]()
	}
	// 创建足够容量的新0长度向量
	newVector := Vector[T]{
		size:     0,
		capacity: defaultCapacity,
		data:     make([]T, defaultCapacity),
	}
	// 复制
	newVector.copyFrom(anotherVector, lo, hi)
	return &newVector
}

// copyFrom 向量区间复制方法：[lo,hi)区间
func (this *Vector[T]) copyFrom(another *Vector[T], lo, hi int) {
	newSize := hi - lo
	this.size = newSize // 更新已用容量
	newCapacity := enoughCapacity(newSize)
	// 惰性空间分配，只有空间不足时才会切换 动态空间
	if this.capacity < newCapacity {
		this.capacity = newCapacity       // 更新总容量
		newData := make([]T, newCapacity) // 分配2倍新空间
		this.data = newData               // 更换新空间，丢弃原数据空间
	}
	copy(this.data, another.data[lo:hi]) // 复制数据
}

// expand 空间拓展 n 个单位
func (this *Vector[T]) expand(n int) {
	newSize := this.size + n
	if newSize <= this.capacity { // 容量足够，无需扩容
		return
	}
	// 容量不足，进行扩容
	newCapacity := enoughCapacity(this.capacity)
	this.capacity = newCapacity
	newData := make([]T, newCapacity)
	copy(newData, this.data[0:this.size])
	this.data = newData
}

// Size 接口，返回已用空间
func (this *Vector[T]) Size() (usedSize int) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	return this.size
}

// Get 接口，读取元素
// 警告：当给出的秩r不在有效范围内时，会返回错误 err
func (this *Vector[T]) Get(r int) (element T, err error) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	// 范围合法性检查
	if 0 <= r && r < this.size {
		return this.data[r], nil
	} else {
		return 0, fmt.Errorf("索引/秩越界 %v", r)
	}
}

// Put 接口，更改元素
// 警告：当所给秩r不再有效范围内时，返回错误err
func (this *Vector[T]) Put(r int, newElement T) (err error) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	// 范围合法性检查
	if 0 <= r && r < this.size {
		this.data[r] = newElement
		return nil
	} else {
		return fmt.Errorf("索引/秩越界 %v", r)
	}
}

// 序列化方法
func (this *Vector[T]) String() string {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	return fmt.Sprintf("{%v %v %v}", this.size, this.capacity, this.data[:this.size])
}

// 协助空间计算函数，计算给定已用容量下的新容量
// 策略：总是给出2倍空间，当倍增后依然小于最低容量时给出更大的最低容量
func enoughCapacity(newSize int) int {
	newCapacity := newSize * 2
	if newSize < defaultCapacity {
		newCapacity = defaultCapacity
	}
	return newCapacity
}
