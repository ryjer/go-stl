package doublyLinkedlist

import (
	"fmt"

	num "gitee.com/ryjer/go-generic/number"
)

// 双向链表结构体定义，其中默认有两个哨兵节点：头节点header 和 尾节点trailer
type List[T num.Q] struct {
	size            int
	header, trailer *Node[T] // 头节点，尾节点。作为[首, 末] 节点的哨兵
}

// 初始化，创建空列表
func Init[T num.Q]() *List[T] {
	// 构造头尾哨兵节点
	headerNode := NewNode[T](0)
	trailerNode := NewNode[T](0)
	// 连接两哨兵
	headerNode.pre = nil
	headerNode.next = trailerNode
	trailerNode.pre = headerNode
	trailerNode.next = nil
	// 构造权柄
	return &List[T]{
		size:    0,
		header:  headerNode,
		trailer: trailerNode,
	}
}

// 构造方法，返回空列表
func New[T num.Q]() *List[T] {
	return Init[T]()
}

// 从切片构造
func NewFromSlice[T num.Q](sourceSlice []T) *List[T] {
	newList := Init[T]()            //初始化空链表
	for _, v := range sourceSlice { // 逐个复制元素
		newList.InsertAsLast(v)
	}
	return newList
}

// 从已有列表复制构造，从自节点p起（包含p）的n个节点复制到新列表中
func NewFormList[T num.Q](p *Node[T], n int) (newList *List[T]) {
	newList = Init[T]()
	for 0 < n {
		newList.InsertAsLast(p.data) // 新列表插入一个新末节点
		p = p.next                   // 当前节点p后移一个
		n--                          // n计数器递减
	}
	return newList
}

// 转换为一个新的切片返回
func (this *List[T]) ToSlice() (newSlice []T) {
	newSlice = make([]T, this.size) // 构造新切片
	// 逐个填入数据
	for i, currentNode := 0, this.FirstNode(); i < len(newSlice) && currentNode != this.trailer; {
		newSlice[i] = currentNode.data
		// 后移索引
		i++
		currentNode = currentNode.next
	}
	return newSlice
}

// 内部方法：获取首节点
func (this *List[T]) FirstNode() *Node[T] {
	return this.header.next // 头哨兵的后继
}

// 内部方法：获取末节点
func (this *List[T]) LastNode() *Node[T] {
	return this.trailer.pre // 尾哨兵的前驱
}

// 清空链表
func (this *List[T]) Clear() {
	this.size = 0
	this.header.next = this.trailer
	this.trailer.pre = this.header
}

// 容量
func (this *List[T]) Size() (size int) {
	return this.size
}

// 空判断
func (this *List[T]) IsEmpty() (isEmpty bool) {
	if this.size == 0 {
		isEmpty = true
	} else {
		isEmpty = false
	}
	return
}

// 读取元素，寻秩访问
func (this *List[T]) Get(r int) (element T) {
	// 初始化扫描器
	currentNode := this.FirstNode()
	for 0 < r {
		currentNode = currentNode.next
		r--
	}
	return currentNode.data
}

// 修改元素
func (this *List[T]) Put(p *Node[T], e T) {
	p.data = e
}

// 移除节点
func (this *List[T]) Remove(node *Node[T]) (element T) {
	element = node.data
	node.pre.next = node.next
	node.next.pre = node.pre
	this.size--
	return element
}

// 插入前驱节点，将e作为p的前驱插入，返回新节点的地址
func (this *List[T]) InsertBefore(p *Node[T], e T) *Node[T] {
	newNode := p.InsertAsPre(e) // 插入新前驱 newNode
	this.size++                 // 更新列表容量
	return newNode
}

// 插入后继节点，将e作为p的后继插入，返回新节点的地址
func (this *List[T]) InsertAfter(p *Node[T], e T) *Node[T] {
	newNode := p.InsertAsNext(e) // 插入新后继 newNode
	this.size++                  // 更新列表容量
	return newNode
}

// 作为首节点插入，将元素e作为整个列表的首节点插入
func (this *List[T]) InsertAsFirst(e T) *Node[T] {
	return this.InsertAfter(this.header, e) // 在头哨兵之后插入
}
func (this *List[T]) PushFront(e T) *Node[T] { //与InsertAsFirst相同，只有名字不同
	return this.InsertAfter(this.header, e) // 在头哨兵之后插入
}

// 弹出首节点
// 警告：不会检查链表是否为空，调用方需自行检查保证
func (this *List[T]) PopFront() (element T) {
	return this.Remove(this.FirstNode())
}

// 作为末节点插入，将元素e作为整个列表的末节点插入
func (this *List[T]) InsertAsLast(e T) *Node[T] {
	return this.InsertBefore(this.trailer, e) // 在尾哨兵之前插入
}
func (this *List[T]) PushBack(e T) *Node[T] { // 与 InsertAsLast相同，只有名字不同
	return this.InsertBefore(this.trailer, e) // 在尾哨兵之前插入
}

// 弹出末节点
// 警告：不会检查链表是否为空，调用方需自行检查保证
func (this *List[T]) PopBack() (element T) {
	return this.Remove(this.LastNode())
}

// 向前查找, 在节点p之前（不包括p本身）的n个节点中，从p向前查找元素e，返回第一个包含元素e的节点的地址
func (this *List[T]) FindBefore(e T, n int, p *Node[T]) (targetNode *Node[T]) {
	for 0 < n {
		p = p.pre
		if p.data == e {
			return p
		}
		n--
	}
	return nil
}

// 向后查找, 在节点p之后（不包括p本身）的n个节点中，从p向后查找元素e，返回第一个包含元素e的节点的地址
func (this *List[T]) FindAfter(e T, p *Node[T], n int) (targetNode *Node[T]) {
	for 0 < n {
		p = p.next
		if e == p.data {
			return p
		}
		n--
	}
	return nil
}

// 向前搜索，在节点p之前（不包括p）的n个节点中，向前逐个查找元素e，返回第一个不大于e的节点
// 由于数据结构本身的限制，有序列表的搜索并不会比无序列表更好
func (this *List[T]) SearchBefore(e T, n int, p *Node[T]) (targetNode *Node[T]) {
	for 0 <= n {
		p = p.pre
		if p.data <= e {
			break
		}
		n--
	}
	return p
}

// 唯一化，不要求列表元素有序
func (this *List[T]) Deduplicate() (removedNumber int) {
	if this.size < 2 { // 平凡链表(只有1个元素)，自然无重复
		return 0
	}
	// 非平凡链表
	oldSize := this.size                                                  //原规模
	for p, r := this.FirstNode().next, 1; p != this.trailer; p = p.next { // 从第2个节点开始
		// p为当前待比较节点，r为p节点之前待比较的节点数
		if q := this.FindBefore(p.data, r, p); q != nil {
			this.Remove(q) // 发现重复，删除前驱中的重复节点q，后方所有节点前移1位
		} else {
			r++ // 没有重复节点，待比较部分后移1位
		}
	}
	return (oldSize - this.size)
}

// 有序列表唯一化
func (this *List[T]) Uniquify() (removedNumber int) {
	if this.size < 2 { // 平凡链表，自然无重复
		return 0
	}
	// 非平凡链表
	oldSize := this.size
	for p, q := this.FirstNode(), this.FirstNode().next; this.trailer != q; {
		// 反复考察邻近节点对 (p, q)，q定义位p的直接后继
		if p.data != q.data {
			p = q // 互异，p移动到新区段的起始
		} else {
			this.Remove(q) // 否则雷同，删除后者q
		}
		q = p.next
	}
	return (oldSize - this.size)
}

// 选择排序，对链表中自p开始(包括p)的连续n个元素([p, p+n)区间)进行非降序排序
func (this *List[T]) SelectionSort(p *Node[T], n int) {
	head := p.pre
	tail := p                // 待排序区间：(head,tail)，有序区间[tail,p+n)
	for i := 0; i < n; i++ { // tail后移到 p+n 上界位置
		tail = tail.next
	}
	for 1 < n { // 反复从（非平凡的）待排序区间内找出最大者，并移至有序区间前端
		// 交换节点内容，减少碎片(这里data仅为基础类型，因而交换节点内容比交换节点更块)
		maxNode := this.selectMaxAfter(head.next, n)
		maxNode.data, tail.pre.data = tail.pre.data, maxNode.data
		tail = tail.pre
		n--
	}
}

// 向后求最大内容节点，查找自p节点开始的n个节点（包括p）中元素最大的节点
func (this *List[T]) selectMaxAfter(p *Node[T], n int) (maxNode *Node[T]) {
	maxNode = p
	for currentNode := p; 1 < n; n-- {
		currentNode = currentNode.next
		if currentNode.data >= maxNode.data {
			maxNode = currentNode
		}
	}
	return maxNode
}

// 插入排序，对链表中自p开始（包括p）的连续n个元素做插入排序
func (this *List[T]) InsertionSort(p *Node[T], n int) {
	for r := 0; r < n; r++ { // r为 有序前缀 节点数：有序部分 [header, header+r]；无序部分[]
		beforeNode := this.SearchBefore(p.data, r, p) // 查找插入位点
		p = p.next                                    //待比较节点指针后移
		(p.pre).MoveToAfter(beforeNode)               //将原待比较节点插入有序区间对应位置
	}
}

// 值判等，以值相等原则进行比较
// 定义：一个列表在"内容视图"上的相等包括：容量、链表的元素序列 相等，忽略其中的指针
func (this *List[T]) DeepEqual(another *List[T]) (equal bool) {
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
func (this *List[T]) String() (retString string) {
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
