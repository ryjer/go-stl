package twoWayNode

import (
	"testing"

	num "gitee.com/ryjer/go-generic/number"
)

// 值判等
func Test_deepEquals(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string   // 测试用例名称
		Receiver *Node[T] // 接收对象
		arg      *Node[T] // 参数
		want     bool     // 期望结果
	}
	// int 类型测试
	preNode := &Node[int]{
		data: 1,
		pre:  nil,
		next: nil,
	}
	nextNode := &Node[int]{
		data: 8,
		pre:  nil,
		next: nil,
	}
	intTests := []testCase[int]{
		{"int  nil相等", &Node[int]{8, nil, nil}, &Node[int]{8, nil, nil}, true},
		{"int 随机测试", &Node[int]{10, preNode, nextNode}, &Node[int]{10, preNode, nextNode}, true},
		{"int data测试", &Node[int]{1, preNode, nextNode}, &Node[int]{10, preNode, nextNode}, false},
		{"int  pre测试", &Node[int]{1, preNode, nextNode}, &Node[int]{1, nil, nextNode}, false},
		{"int  pre测试", &Node[int]{1, preNode, nextNode}, &Node[int]{1, nextNode, nextNode}, false},
		{"int next测试", &Node[int]{1, preNode, nextNode}, &Node[int]{1, preNode, nil}, false},
		{"int next测试", &Node[int]{1, preNode, nextNode}, &Node[int]{1, preNode, preNode}, false},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.DeepEqual(tt.arg); got != tt.want {
				t.Errorf("this.valueEquals(%v) = %v, want %v %v", tt.arg, got, tt.want, tt.Receiver)
			}
		})
	}
}

// 节点构造
func Test_New(t *testing.T) {
	type testCase[T num.Q] struct {
		name string
		arg  T
		want *Node[T]
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 0测试", 0, &Node[int]{0, nil, nil}},
		{"int 随机测试", 99, &Node[int]{99, nil, nil}},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.arg); !got.DeepEqual(tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
	// uint 类型测试
	uintTests := []testCase[uint]{
		{"int 0测试", 0, &Node[uint]{0, nil, nil}},
		{"int 随机测试", 99, &Node[uint]{99, nil, nil}},
	}
	for _, tt := range uintTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.arg); !got.DeepEqual(tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
	// float64 类型测试
	float64Tests := []testCase[float64]{
		{"int 0测试", 0, &Node[float64]{0, nil, nil}},
		{"int 随机测试", 99, &Node[float64]{99, nil, nil}},
	}
	for _, tt := range float64Tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.arg); !got.DeepEqual(tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 节点完全构造
func Test_FullNew(t *testing.T) {
	type args[T num.Q] struct {
		e    T
		pre  *Node[T]
		next *Node[T]
	}
	type testCase[T num.Q] struct {
		name string
		args args[T]
		want *Node[T]
	}
	// int 类型测试
	preNode := &Node[int]{
		data: 1,
		pre:  nil,
		next: nil,
	}
	nextNode := &Node[int]{
		data: 8,
		pre:  nil,
		next: nil,
	}
	intTests := []testCase[int]{
		{"int 空指针测试", args[int]{8, nil, nil}, &Node[int]{8, nil, nil}},
		{"int 空前驱测试", args[int]{1, nil, nextNode}, &Node[int]{1, nil, nextNode}},
		{"int 空后继测试", args[int]{1, preNode, nil}, &Node[int]{1, preNode, nil}},
		{"int 满指针测试", args[int]{1, preNode, nextNode}, &Node[int]{1, preNode, nextNode}},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FullNew(tt.args.e, tt.args.pre, tt.args.next); !got.DeepEqual(tt.want) {
				t.Errorf("FullNew() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 获取元素值
func Test_Get(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *Node[T]
		want     T
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 0测试", &Node[int]{0, nil, nil}, 0},
		{"int 随机测试", &Node[int]{8, nil, nil}, 8},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.Get(); got != tt.want {
				t.Errorf("Receiver.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 获取前驱节点
func Test_PreNode(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *Node[T]
		want     *Node[T]
	}
	// int 类型测试
	// 构造一个简单的 双向链表
	headNode := New(0)
	tailNode := New(8)
	headNode.next = tailNode // 将两个节点互相连接
	tailNode.pre = headNode
	intTests := []testCase[int]{
		{"int 简单链表测试", tailNode, headNode},
		{"int 构造链表测试", FullNew[int](2, tailNode, nil), tailNode},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.PreNode(); got != tt.want {
				t.Errorf("Receiver.PreNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 获取后继节点
func Test_NextNode(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *Node[T]
		want     *Node[T]
	}
	// int 类型测试
	// 构造一个简单的 双向链表
	headNode := New(0)
	tailNode := New(8)
	headNode.next = tailNode // 将两个节点互相连接
	tailNode.pre = headNode
	intTests := []testCase[int]{
		{"int 简单链表测试", headNode, tailNode},
		{"int 构造链表测试", FullNew[int](2, nil, headNode), headNode},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.NextNode(); got != tt.want {
				t.Errorf("Receiver.NextNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 作为前驱节点插入
func Test_insertAsPre(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *Node[T]
		arg      T
		want     *Node[T]
	}
	// int 类型测试
	// 构造一个简单的 双向链表
	headNode := New(0)
	tailNode := New(8)
	headNode.next = tailNode // 将两个节点互相连接
	tailNode.pre = headNode
	intTests := []testCase[int]{
		{"int 插入测试", tailNode, 1, &Node[int]{1, headNode, tailNode}}, // 只支持一个测试用例
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.insertAsPre(tt.arg); !((headNode.next == got && got.pre == headNode) && (got.next == tailNode && tailNode.pre == got)) {
				t.Errorf("this.insertAsPre() = %v, Receiver => %v, want %v", got, tt.Receiver, tt.want)
			}
		})
	}
}

// 作为后继节点插入
func Test_insertAsNext(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *Node[T]
		arg      T
		want     *Node[T]
	}
	// int 类型测试
	// 构造一个简单的 双向链表
	headNode := New(0)
	tailNode := New(8)
	headNode.next = tailNode // 将两个节点互相连接
	tailNode.pre = headNode
	intTests := []testCase[int]{
		{"int 插入测试", headNode, 1, &Node[int]{1, headNode, tailNode}}, // 只支持一个测试用例
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.insertAsNext(tt.arg); !((headNode.next == got && got.pre == headNode) && (got.next == tailNode && tailNode.pre == got)) {
				t.Errorf("this.insertAsNext() = %v, Receiver => %v, want %v", got, tt.Receiver, tt.want)
			}
		})
	}
}
