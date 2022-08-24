package doublyLinkedlist

import (
	"testing"

	num "gitee.com/ryjer/go-generic/number"
)

func Test_deepEqual(t *testing.T) {
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
		{"int  nil相等", &Node[int]{0, nil, nil}, &Node[int]{0, nil, nil}, true},
		{"int 随机测试", &Node[int]{10, preNode, nextNode}, &Node[int]{10, preNode, nextNode}, true},
		{"int data测试", &Node[int]{1, preNode, nextNode}, &Node[int]{10, preNode, nextNode}, false},
		{"int  pre测试", &Node[int]{1, preNode, nextNode}, &Node[int]{1, nextNode, nextNode}, false},
		{"int  pre测试", &Node[int]{1, preNode, nextNode}, &Node[int]{1, nil, nextNode}, false},
		{"int next测试", &Node[int]{1, preNode, nextNode}, &Node[int]{1, preNode, preNode}, false},
		{"int next测试", &Node[int]{1, preNode, nextNode}, &Node[int]{1, preNode, nil}, false},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.deepEqual(tt.arg); got != tt.want {
				t.Errorf("this.deepEqual(%v) = %v, want %v %v", tt.arg, got, tt.want, tt.Receiver)
			}
		})
	}
}

func Test_NewNode(t *testing.T) {
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
			if got := NewNode(tt.arg); !got.deepEqual(tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
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
			if got := NewNode(tt.arg); !got.deepEqual(tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
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
			if got := NewNode(tt.arg); !got.deepEqual(tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			if got := FullNew(tt.args.e, tt.args.pre, tt.args.next); !got.deepEqual(tt.want) {
				t.Errorf("FullNew() = %v, want %v", got, tt.want)
			}
		})
	}
}
