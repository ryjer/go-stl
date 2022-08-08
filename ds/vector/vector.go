package vector

import (
	"fmt"

	num "gitee.com/ryjer/go-generic/number"
)

const (
	defaultCapacity = 8 // 最小容量
)

// 数组列表数据结构，默认采用切片slice，使用数组需使用特定的方法
type Vector[T num.Q] struct {
	size     int // 已用容量
	capacity int // 总容量
	data     []T // 使用内置切片作为动态数组，但总是满的(len==cap)
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
		size:     len(elements),
		capacity: newCapacity,
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

// expand 空间拓展 n 个单位，未测试
func (this *Vector[T]) expand(n int) {
	newSize := this.size + n
	if newSize <= this.capacity { // 容量足够，无需扩容
		return
	}
	// 容量不足，进行扩容
	newCapacity := enoughCapacity(this.capacity)
	this.capacity = newCapacity // 更新容量
	newData := make([]T, newCapacity)
	copy(newData, this.data[0:this.size])
	this.data = newData
}

// Size 接口，返回已用空间
func (this *Vector[T]) Size() (usedSize int) {
	return this.size
}

// 检查是否为空
func (this *Vector[T]) IsEmpty() bool {
	if this.size == 0 {
		return true
	} else {
		return false
	}
}

// 读取元素
// 警告：当给出的秩r不在有效范围内时，会返回错误 err
func (this *Vector[T]) Get(r int) (element T, err error) {
	// 范围合法性检查
	if 0 <= r && r < this.size {
		return this.data[r], nil
	} else {
		return 0, fmt.Errorf("索引/秩越界 %v", r)
	}
}

// 更改元素
// 警告：当所给秩r不再有效范围内时，返回错误err
func (this *Vector[T]) Put(r int, newElement T) (err error) {
	// 范围合法性检查
	if 0 <= r && r < this.size {
		this.data[r] = newElement
		return nil
	} else {
		return fmt.Errorf("索引/秩越界 %v", r)
	}
}

// 插入元素，插入元素 element 到已被占用的秩 r
// 警告：不得插入未使用的秩处，尤其是最后一个秩之后的一个位置
func (this *Vector[T]) Insert(r int, element T) (rank int) {
	this.expand(1)                   // 检查扩容1个单位
	for i := this.size; i > r; i-- { // 从后向前
		this.data[i] = this.data[i-1] // 后继元素依次后移一个单位
	}
	this.data[r] = element // 置入新元素
	this.size++            //更新容量
	return rank
}

// 移除区间，并将其后的元素前移补全
func (this *Vector[T]) Remove(lo, hi int) (removedNumber int) {
	// lo，hi区间有效性检查
	if lo < 0 {
		lo = 0
	}
	if hi > this.size {
		hi = this.size
	}
	if lo >= hi {
		return (hi - lo)
	}
	for hi < this.size { // [lo,hi) 依次前移 hi-lo 位
		this.data[lo] = this.data[hi]
		lo++
		hi++
	}
	this.size = lo // 更新已用容量
	// this.shrink()  // 自动缩容，可选
	return (hi - lo)
}

// 移除单个元素，其后元素依次前移补全，返回被移除的元素
func (this *Vector[T]) Remove1(r int) (removedElement T) {
	removedElement = this.data[r]
	this.Remove(r, r+1)
	return
}

// 无序向量精确区间查找，从后向前精确查找 [lo, hi) 区间内元素的e，返回第一个匹配元素的秩，没有找到就返回-1
func (this *Vector[T]) Find(e T, lo, hi int) (rank int) {
	if this.size == 0 { // 空向量返回-1
		return -1
	}
	for (lo < hi) && (e != this.data[hi-1]) { // 使用 hi 从后向前扫描，直到找到e或者到达lo
		hi--
	}
	if (lo < hi) && this.data[hi-1] == e { // 当未到达lo且找到e时，返回对应秩
		return hi - 1
	} else { // 到达lo停止且没有找到时，返回-1
		return -1
	}
}

// 无序去重，可以保持低秩方向不同元素间的稳定性
func (this *Vector[T]) Deduplicate() (removedNumber int) {
	if this.size == 0 { // 空向量直接返回
		return 0
	}
	oldSize := this.size
	i := 1
	for i < this.size {
		if this.Find(this.data[i], 0, i) < 0 {
			i++
		} else {
			this.Remove1(i)
		}
	}
	return (oldSize - this.size)
}

// 遍历，使用参数中给定的函数，逐一处理向量中的元素
func (this *Vector[T]) Traverse(visit func(element *T)) {
	for i, _ := range this.data {
		visit(&this.data[i])
	}
}

// 逆序对统计
func (this *Vector[T]) Disordered() (disorderedNumber int) {
	var count int = 0
	for i := 1; i < this.size; i++ {
		if this.data[i-1] > this.data[i] {
			count++
		}
	}
	return count
}

// 有序向量 去重，返回被删除元素的个数
func (this *Vector[T]) Uniquify() (deletedNumber int) {
	// 空向量不处理
	if this.size == 0 {
		return 0
	}
	i, j := 0, 1
	for ; j < this.size; j++ {
		if this.data[i] != this.data[j] {
			i++
			this.data[i] = this.data[j]
		}
	}
	i++
	this.size = i // 更新size信息
	// this.shrink()     // 缩容
	return (j - i)
}

// 二分近似查找，在 [lo, hi) 区间查找元素 e，返回不大于e的元素的秩
// 警告：对空向量无法进行查找，会返回-2
func (this *Vector[T]) BinSearch(e T, lo, hi int) (rank int) {
	if this.size == 0 { // 空向量不查找，直接返回-1
		return -2
	}
	var mi int
	for lo < hi { // 不变性：arr[0,lo) <= e < arr[hi,n)
		mi = (lo + hi) / 2
		if e < this.data[mi] { // e < this.data[mi]，e∈[0, mi)
			hi = mi
		} else { // this.data[mi] <= e，e∈(mi, hi)
			lo = mi + 1
		}
	}
	return lo - 1
}

// 序列化方法
func (this *Vector[T]) String() string {
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
