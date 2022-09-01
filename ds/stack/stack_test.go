package stack

import (
	"testing"

	num "gitee.com/ryjer/go-generic/number"
)

// 无参构造
func Test_New(t *testing.T) {
	type testCase[T num.Q] struct {
		name         string
		wantNewStack *Stack[T]
	}
	intTests := []testCase[int]{
		{"int 测试", &Stack[int]{0, defaultCapacity, []int{}}},
		{"int 测试", &Stack[int]{0, defaultCapacity, []int{0, 0, 0, 0}}},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewStack := New[int](); !gotNewStack.DeepEqual(tt.wantNewStack) {
				t.Errorf("New() = %v, want %v", gotNewStack, tt.wantNewStack)
			}
		})
	}
}

// 从切片构造
func Test_NewFromSlice(t *testing.T) {
	type testCase[T num.Q] struct {
		name         string
		arg          []T
		wantNewStack *Stack[T]
	}
	intTests := []testCase[int]{
		{"int 空构造", []int{}, &Stack[int]{0, defaultCapacity, []int{}}},
		{"int 非空构造", []int{1, 2, 3, 4, 5, 6, 7, 8}, &Stack[int]{8, defaultCapacity, []int{1, 2, 3, 4, 5, 6, 7, 8}}},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewStack := NewFromSlice(tt.arg); !gotNewStack.DeepEqual(tt.wantNewStack) {
				t.Errorf("NewFromSlice(%v) = %v, want %v", tt.arg, gotNewStack, tt.wantNewStack)
			}
		})
	}
}

// 已用容量测试
func Test_Size(t *testing.T) {
	type testCases[T num.Q] struct {
		name string    // 测试用例名
		Recv *Stack[T] // 接收对象
		want int       // 期望结果
	}
	// int 测试
	intTests := []testCases[int]{
		{"int 空栈", New[int](), 0},
		{"int 非空栈", &Stack[int]{4, defaultCapacity, []int{1, 2, 3, 4}}, 4},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Recv.Size(); got != tt.want {
				t.Errorf("this.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 空向量判断
func Test_IsEmpty(t *testing.T) {
	type testCase[T num.Q] struct {
		name string    // 测试用例名
		Recv *Stack[T] // 接收对象
		want bool      // 预期结果
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 空栈", New[int](), true},
		{"int 非空栈", &Stack[int]{4, defaultCapacity, []int{1, 2, 3, 4}}, false},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := tt.Recv.IsEmpty(); gotRet != tt.want {
				t.Errorf("this.IsEmpty() = %v, Recv = %v, want %v", gotRet, tt.Recv, tt.want)
			}
		})
	}
}

// 清空
func Test_Clear(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string    // 测试用例名
		Recv     *Stack[T] // 接收对象
		wantRecv *Stack[T] // 预期结果
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 空向量", New[int](), &Stack[int]{0, defaultCapacity, []int{}}},
		{"int 非空向量", &Stack[int]{4, defaultCapacity, []int{1, 2, 3, 4}}, &Stack[int]{0, defaultCapacity, []int{}}},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.Recv.Clear(); !tt.Recv.DeepEqual(tt.wantRecv) {
				t.Errorf("this.Clear(), Recv = %v, want %v", tt.Recv, tt.wantRecv)
			}
		})
	}
}

// 翻倍扩容测试
func Test_expand(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string    // 测试用例名
		Recv     *Stack[T] // 接收对象
		wantRecv *Stack[T] // 预期结果
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 空向量", New[int](), &Stack[int]{0, defaultCapacity, []int{}}},
		{"int 非空不扩容", &Stack[int]{defaultCapacity - 1, defaultCapacity, make([]int, defaultCapacity*2)}, &Stack[int]{defaultCapacity - 1, defaultCapacity, make([]int, defaultCapacity*2)}},
		{"int 非空扩容", &Stack[int]{defaultCapacity, defaultCapacity, make([]int, defaultCapacity)}, &Stack[int]{defaultCapacity, defaultCapacity * 2, make([]int, defaultCapacity)}},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.Recv.expand(); !tt.Recv.DeepEqual(tt.wantRecv) {
				t.Errorf("this.expand(), Recv = %v, want %v", tt.Recv, tt.wantRecv)
			}
		})
	}
}

// 压栈
func Test_Push(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string    // 测试用例名
		Recv     *Stack[T] // 接收对象
		arg      T         // 单参数，秩
		wantRecv *Stack[T] // 预期结果
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 空栈压栈1", New[int](), 1, &Stack[int]{1, defaultCapacity, []int{1}}},
		{"int 非空压栈3", NewFromSlice([]int{1, 2, 3}), 4, &Stack[int]{4, 8, []int{1, 2, 3, 4}}},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.Recv.Push(tt.arg); !tt.Recv.DeepEqual(tt.wantRecv) {
				t.Errorf("this.Push(%v), Recv = %v, wantRecv %v", tt.arg, tt.Recv, tt.wantRecv)
			}
		})
	}
}

// 减半缩容测试
func Test_shrink(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string    // 测试用例名
		Recv     *Stack[T] // 接收对象
		wantRecv *Stack[T] // 预期结果
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 空栈不缩容", New[int](), &Stack[int]{0, defaultCapacity, []int{}}},
		{"int 空栈缩容", &Stack[int]{0, defaultCapacity * 4, make([]int, defaultCapacity*4)}, &Stack[int]{0, defaultCapacity * 2, []int{}}},
		{"int 满栈不缩容", &Stack[int]{defaultCapacity, defaultCapacity, make([]int, defaultCapacity)}, &Stack[int]{defaultCapacity, defaultCapacity, make([]int, defaultCapacity)}},
		{"int 少于一半缩容", &Stack[int]{defaultCapacity, defaultCapacity * 2, make([]int, defaultCapacity*2)}, &Stack[int]{defaultCapacity, defaultCapacity, make([]int, defaultCapacity*2)}},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			oldCap := tt.Recv.capacity
			if tt.Recv.shrink(); !tt.Recv.DeepEqual(tt.wantRecv) {
				t.Errorf("this.shrink(), oldCap=%v, newRecv=%v,  wantRecv %v", oldCap, tt.Recv, tt.wantRecv)
			}
		})
	}
}

// 出栈
func Test_Pop(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string    // 测试用例名
		Recv     *Stack[T] // 接收对象
		wantRecv *Stack[T] // 预期对象变化
		wantRet  T         // 预期结果
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 上界弹出", NewFromSlice([]int{1}), NewFromSlice([]int{}), 1},
		{"int 非空向量", NewFromSlice([]int{1, 2, 3, 4}), NewFromSlice([]int{1, 2, 3}), 4},
		{"int 非空向量", NewFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}), NewFromSlice([]int{1, 2, 3, 4, 5, 6, 7}), 8},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := tt.Recv.Pop(); !(tt.Recv.DeepEqual(tt.wantRecv) && (gotRet == tt.wantRet)) {
				t.Errorf("this.Pop() = %v, Recv = %v, want %v %v", gotRet, tt.Recv, tt.wantRet, tt.wantRecv)
			}
		})
	}
}

// 取顶
func Test_Top(t *testing.T) {
	type testCase[T num.Q] struct {
		name    string    // 测试用例名
		Recv    *Stack[T] // 接收对象
		wantRet T         // 预期结果
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 上界", NewFromSlice([]int{1}), 1},
		{"int 非空", NewFromSlice([]int{1, 2, 3, 4}), 4},
		{"int 非空", NewFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}), 8},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := tt.Recv.Pop(); !(gotRet == tt.wantRet) {
				t.Errorf("this.Top() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
