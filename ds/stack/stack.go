package stack

import (
	"fmt"

	num "gitee.com/ryjer/go-generic/number"
)

const defaultCapacity = 4 // 最小容量

type Stack[T num.Q] struct {
	size     int // 已用容量，指向栈顶第一个空位
	capacity int // 总容量
	data     []T // 使用内置切片作为栈数据域，但总是满的(len==cap)
}

// New 空构造
func New[T num.Q]() (newStack *Stack[T]) {
	return &Stack[T]{
		size:     0,
		capacity: defaultCapacity,
		data:     make([]T, defaultCapacity),
	}
}

// 从切片构造
func NewFromSlice[T num.Q](sourceSlice []T) (newStack *Stack[T]) {
	newSize := len(sourceSlice)
	newCapacity := cap(sourceSlice)
	newData := make([]T, newSize)
	copy(newData, sourceSlice)
	return &Stack[T]{
		size:     newSize,
		capacity: newCapacity,
		data:     newData,
	}
}

// 获取已用容量
func (this *Stack[T]) Size() int {
	return this.size
}

// 判空
func (this *Stack[T]) IsEmpty() bool {
	if this.size == 0 {
		return true
	} else {
		return false
	}
}

// 清空
func (this *Stack[T]) Clear() {
	this.data = make([]T, 0) // 将数据域替换为空切片
	this.size = 0
	this.capacity = defaultCapacity // 初始化容量和长度
}

// 检查拓展，自动扩容保证有至少1个余量可用
func (this *Stack[T]) expand() {
	if this.size < this.capacity { // 容量足够，无需扩容
		return
	} else { // 容量不足，进行扩容
		this.capacity *= 2                    // 容量翻倍扩展
		newData := make([]T, this.capacity)   // 创建新切片
		copy(newData, this.data[0:this.size]) // 转移数据到新切片
		this.data = newData                   // 切换描述块中的数据域为新切片
	}
}

// 压栈
func (this *Stack[T]) Push(e T) {
	this.expand()            // 检查拓展容量
	this.data[this.size] = e // 栈顶
	this.size++              // 更新容量
}

// 检查缩容一个单位
func (this *Stack[T]) shrink() {
	newSize := this.size - 1
	// 当已用容量少于一半，且总容量超出最小容量的2倍时进行缩容
	if (newSize < (this.capacity / 2)) && (this.capacity >= (defaultCapacity * 2)) {
		this.capacity /= 2                    // 减半缩容
		newData := make([]T, this.capacity)   // 创建新切片
		copy(newData, this.data[0:this.size]) // 转移数据到新切片
		this.data = newData                   // 切换描述块中的数据域为新切片
	}
}

// 出栈
// 警告：不会检查自身是否为空，请调用方自行检查
func (this *Stack[T]) Pop() (element T) {
	this.size--
	element = this.data[this.size] // 先移动指针，后弹出数据
	this.shrink()                  // 检查缩容
	return element
}

// 获取栈顶元素
func (this *Stack[T]) Top() (element T) {
	return this.data[this.size-1]
}

// 值判等，内容视图判等
func (this *Stack[T]) DeepEqual(another *Stack[T]) bool {
	// 已用容量不同，不相等
	if this.size != another.size {
		return false
	}
	// 当容量相同时，且为0时。不用逐个比较元素
	if this.size == 0 && another.size == 0 {
		return true
	}
	// 已用容量相同且不为0时，逐个元素比对
	for i := 0; i < this.size; i++ {
		if this.data[i] != another.data[i] {
			return false
		}
	}
	// 都相等，则判断
	return true
}

// 序列化方法
func (this *Stack[T]) String() string {
	return fmt.Sprintf("{%v %v %v}", this.size, this.capacity, this.data[:this.size])
}
