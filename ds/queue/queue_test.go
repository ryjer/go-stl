package queue

import (
	"testing"

	num "gitee.com/ryjer/go-generic/number"
)

// 内容视图判等
func Test_DeepEqual(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string
		Receiver *Queue[T]
		arg      *Queue[T]
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
				t.Errorf("Receiver = %v .DeepEqual(%v) = %v, want %v", tt.Receiver, tt.arg, got, tt.want)
			}
		})
	}
}

// 空队列构造
func Test_New(t *testing.T) {
	t.Run("空列表初始化", func(t *testing.T) {
		if got := New[int](); !((got.size == 0) && (got.header.pre == nil && got.header.next == got.trailer) && (got.trailer.pre == got.header && got.trailer.next == nil)) {
			t.Errorf("New() = %v, want %v", got, got.String())
		}
	})
}

// 作为末节点插入
func Test_Enqueue(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string    // 测试用例名
		Receiver *Queue[T] // 接收对象
		arg      T         // 多参数
	}
	// int 类型测试，只插入一个节点
	// 构造队列 {4 [1 2 3 4]}
	intList := New[int]()
	intList.trailer.InsertAsPre(1)
	intList.trailer.InsertAsPre(2)
	intList.trailer.InsertAsPre(3)
	intList.trailer.InsertAsPre(4)
	intList.size = 4
	intTests := []testCase[int]{
		{"int 插入一个数", intList, 10},
		{"int 插入另一个数", intList, 11},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			oldLast := tt.Receiver.trailer.pre // 记录旧首节点
			// 插入节点，进行测试
			if got := tt.Receiver.Enqueue(tt.arg); !((got.data == tt.arg) && got.IsBetween(oldLast, tt.Receiver.trailer) && (intList.size == 5)) {
				t.Errorf("this.Enqueue(%v).data = %v, Receiver => %v, want %v", tt.arg, got.data, intList.String(), got.data)
			}
			// 复原链表，保持每个测试用例环境一致性
			oldLast.next = tt.Receiver.trailer
			tt.Receiver.trailer.pre = oldLast
			intList.size--
		})
	}
}

// 弹出首节点
func Test_Dequeue(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string    // 测试用例名
		Receiver *Queue[T] // 接收对象
		want     T
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 下界弹出", NewFromSlice([]int{4}), 4},
		{"int 随机长度弹出", NewFromSlice([]int{1, 2, 3, 4}), 1},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.Dequeue(); !(got == tt.want) {
				t.Errorf("this.Dequeue() = %v, Receiver => %v, want %v", got, tt.Receiver, tt.want)
			}
		})
	}
}

// 获取首节点
func Test_FrontNode(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string    // 测试用例名
		Receiver *Queue[T] // 接收对象
		want     T
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 下界弹出", NewFromSlice([]int{4}), 4},
		{"int 随机长度弹出", NewFromSlice([]int{1, 2, 3, 4}), 1},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.FrontNode(); !(got.data == tt.want) {
				t.Errorf("this.FrontNode().data = %v, Receiver => %v, want %v", got.data, tt.Receiver, tt.want)
			}
		})
	}
}

// 获取首元素
func Test_Front(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string    // 测试用例名
		Receiver *Queue[T] // 接收对象
		want     T
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 下界弹出", NewFromSlice([]int{4}), 4},
		{"int 随机长度弹出", NewFromSlice([]int{1, 2, 3, 4}), 1},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.Front(); !(got == tt.want) {
				t.Errorf("this.Front() = %v, Receiver => %v, want %v", got, tt.Receiver, tt.want)
			}
		})
	}
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
