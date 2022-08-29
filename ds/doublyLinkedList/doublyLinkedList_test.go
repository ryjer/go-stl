package doublyLinkedlist

import (
	"testing"

	num "gitee.com/ryjer/go-generic/number"
)

// 内容视图判等
func Test_DeepEqual(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *DoublyLinkedList[T]
		arg      *DoublyLinkedList[T]
		want     bool
	}
	// int 类型测试
	intList := New[int]()
	// 在列表尾部插入元素，构造链表
	intList.trailer.InsertAsPre(0)
	intList.trailer.InsertAsPre(1)
	intList.trailer.InsertAsPre(2)
	intList.trailer.InsertAsPre(3)
	intList.size = 4
	intTests := []testCase[int]{
		{"int 空判等", New[int](), NewFromSlice[int]([]int{}), true},
		{"int 非空判等", intList, NewFromSlice[int]([]int{0, 1, 2, 3}), true},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.DeepEqual(tt.arg); got != tt.want {
				t.Errorf("Receiver.DeepEqual(%v) = %v, Receiver => %v, want %v", tt.arg, got, tt.Receiver.String(), tt.want)
			}
		})
	}
}

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

// 从切片构造
func Test_NewFromSlice(t *testing.T) {
	type testCase[T num.Q] struct {
		name string
		arg  []T
		want string
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 空链表", []int{}, "{0 []}"},
		{"int 非空链表", []int{0, 1, 2, 3}, "{4 [0 1 2 3]}"},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromSlice(tt.arg); got.String() != tt.want {
				t.Errorf("NewFromSlice(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

// // 通过复制构造
// func Test_NewFormList(t testing.T) {
// 	type args[T num.Q] struct {
// 		p *Node[T]
// 		n int
// 	}
// 	type testCase[T num.Q] struct {
// 		name     string               // 测试用例名
// 		Receiver *DoublyLinkedList[T] // 接收对象
// 		args     args[T]              // 多参数
// 		want     T                    // 目标节点中的元素值
// 	}
// 	// int 类型测试，只插入一个节点
// 	// 构造链表
// 	intList := New[int]()
// 	node0 := intList.trailer.InsertAsPre(0)
// 	intList.trailer.InsertAsPre(1)
// 	node2 := intList.trailer.InsertAsPre(2)
// 	intList.trailer.InsertAsPre(3)
// 	intList.trailer.InsertAsPre(4)
// 	node5 := intList.trailer.InsertAsPre(5)
// 	intList.size = 6
// 	intTests := []testCase[int]{
// 		{"int 全复制", intList, args[int]{node0, -1}, -1},
// 		{"int 上界复制", intList, args[int]{node5, -1}, -1},
// 		{"int 下界复制", intList, args[int]{node3, -1}, -1},
// 		{"int 中间复制", intList, args[int]{node3, -1}, -1},
// 	}
// 	for _, tt := range intTests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// 插入节点，进行测试
// 			var got *Node[int]
// 			if got = tt.Receiver.NewFormList(tt.args.p, tt.args.n); !((got.data == tt.want) && (got.pre.next == got && got.next.pre == got) && (intList.size == 7)) {
// 				t.Errorf("this.NewFormList(%v, %v).data = %v, Receiver => %v, want %v", tt.args.p, tt.args.n, got.data, intList.String(), tt.want)
// 			}
// 			// 复原链表，保持每个测试用例环境一致性
// 			intList.size--
// 			got.pre.next = got.next
// 			got.next.pre = got.pre
// 		})
// 	}
// }

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
	// 在列表尾部插入元素，构造链表
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

// 作为前驱插入
func Test_InsertBefore(t *testing.T) {
	type args[T num.Q] struct {
		p *Node[T]
		e T
	}
	type testCase[T num.Q] struct {
		name     string               // 测试用例名
		Receiver *DoublyLinkedList[T] // 接收对象
		args     args[T]              // 多参数
		want     T                    // 目标节点中的元素值
	}
	// int 类型测试，只插入一个节点
	// 构造链表
	intList := New[int]()
	node0 := intList.trailer.InsertAsPre(0)
	intList.trailer.InsertAsPre(1)
	intList.trailer.InsertAsPre(2)
	node3 := intList.trailer.InsertAsPre(3)
	intList.trailer.InsertAsPre(4)
	node5 := intList.trailer.InsertAsPre(5)
	intList.size = 6
	intTests := []testCase[int]{
		{"int 首节点之前插入", intList, args[int]{node0, -1}, -1},
		{"int 末节点之前插入", intList, args[int]{node5, -1}, -1},
		{"int 中间随机插入", intList, args[int]{node3, -1}, -1},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			// 插入节点，进行测试
			var got *Node[int]
			if got = tt.Receiver.InsertBefore(tt.args.p, tt.args.e); !((got.data == tt.want) && (got.pre.next == got && got.next.pre == got) && (intList.size == 7)) {
				t.Errorf("this.InsertBefore(%v, %v).data = %v, Receiver => %v, want %v", tt.args.p, tt.args.e, got.data, intList.String(), tt.want)
			}
			// 复原链表，保持每个测试用例环境一致性
			intList.size--
			got.pre.next = got.next
			got.next.pre = got.pre
		})
	}
}

// 作为前驱插入
func Test_InsertAfter(t *testing.T) {
	type args[T num.Q] struct {
		p *Node[T]
		e T
	}
	type testCase[T num.Q] struct {
		name     string               // 测试用例名
		Receiver *DoublyLinkedList[T] // 接收对象
		args     args[T]              // 多参数
		want     T                    // 目标节点中的元素值
	}
	// int 类型测试，只插入一个节点
	// 构造链表
	intList := New[int]()
	node0 := intList.trailer.InsertAsPre(0)
	intList.trailer.InsertAsPre(1)
	intList.trailer.InsertAsPre(2)
	node3 := intList.trailer.InsertAsPre(3)
	intList.trailer.InsertAsPre(4)
	node5 := intList.trailer.InsertAsPre(5)
	intList.size = 6
	intTests := []testCase[int]{
		{"int 首节点之后插入", intList, args[int]{node0, -1}, -1},
		{"int 末节点之后插入", intList, args[int]{node5, -1}, -1},
		{"int 中间随机插入", intList, args[int]{node3, -1}, -1},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			// 插入节点，进行测试
			var got *Node[int]
			if got = tt.Receiver.InsertAfter(tt.args.p, tt.args.e); !((got.data == tt.want) && (got.pre.next == got && got.next.pre == got) && (intList.size == 7)) {
				t.Errorf("this.InsertAfter(%v, %v).data = %v, Receiver => %v, want %v", tt.args.p, tt.args.e, got.data, intList.String(), tt.want)
			}
			// 复原链表，保持每个测试用例环境一致性
			intList.size--
			got.pre.next = got.next
			got.next.pre = got.pre
		})
	}
}

// 作为首节点插入
func Test_InsertAsFirst(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string               // 测试用例名
		Receiver *DoublyLinkedList[T] // 接收对象
		arg      T                    // 多参数
	}
	// int 类型测试，只插入一个节点
	// 构造链表
	intList := New[int]()
	intList.trailer.InsertAsPre(0)
	intList.trailer.InsertAsPre(1)
	intList.trailer.InsertAsPre(2)
	intList.trailer.InsertAsPre(3)
	intList.trailer.InsertAsPre(4)
	intList.trailer.InsertAsPre(5)
	intList.size = 6
	intTests := []testCase[int]{
		{"int 插入一个数", intList, 0},
		{"int 插入另一个数", intList, -1},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			oldFirst := tt.Receiver.header.next // 记录旧首节点
			// 插入节点，进行测试
			if got := tt.Receiver.InsertAsFirst(tt.arg); !((got.data == tt.arg) && ((got.pre == tt.Receiver.header && tt.Receiver.header.next == got) && (got.next == oldFirst && oldFirst.pre == got)) && (intList.size == 7)) {
				t.Errorf("this.InsertAsFirst(%v).data = %v, Receiver => %v, want %v", tt.arg, got.data, intList.String(), got.data)
			}
			// 复原链表，保持每个测试用例环境一致性
			tt.Receiver.header.next = oldFirst
			oldFirst.pre = tt.Receiver.header
			intList.size--
		})
	}
}

// 作为末节点插入
func Test_InsertAsLast(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string               // 测试用例名
		Receiver *DoublyLinkedList[T] // 接收对象
		arg      T                    // 多参数
	}
	// int 类型测试，只插入一个节点
	// 构造链表
	intList := New[int]()
	intList.trailer.InsertAsPre(0)
	intList.trailer.InsertAsPre(1)
	intList.trailer.InsertAsPre(2)
	intList.trailer.InsertAsPre(3)
	intList.trailer.InsertAsPre(4)
	intList.trailer.InsertAsPre(5)
	intList.size = 6
	intTests := []testCase[int]{
		{"int 插入一个数", intList, 0},
		{"int 插入另一个数", intList, -1},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			oldLast := tt.Receiver.trailer.pre // 记录旧首节点
			// 插入节点，进行测试
			if got := tt.Receiver.InsertAsLast(tt.arg); !((got.data == tt.arg) && ((got.pre == oldLast && oldLast.next == got) && (got.next == tt.Receiver.trailer && tt.Receiver.trailer.pre == got)) && (intList.size == 7)) {
				t.Errorf("this.InsertAsLast(%v).data = %v, Receiver => %v, want %v", tt.arg, got.data, intList.String(), got.data)
			}
			// 复原链表，保持每个测试用例环境一致性
			oldLast.next = tt.Receiver.trailer
			tt.Receiver.trailer.pre = oldLast
			intList.size--
		})
	}
}

// 向前查找
func Test_FindBefore(t *testing.T) {
	type args[T num.Q] struct {
		e T
		n int
		p *Node[T]
	}
	type testCase[T num.Q] struct {
		name     string               // 测试用例名
		Receiver *DoublyLinkedList[T] // 接收对象
		args     args[T]              // 多参数
		want     T                    // 目标节点中的元素值
	}
	// int 类型测试
	// 构造链表
	intList := New[int]()
	intList.trailer.InsertAsPre(0)
	intList.trailer.InsertAsPre(1)
	intList.trailer.InsertAsPre(2)
	intList.trailer.InsertAsPre(3)
	intList.trailer.InsertAsPre(4)
	intList.trailer.InsertAsPre(5)
	intList.size = 6
	intTests := []testCase[int]{
		{"int 全查找，循环下界", intList, args[int]{4, 5, intList.lastNode()}, 4},
		{"int 全查找，中间随机", intList, args[int]{2, 5, intList.lastNode()}, 2},
		{"int 全查找，循环上界", intList, args[int]{0, 5, intList.lastNode()}, 0},
		{"int 部分查找，循环下界，", intList, args[int]{3, 3, intList.lastNode().pre}, 3},
		{"int 部分查找，循环上界，", intList, args[int]{1, 3, intList.lastNode().pre}, 1},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.FindBefore(tt.args.e, tt.args.n, tt.args.p); got.data != tt.want {
				t.Errorf("this.FindBefore(%v, %v, %v).data = %v, want %v", tt.args.e, tt.args.n, tt.args.p.data, got.data, tt.want)
			}
		})
	}
}

// 向后查找
func Test_FindAfter(t *testing.T) {
	type args[T num.Q] struct {
		e T
		p *Node[T]
		n int
	}
	type testCase[T num.Q] struct {
		name     string               // 测试用例名
		Receiver *DoublyLinkedList[T] // 接收对象
		args     args[T]              // 多参数
		want     T                    // 目标节点中的元素值
	}
	// int 类型测试
	// 构造链表
	intList := New[int]()
	intList.trailer.InsertAsPre(0)
	intList.trailer.InsertAsPre(1)
	intList.trailer.InsertAsPre(2)
	intList.trailer.InsertAsPre(3)
	intList.trailer.InsertAsPre(4)
	intList.trailer.InsertAsPre(5)
	intList.size = 6
	intTests := []testCase[int]{
		{"int 全查找，循环下界", intList, args[int]{1, intList.firstNode(), 5}, 1},
		{"int 全查找，中间随机", intList, args[int]{2, intList.firstNode(), 5}, 2},
		{"int 全查找，循环上界", intList, args[int]{5, intList.firstNode(), 5}, 5},
		{"int 部分查找，循环下界，", intList, args[int]{2, intList.firstNode().next, 3}, 2},
		{"int 部分查找，循环上界，", intList, args[int]{4, intList.firstNode().next, 3}, 4},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.FindAfter(tt.args.e, tt.args.p, tt.args.n); got.data != tt.want {
				t.Errorf("this.FindAfter(%v, %v, %v).data = %v, want %v", tt.args.e, tt.args.p, tt.args.n, got.data, tt.want)
			}
		})
	}
}
