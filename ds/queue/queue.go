package queue

import (
	"fmt"

	num "gitee.com/ryjer/go-generic/number"
)

// 不使用链表，从头构建。但会复制链表相关的代码
type Queue[T num.Q] struct {
	size            int
	header, trailer *Node[T] // 头节点，尾节点。作为[首, 末] 节点的哨兵
}

// 构造空队列
func New[T num.Q]() *Queue[T] {
	// 构造头尾哨兵节点
	headerNode := NewNode[T](0)
	trailerNode := NewNode[T](0)
	// 连接两哨兵
	headerNode.pre = nil
	headerNode.next = trailerNode
	trailerNode.pre = headerNode
	trailerNode.next = nil
	// 构造权柄
	return &Queue[T]{
		size:    0,
		header:  headerNode,
		trailer: trailerNode,
	}
}

// 从切片构造
func NewFromSlice[T num.Q](sourceSlice []T) *Queue[T] {
	newList := New[T]()             //初始化空链表
	for _, v := range sourceSlice { // 逐个入队元素
		newList.Enqueue(v)
	}
	return newList
}

// 转换为一个新的切片返回
func (this *Queue[T]) ToSlice() (newSlice []T) {
	newSlice = make([]T, this.size) // 构造新切片
	// 逐个填入数据
	for i, currentNode := 0, this.FrontNode(); i < len(newSlice) && currentNode != this.trailer; {
		newSlice[i] = currentNode.data
		// 后移索引
		i++
		currentNode = currentNode.next
	}
	return newSlice
}

// 容器方法
// 容量
func (this *Queue[T]) Size() int {
	return this.size
}

// 判空
func (this *Queue[T]) IsEmpty() (isEmpty bool) {
	if this.size == 0 {
		isEmpty = true
	} else {
		isEmpty = false
	}
	return
}

// 清空队列
func (this *Queue[T]) Clear() {
	this.size = 0
	this.header.next = this.trailer
	this.trailer.pre = this.header
}

// 尾部入列，将元素e加入队列尾部
func (this *Queue[T]) Enqueue(e T) *Node[T] {
	this.size++
	return this.trailer.InsertAsPre(e)
}

// 头部出列，删除队列头部节点，并返回被删除节点内的元素e
func (this *Queue[T]) Dequeue() (element T) {
	this.size--
	return this.header.next.Remove()
}

// 队首节点，返回队首节点
func (this *Queue[T]) FrontNode() *Node[T] {
	return this.header.next
}

// 队首
func (this *Queue[T]) Front() (element T) {
	return this.FrontNode().data
}

// 值判等，以值相等原则进行比较
// 定义：一个列表在"内容视图"上的相等包括：容量、链表的元素序列 相等，忽略其中的指针
func (this *Queue[T]) DeepEqual(another *Queue[T]) (equal bool) {
	if this.size != another.size { // 当容量不同时，不相等
		return false
	}
	// 当容量相同时，且为0时。不用逐个比较元素
	if this.size == 0 && another.size == 0 {
		return true
	}
	// 容量不为0时，逐个比较元素值
	for thisCurrent, anotherCurrent := this.header.next, another.header.next; (thisCurrent != this.trailer) || (anotherCurrent != another.trailer); thisCurrent, anotherCurrent = thisCurrent.next, anotherCurrent.next {
		if thisCurrent.data != anotherCurrent.data { // 一旦发现一个节点中的元素不相等，则认定数据视图不同
			return false
		}
	}
	// 全部相等后，可以判定相等
	return true
}

// 序列化函数
func (this *Queue[T]) String() (retString string) {
	// 打印头部
	retString = fmt.Sprintf("{%v [", this.size)
	// 打印列表元素序列，打印到倒数第2个
	if this.size > 0 { // 当列表不为空时进行打印
		p := this.header.next                     // 初始化p为首元素
		for ; p != this.trailer.pre; p = p.next { // 打印到倒数第2个
			retString += fmt.Sprintf("%v ", p.data)
		}
		retString += fmt.Sprintf("%v", p.data) // 打印最后一个元素值
	}
	// 补充尾部
	retString += "]}"
	return
}
