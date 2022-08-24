package doublyLinkedlist

import (
	num "gitee.com/ryjer/go-generic/number"
)

// 链表节点定义
type Node[T num.Q] struct {
	data      T
	pre, next *Node[T]
}

// 节点构造方法
func NewNode[T num.Q](e T) *Node[T] {
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

// 前一节点
func (this *Node[T]) Prev() *Node[T] {
	return this.pre
}

// 后一节点
func (this *Node[T]) Next() *Node[T] {
	return this.next
}

// 前插入算法
func (this *Node[T]) insertPre(e T) (xnode *Node[T]) {
	xnode = FullNew(e, this.pre, this)
	this.pre.next = xnode
	this.pre = xnode
	return xnode
}

// 后插入算法
func (this *Node[T]) insertNext(e T) (xnode *Node[T]) {
	xnode = FullNew(e, this, this.next)
	this.next.pre = xnode
	this.next = xnode
	return xnode
}

// 节点值判等
func (this *Node[T]) deepEqual(another *Node[T]) (equal bool) {
	if (this.data == another.data) && (this.pre == another.pre) && (this.next == another.next) {
		return true
	} else {
		return false
	}
}
