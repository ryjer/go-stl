package twoWayNode

import (
	num "gitee.com/ryjer/go-generic/number"
)

// 链表节点定义
type Node[T num.Q] struct {
	data      T
	pre, next *Node[T]
}

// 节点值判等，底层值相等，会比较其中的指针指向
func (this *Node[T]) DeepEqual(another *Node[T]) (equal bool) {
	if (this.data == another.data) && (this.pre == another.pre) && (this.next == another.next) {
		return true
	} else {
		return false
	}
}

// 节点构造方法
func New[T num.Q](e T) *Node[T] {
	return &Node[T]{
		data: e,
		pre:  nil, next: nil,
	}
}

// 节点完全构造，提供节点的所有信息进行构造
func FullNew[T num.Q](e T, pre, next *Node[T]) *Node[T] {
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

// 前一节点，返回当前节点的直接前驱节点的引用
func (this *Node[T]) PreNode() *Node[T] {
	return this.pre
}

// 后一节点，返回当前节点的直接后继节点的引用
func (this *Node[T]) NextNode() *Node[T] {
	return this.next
}

// 前插入算法，作为当前节点的直接前驱插入，返回插入节点的地址
func (this *Node[T]) insertAsPre(e T) (xnode *Node[T]) {
	xnode = FullNew(e, this.pre, this)
	this.pre.next = xnode
	this.pre = xnode
	return xnode
}

// 后插入算法，作为当前节点的直接后继插入，返回插入节点的地址
func (this *Node[T]) insertAsNext(e T) (xnode *Node[T]) {
	xnode = FullNew(e, this, this.next)
	this.next.pre = xnode // 重定位前驱节点的 next 指针
	this.next = xnode     // 重定位后继节点的 pre 指针
	return xnode
}
