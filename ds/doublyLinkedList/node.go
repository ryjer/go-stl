package doublyLinkedlist

import (
	num "gitee.com/ryjer/go-generic/number"
)

// 链表节点定义
type Node[T num.Q] struct {
	data      T
	pre, next *Node[T]
}

// 节点值判等，底层值相等，会比较其中的指针指向
// 定义：一个节点在 值 层次上的相等包括 容纳的元素e和前后指针
func (this *Node[T]) DeepEqual(another *Node[T]) (equal bool) {
	if (this.data == another.data) && (this.pre == another.pre) && (this.next == another.next) {
		return true
	} else {
		return false
	}
}

// 节点构造方法
func NewNode[T num.Q](e T) *Node[T] {
	return &Node[T]{
		data: e,
		pre:  nil, next: nil,
	}
}

// 节点完全构造，提供节点的所有信息进行构造
// 提示：不会修改参数中被指向的节点，调用方需要自行调整被指向节点的指针信息
func FullNewNode[T num.Q](e T, pre, next *Node[T]) *Node[T] {
	return &Node[T]{
		data: e,
		pre:  pre,
		next: next,
	}
}

// 获取元素
func (this *Node[T]) Get() (element T) {
	return this.data
}

// 修改元素，并返回原元素
func (this *Node[T]) Put(e T) (element T) {
	element = this.data
	this.data = e
	return element
}

// 前一节点，返回当前节点的直接前驱节点的引用
func (this *Node[T]) PreNode() *Node[T] {
	return this.pre
}

// 后一节点，返回当前节点的直接后继节点的引用
func (this *Node[T]) NextNode() *Node[T] {
	return this.next
}

// 前插入算法，作为当前节点的直接前驱插入，返回插入节点的地址
// 提示：可以在链表的头部插入，即使头部没有前驱节点
func (this *Node[T]) InsertAsPre(e T) (xnode *Node[T]) {
	xnode = FullNewNode(e, this.pre, this)
	if this.pre != nil { // 当前驱哨兵存在时
		this.pre.next = xnode // 重定位前驱哨兵的 next 指针
	}
	this.pre = xnode
	return xnode
}

// 后插入算法，作为当前节点的直接后继插入，返回插入节点的地址
func (this *Node[T]) InsertAsNext(e T) (xnode *Node[T]) {
	xnode = FullNewNode(e, this, this.next)
	if this.next != nil { // 当后继哨兵存在时
		this.next.pre = xnode // 重定位后继哨兵的 pre 指针
	}
	this.next = xnode // 重定位后继节点的 pre 指针
	return xnode
}

// 移除节点，返回被移除节点内的元素值
func (this *Node[T]) Remove() (element T) {
	element = this.data
	pre := this.pre
	next := this.next
	pre.next = next
	next.pre = pre
	return
}
