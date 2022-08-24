package doublyLinkedlist

import (
	num "gitee.com/ryjer/go-generic/number"
)

// 双向链表结构体定义，其中默认有两个哨兵节点：头节点header 和 尾节点trailer
type DoublyLinkedList[T num.Q] struct {
	size            int
	header, trailer *Node[T] // 头节点，尾节点
}

// 构造方法
func New[T num.Q]() *DoublyLinkedList[T] {
	// 构造头尾哨兵节点
	headerNode := NewNode[T](0)
	trailerNode := NewNode[T](0)
	// 调整指针
	headerNode.pre = nil
	headerNode.next = trailerNode
	trailerNode.pre = headerNode
	trailerNode.next = nil
	// 构造权柄
	return &DoublyLinkedList[T]{
		size:    0,
		header:  headerNode,
		trailer: trailerNode,
	}
}

// 容量
func (this *DoublyLinkedList[T]) Size() (size int) {
	return this.size
}

// 空判断
func (this *DoublyLinkedList[T]) IsEmpty(isEmpty bool) {
	if this.size == 0 {
		isEmpty = true
	} else {
		isEmpty = false
	}
}

// 读取元素，寻秩访问
func (this *DoublyLinkedList[T]) Get(r int) (element T) {
	// 初始化扫描器
	var i int = 0
	currentNode := this.header.next
	for i < r {
		currentNode = currentNode.next
		i++
	}
	return currentNode.data
}

// 插入前驱节点，将e作为p的前驱插入，返回新节点的地址
func (this *DoublyLinkedList[T]) InsertBefore(p *Node[T], e T) *Node[T] {
	this.size++
	return p.insertPre(e)
}

// 插入后继节点，将e作为p的后继插入，返回新节点的地址
func (this *DoublyLinkedList[T]) InsertAfter(p *Node[T], e T) *Node[T] {
	this.size++
	return p.insertNext(e)
}

// 查找, 在节点p之前（不包括p本身）的n个节点中，从p向前查找元素e，返回第一个包含元素e的节点的地址
func (this *DoublyLinkedList[T]) Find(e T, n int, p *Node[T]) *Node[T] {
	for 0 < n {
		p = p.pre
		if e == p.data {
			return p
		}
		n--
	}
	return nil
}
