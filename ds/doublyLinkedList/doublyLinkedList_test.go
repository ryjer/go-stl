package doublyLinkedlist

import (
	"testing"

	num "gitee.com/ryjer/go-generic/number"
)

// 内容视图判等
func Test_DeepEqual(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *List[T]
		arg      *List[T]
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
		{"int 空判等", New[int](), NewFromSlice([]int{}), true},
		{"int 非空判等", intList, NewFromSlice([]int{0, 1, 2, 3}), true},
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

// 通过复制构造
func Test_NewFormList(t *testing.T) {
	type args[T num.Q] struct {
		p *Node[T]
		n int
	}
	type testCase[T num.Q] struct {
		name     string   // 测试用例名
		Receiver *List[T] // 接收对象
		args     args[T]  // 多参数
		want     *List[T] // 目标节点中的元素值
	}
	// int 类型测试，只插入一个节点
	// 构造链表
	intList := New[int]()
	node0 := intList.trailer.InsertAsPre(0)
	intList.trailer.InsertAsPre(1)
	node2 := intList.trailer.InsertAsPre(2)
	node3 := intList.trailer.InsertAsPre(3)
	intList.trailer.InsertAsPre(4)
	intList.trailer.InsertAsPre(5)
	intList.size = 6
	intTests := []testCase[int]{
		{"int 全复制", intList, args[int]{node0, 6}, NewFromSlice([]int{0, 1, 2, 3, 4, 5})},
		{"int 上界复制", intList, args[int]{node3, 3}, NewFromSlice([]int{3, 4, 5})},
		{"int 下界复制", intList, args[int]{node0, 3}, NewFromSlice([]int{0, 1, 2})},
		{"int 中间复制", intList, args[int]{node2, 2}, NewFromSlice([]int{2, 3})},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFormList(tt.args.p, tt.args.n); !got.DeepEqual(tt.want) {
				t.Errorf("NewFormList(%v, %v) = %v, want %v", tt.args.p, tt.args.n, got, tt.want)
			}
		})
	}
}

// 寻秩读取元素
func Test_Get(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *List[T]
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

// 清空链表
func Test_Clear(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *List[T]
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 空链表清空", NewFromSlice([]int{})},
		{"int 非空链表", NewFromSlice([]int{1, 2, 3, 4})},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.Receiver.Clear(); !(tt.Receiver.size == 0 && tt.Receiver.header.next == tt.Receiver.trailer && tt.Receiver.trailer.pre == tt.Receiver.header) {
				t.Errorf("Receiver.Clear(), Receiver => %v", tt.Receiver)
			}
		})
	}
}

// 获取已用容量
func Test_Size(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *List[T]
		want     int
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 空链表清空", NewFromSlice([]int{}), 0},
		{"int 非空链表", NewFromSlice([]int{1, 2, 3, 4}), 4},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.Size(); !(got == tt.want) {
				t.Errorf("Receiver.Size() = %v, Receiver => %v, want %v", got, tt.Receiver, tt.want)
			}
		})
	}
}

// 获取已用容量
func Test_IsEmpty(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *List[T]
		want     bool
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 空链表清空", NewFromSlice([]int{}), true},
		{"int 非空链表", NewFromSlice([]int{1, 2, 3, 4}), false},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.IsEmpty(); !(got == tt.want) {
				t.Errorf("Receiver.IsEmpty() = %v, Receiver => %v, want %v", got, tt.Receiver, tt.want)
			}
		})
	}
}

// 移除指定节点
func Test_Remove(t *testing.T) {
	type want[T num.Q] struct {
		Receiver *List[T] // 处理后的链表
		ret      T        // 返回值
	}
	type testCase[T num.Q] struct {
		name     string   // 测试用例名
		Receiver *List[T] // 接收对象
		arg      *Node[T] // 多参数
		want     want[T]  // 返回值
	}
	// int 类型测试，只插入一个节点
	// 构造链表
	intList := New[int]()
	node0 := intList.trailer.InsertAsPre(0)
	intList.trailer.InsertAsPre(1)
	node2 := intList.trailer.InsertAsPre(2)
	intList.trailer.InsertAsPre(3)
	intList.trailer.InsertAsPre(4)
	node5 := intList.trailer.InsertAsPre(5)
	intList.size = 6
	intTests := []testCase[int]{
		{"int 首节点移除", intList, node0, want[int]{NewFromSlice([]int{1, 2, 3, 4, 5}), 0}},
		{"int 末节点移除", intList, node5, want[int]{NewFromSlice([]int{0, 1, 2, 3, 4}), 5}},
		{"int 中间随机移除", intList, node2, want[int]{NewFromSlice([]int{0, 1, 3, 4, 5}), 2}},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			// 记录被删除节点的前一节点，以备之后复原
			previousNode := tt.arg.pre
			var got int
			if got = tt.Receiver.Remove(tt.arg); !(tt.Receiver.DeepEqual(tt.want.Receiver) && got == tt.want.ret) {
				t.Errorf("this.Remove(%v) = %v, Receiver => %v, want %v %v", tt.arg, got, tt.Receiver, tt.want.ret, tt.want.Receiver)
			}
			// 复原链表，保持每个测试用例环境一致性
			tt.Receiver.InsertAfter(previousNode, got)
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
		name     string   // 测试用例名
		Receiver *List[T] // 接收对象
		args     args[T]  // 多参数
		want     T        // 目标节点中的元素值
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
		name     string   // 测试用例名
		Receiver *List[T] // 接收对象
		args     args[T]  // 多参数
		want     T        // 目标节点中的元素值
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
		name     string   // 测试用例名
		Receiver *List[T] // 接收对象
		arg      T        // 多参数
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

// 弹出首节点
func Test_PopFront(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string   // 测试用例名
		Receiver *List[T] // 接收对象
		want     T
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 下界弹出", NewFromSlice([]int{4}), 4},
		{"int 随机长度弹出", NewFromSlice([]int{1, 2, 3, 4}), 1},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.PopFront(); !(got == tt.want) {
				t.Errorf("this.PopFront() = %v, Receiver => %v, want %v", got, tt.Receiver, tt.want)
			}
		})
	}
}

// 作为末节点插入
func Test_InsertAsLast(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string   // 测试用例名
		Receiver *List[T] // 接收对象
		arg      T        // 多参数
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

// 弹出末节点
func Test_PopBack(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string   // 测试用例名
		Receiver *List[T] // 接收对象
		want     T
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 下界弹出", NewFromSlice([]int{4}), 4},
		{"int 随机长度弹出", NewFromSlice([]int{1, 2, 3}), 3},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.PopBack(); !(got == tt.want) {
				t.Errorf("this.PopBack() = %v, Receiver => %v, want %v", got, tt.Receiver, tt.want)
			}
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
		name     string   // 测试用例名
		Receiver *List[T] // 接收对象
		args     args[T]  // 多参数
		want     T        // 目标节点中的元素值
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
		{"int 全查找，循环下界", intList, args[int]{4, 5, intList.LastNode()}, 4},
		{"int 全查找，中间随机", intList, args[int]{2, 5, intList.LastNode()}, 2},
		{"int 全查找，循环上界", intList, args[int]{0, 5, intList.LastNode()}, 0},
		{"int 部分查找，循环下界，", intList, args[int]{3, 3, intList.LastNode().pre}, 3},
		{"int 部分查找，循环上界，", intList, args[int]{1, 3, intList.LastNode().pre}, 1},
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
		name     string   // 测试用例名
		Receiver *List[T] // 接收对象
		args     args[T]  // 多参数
		want     T        // 目标节点中的元素值
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
		{"int 全查找，循环下界", intList, args[int]{1, intList.FirstNode(), 5}, 1},
		{"int 全查找，中间随机", intList, args[int]{2, intList.FirstNode(), 5}, 2},
		{"int 全查找，循环上界", intList, args[int]{5, intList.FirstNode(), 5}, 5},
		{"int 部分查找，循环下界，", intList, args[int]{2, intList.FirstNode().next, 3}, 2},
		{"int 部分查找，循环上界，", intList, args[int]{4, intList.FirstNode().next, 3}, 4},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.FindAfter(tt.args.e, tt.args.p, tt.args.n); got.data != tt.want {
				t.Errorf("this.FindAfter(%v, %v, %v).data = %v, want %v", tt.args.e, tt.args.p, tt.args.n, got.data, tt.want)
			}
		})
	}
}

// 无序列表去重
func Test_Deduplicate(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *List[T] //原对象
		newRecv  *List[T] //新对象
		ret      int      //返回值
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 空列表去重", NewFromSlice([]int{}), NewFromSlice([]int{}), 0},
		{"int 平凡列表", NewFromSlice([]int{1}), NewFromSlice([]int{1}), 0},
		{"int 非平凡列表", NewFromSlice([]int{4, 8, 6, 8, 2, 2, 9}), NewFromSlice([]int{4, 6, 8, 2, 9}), 2},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.Deduplicate(); !(tt.Receiver.DeepEqual(tt.newRecv) && got == tt.ret) {
				t.Errorf("Receiver.Deduplicate() = %v, Recv => %v, ret %v", got, tt.Receiver, tt.ret)
			}
		})
	}
}

// 有序列表去重
func Test_Uniquify(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Recv     *List[T]
		wantRecv *List[T]
		wantRet  int
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 空链表去重", NewFromSlice([]int{}), NewFromSlice([]int{}), 0},
		{"int 平凡链表去重", NewFromSlice([]int{1}), NewFromSlice([]int{1}), 0},
		{"int 非平凡链表", NewFromSlice([]int{3, 3, 3, 3, 5, 5, 5, 5, 5, 8, 8, 8, 13, 13, 13, 13}), NewFromSlice([]int{3, 5, 8, 13}), 12},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Recv.Uniquify(); !(tt.Recv.DeepEqual(tt.wantRecv) && got == tt.wantRet) {
				t.Errorf("Recv.Uniquify() = %v, Recv => %v, want %v %v", got, tt.Recv, tt.wantRet, tt.wantRecv)
			}
		})
	}
}
