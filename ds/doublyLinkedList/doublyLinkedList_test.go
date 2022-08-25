package doublyLinkedlist

import (
	"testing"

	num "gitee.com/ryjer/go-generic/number"
)

// func Test_deepEqual(t *testing.T) {
// 	type testCase[T num.Q] struct {
// 		name     string               // 测试用例名称
// 		Receiver *DoublyLinkedList[T] // 接收对象
// 		arg      *DoublyLinkedList[T] // 参数
// 		want     bool                 // 期望结果
// 	}
// 	// int 类型测试
// 	preNode := &Node[int]{
// 		data: 1,
// 		pre:  nil,
// 		next: nil,
// 	}
// 	nextNode := &Node[int]{
// 		data: 8,
// 		pre:  nil,
// 		next: nil,
// 	}
// 	intTests := []testCase[int]{
// 		{"int  nil相等", New(), New(), true},
// 		{"int 随机测试", &Node[int]{10, preNode, nextNode}, &Node[int]{10, preNode, nextNode}, true},
// 		{"int data测试", &Node[int]{1, preNode, nextNode}, &Node[int]{10, preNode, nextNode}, false},
// 		{"int  pre测试", &Node[int]{1, preNode, nextNode}, &Node[int]{1, nil, nextNode}, false},
// 		{"int  pre测试", &Node[int]{1, preNode, nextNode}, &Node[int]{1, nextNode, nextNode}, false},
// 		{"int next测试", &Node[int]{1, preNode, nextNode}, &Node[int]{1, preNode, nil}, false},
// 		{"int next测试", &Node[int]{1, preNode, nextNode}, &Node[int]{1, preNode, preNode}, false},
// 	}
// 	for _, tt := range intTests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.Receiver.deepEquals(tt.arg); got != tt.want {
// 				t.Errorf("this.valueEquals(%v) = %v, want %v %v", tt.arg, got, tt.want, tt.Receiver)
// 			}
// 		})
// 	}
// }

// 内部初始化函数
func Test_Init(t *testing.T) {
	t.Run("空列表初始化", func(t *testing.T) {
		if got := Init[int](); !(got.size == 0) && (got.header.pre == nil && got.header.next == got.trailer) && (got.trailer.pre == got.header && got.trailer.next == nil) {
			t.Errorf("Init() = %v, want %v", got, got.String())
		}
	})
}

// 空链表构造器
func Test_New(t *testing.T) {
	t.Run("空列表初始化", func(t *testing.T) {
		if got := New[int](); !(got.size == 0) && (got.header.pre == nil && got.header.next == got.trailer) && (got.trailer.pre == got.header && got.trailer.next == nil) {
			t.Errorf("New() = %v, want %v", got, got.String())
		}
	})
}

// 寻秩读取元素
func Test_Get(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *DoublyLinkedList[T]
		arg      int
		want     T
	}
	// int 类型测试
	intList := New[int]()
	// 在列表尾部插入元素，构造测试对象
	intList.trailer.InsertAsPre(0)
	intList.trailer.InsertAsPre(1)
	intList.trailer.InsertAsPre(2)
	intList.trailer.InsertAsPre(3)
	intList.size = 4
	intTests := []testCase[int]{
		{"int 上界读取", intList, 0, 0},
		{"int 随机测试", intList, 1, 1},
		{"int 下界测试", intList, 3, 3},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.Get(tt.arg); got != tt.want {
				t.Errorf("Receiver.Get(%v) = %v, Receiver => %v, want %v", tt.arg, got, tt.Receiver.String(), tt.want)
			}
		})
	}
}
