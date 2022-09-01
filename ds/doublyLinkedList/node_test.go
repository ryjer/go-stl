package doublyLinkedlist

import (
	"testing"

	num "gitee.com/ryjer/go-generic/number"
)

// 值判等
func Test_node_DeepEquals(t *testing.T) {
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
func Test_node_NewNode(t *testing.T) {
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
			if got := NewNode(tt.arg); !got.DeepEqual(tt.want) {
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
			if got := NewNode(tt.arg); !got.DeepEqual(tt.want) {
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
			if got := NewNode(tt.arg); !got.DeepEqual(tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 节点完全构造
func Test_node_FullNewNode(t *testing.T) {
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
			if got := FullNewNode(tt.args.e, tt.args.pre, tt.args.next); !got.DeepEqual(tt.want) {
				t.Errorf("FullNew() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 获取元素值
func Test_node_Get(t *testing.T) {
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

// 修改元素
func Test_node_Put(t *testing.T) {
	type want[T num.Q] struct {
		ret     T
		newData T
	}
	type testCase[T num.Q] struct {
		name     string
		Receiver *Node[T]
		arg      T
		want     want[T]
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 测试", &Node[int]{1, nil, nil}, 1, want[int]{1, 1}},
		{"int 随机测试", &Node[int]{0, nil, nil}, 8, want[int]{0, 8}},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.Put(tt.arg); !(got == tt.want.ret && tt.Receiver.data == tt.want.newData) {
				t.Errorf("Receiver.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 获取前驱节点
func Test_node_PreNode(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *Node[T]
		want     *Node[T]
	}
	// int 类型测试
	// 构造一个简单的 双向链表
	headNode := NewNode(0)
	tailNode := NewNode(8)
	headNode.next = tailNode // 将两个节点互相连接
	tailNode.pre = headNode
	intTests := []testCase[int]{
		{"int 简单链表测试", tailNode, headNode},
		{"int 构造链表测试", FullNewNode(2, tailNode, nil), tailNode},
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
func Test_node_NextNode(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *Node[T]
		want     *Node[T]
	}
	// int 类型测试
	// 构造一个简单的 双向链表
	headNode := NewNode(0)
	tailNode := NewNode(8)
	headNode.next = tailNode // 将两个节点互相连接
	tailNode.pre = headNode
	intTests := []testCase[int]{
		{"int 简单链表测试", headNode, tailNode},
		{"int 构造链表测试", FullNewNode(2, nil, headNode), headNode},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.NextNode(); got != tt.want {
				t.Errorf("Receiver.NextNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 元素作为前驱插入
func Test_node_InsertAsPre(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *Node[T]
		arg      T
		want     *Node[T]
	}
	// int 类型测试
	// 构造一个简单的 双向链表
	headNode := NewNode(0)
	tailNode := NewNode(8)
	headNode.next = tailNode // 将两个节点互相连接
	tailNode.pre = headNode
	intTests := []testCase[int]{
		{"int 插入测试", tailNode, 1, &Node[int]{1, headNode, tailNode}}, // 在 head <-> tail 中插入一个节点，变为 head(0) <-> new(1) <-> tail(8)
		{"int 无前驱哨兵", headNode, -2, &Node[int]{-2, nil, headNode}},   // 在 head 前插入一个节点，变为 new(-2) <-> head(0) <-> (1) <->tail(8)
	}
	// 有前驱插入
	t.Run(intTests[0].name, func(t *testing.T) {
		if got := intTests[0].Receiver.InsertAsPre(intTests[0].arg); !((headNode.next == got && tailNode.pre == got) && got.DeepEqual(intTests[0].want)) {
			t.Errorf("this.InsertAsPre() = %v, Receiver => %v, want %v", got, intTests[0].Receiver, intTests[0].want)
		}
	})
	// 提示，此时链表已变为 head(0) <-> (1) <-> tail(8)
	// 无前驱插入
	t.Run(intTests[1].name, func(t *testing.T) {
		if got := intTests[1].Receiver.InsertAsPre(intTests[1].arg); !(headNode.pre == got && got.DeepEqual(intTests[1].want)) {
			t.Errorf("this.InsertAsPre() = %v, Receiver => %v, want %v", got, intTests[1].Receiver, intTests[1].want)
		}
	})
}

// 元素作为后继插入
func Test_node_InsertAsNext(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *Node[T]
		arg      T
		want     *Node[T]
	}
	// int 类型测试
	// 构造一个简单的 双向链表
	headNode := NewNode(0)
	tailNode := NewNode(8)
	headNode.next = tailNode // 将两个节点互相连接
	tailNode.pre = headNode
	intTests := []testCase[int]{
		{"int 插入测试", headNode, 1, &Node[int]{1, headNode, tailNode}}, // 在 head <-> tail 中插入一个节点，变为： head(0) <-> new(1) <-> tail(8)
		{"int 无后继哨兵", tailNode, 16, &Node[int]{16, tailNode, nil}},   //  在 tail 后入一个节点，变为： head(0) <-> (1) <-> tail(8) <-> new(16)
	}
	// 有后继哨兵插入
	t.Run(intTests[0].name, func(t *testing.T) {
		if got := intTests[0].Receiver.InsertAsNext(intTests[0].arg); !((headNode.next == got && tailNode.pre == got) && got.DeepEqual(intTests[0].want)) {
			t.Errorf("this.InsertAsNext() = %v, Receiver => %v, want %v", got, intTests[0].Receiver, intTests[0].want)
		}
	})
	// 无后继哨兵插入
	t.Run(intTests[1].name, func(t *testing.T) {
		if got := intTests[1].Receiver.InsertAsNext(intTests[1].arg); !(tailNode.next == got && got.DeepEqual(intTests[1].want)) {
			t.Errorf("this.InsertAsNext() = %v, Receiver => %v, want %v", got, intTests[1].Receiver, intTests[1].want)
		}
	})
}

// 节点作为前驱插入
func Test_node_InsertNodeAsPre(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *Node[T]
		arg      *Node[T]
		want     *Node[T]
	}
	// int 类型测试
	// 构造一个简单的 双向链表
	headNode := NewNode(1)
	tailNode := NewNode(3)
	headNode.next = tailNode // 将两个节点互相连接
	tailNode.pre = headNode
	midNode := tailNode.InsertAsPre(2)
	intTests := []testCase[int]{
		{"int 作为前节点插入", midNode, &Node[int]{-1, nil, nil}, headNode},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.InsertNodeAsPre(tt.arg); !tt.arg.IsBetween(got, tt.Receiver) {
				t.Errorf("Receiver.InsertNodeAsPre(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

// 节点作为后继插入
func Test_node_InsertNodeAsNext(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *Node[T]
		arg      *Node[T]
		want     *Node[T]
	}
	// int 类型测试
	// 构造一个简单的 双向链表
	headNode := NewNode(1)
	tailNode := NewNode(3)
	headNode.next = tailNode // 将两个节点互相连接
	tailNode.pre = headNode
	midNode := tailNode.InsertAsPre(2)
	intTests := []testCase[int]{
		{"int 作为后节点插入", midNode, &Node[int]{-1, nil, nil}, headNode},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.InsertNodeAsNext(tt.arg); !tt.arg.IsBetween(tt.Receiver, got) {
				t.Errorf("Receiver.InsertNodeAsNext(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}
func Test_node_Remove(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *Node[T]
		want     T
	}
	// int 类型测试
	// 构造一个简单的 双向链表
	headNode := NewNode(0)
	tailNode := NewNode(8)
	headNode.next = tailNode // 将两个节点互相连接
	tailNode.pre = headNode
	intTests := []testCase[int]{
		{"int 中间态移除", FullNewNode(2, headNode, tailNode), 2},
		{"int 正常态移除", headNode.InsertAsNext(4), 4}, // 在 head节点后插入一个节点，形成  head(2) <-> new(4) <-> tail(8) 的结构
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.Remove(); !(headNode.next == tailNode && tailNode.pre == headNode) {
				t.Errorf("this.Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
