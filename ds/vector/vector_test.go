package vector

import (
	"strconv"
	"testing"

	num "gitee.com/ryjer/go-generic/number"
)

// 无参构造
func Test_New(t *testing.T) {
	type testCase struct {
		name string
		want string
	}
	// int 类型测试
	intTest := testCase{"int 测试", "{0 " + strconv.Itoa(defaultCapacity) + " []}"}
	if got := New[int](); got.String() != intTest.want {
		t.Errorf("New() = %v, want %v", got, intTest.want)
	}
	// uint 类型测试
	uintTest := testCase{"uint 测试", "{0 " + strconv.Itoa(defaultCapacity) + " []}"}
	if got := New[uint](); got.String() != uintTest.want {
		t.Errorf("New() = %v, want %v", got, uintTest.want)
	}
	// float32 类型测试
	float32Test := testCase{"int 测试", "{0 " + strconv.Itoa(defaultCapacity) + " []}"}
	if got := New[float32](); got.String() != float32Test.want {
		t.Errorf("New() = %v, want %v", got, float32Test.want)
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
		{"int 0测试", []int{}, "{0 " + strconv.Itoa(defaultCapacity) + " []}"},
		{"int 最小容量", []int{1, 2, 3, 4, 5, 6, 7, 8}, "{8 " + strconv.Itoa(defaultCapacity*2) + " [1 2 3 4 5 6 7 8]}"},
		{"int 扩容1次", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, "{9 18 [1 2 3 4 5 6 7 8 9]}"},
		{"int 扩容2次", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, "{17 34 [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17]}"},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromSlice(tt.arg); got.String() != tt.want {
				t.Errorf("NewFromSlicew() = %v, want %v", got, tt.want)
			}
		})
	}
	// uint 类型测试
	uintTests := []testCase[uint]{
		{"uint 0测试", []uint{}, "{0 8 []}"},
		{"uint 最小容量", []uint{1, 2, 3, 4, 5, 6, 7, 8}, "{8 " + strconv.Itoa(defaultCapacity*2) + " [1 2 3 4 5 6 7 8]}"},
		{"uint 扩容1次", []uint{1, 2, 3, 4, 5, 6, 7, 8, 9}, "{9 18 [1 2 3 4 5 6 7 8 9]}"},
		{"uint 扩容2次", []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, "{17 34 [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17]}"},
	}
	for _, tt := range uintTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromSlice(tt.arg); got.String() != tt.want {
				t.Errorf("NewFromSlice() = %v, want %v", got, tt.want)
			}
		})
	}
	// float32 类型测试
	float32Tests := []testCase[float32]{
		{"float32 0测试", []float32{}, "{0 8 []}"},
		{"float32 最小容量", []float32{1, 2, 3, 4, 5, 6, 7, 8}, "{8 " + strconv.Itoa(defaultCapacity*2) + " [1 2 3 4 5 6 7 8]}"},
		{"float32 扩容1次", []float32{1, 2, 3, 4, 5, 6, 7, 8, 9}, "{9 18 [1 2 3 4 5 6 7 8 9]}"},
		{"float32 扩容2次", []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, "{17 34 [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17]}"},
	}
	for _, tt := range float32Tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromSlice(tt.arg); got.String() != tt.want {
				t.Errorf("NewFromSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

// NewFromVector 从已有列表/向量子区间构造
func Test_NewFromVector(t *testing.T) {
	type args[T num.Q] struct {
		Vector *Vector[T]
		lo, hi int
	}
	type testCase[T num.Q] struct {
		name string
		args args[T]
		want string
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 对空向量", args[int]{NewFromSlice[int]([]int{}), 0, 0}, "{0 " + strconv.Itoa(defaultCapacity) + " []}"},
		{"int 0区间", args[int]{NewFromSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8}), 1, 1}, "{0 " + strconv.Itoa(defaultCapacity) + " []}"},
		{"int 负区间", args[int]{NewFromSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8}), 0, 0}, "{0 " + strconv.Itoa(defaultCapacity) + " []}"},
		{"int 完全区间", args[int]{NewFromSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8}), 0, 8}, "{8 " + strconv.Itoa(defaultCapacity*2) + " [1 2 3 4 5 6 7 8]}"},
		{"int 下边界", args[int]{NewFromSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8}), 0, 4}, "{4 " + strconv.Itoa(defaultCapacity) + " [1 2 3 4]}"},
		{"int 上边界", args[int]{NewFromSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8}), 4, 8}, "{4 " + strconv.Itoa(defaultCapacity) + " [5 6 7 8]}"},
		{"int 中间随机", args[int]{NewFromSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8}), 2, 6}, "{4 " + strconv.Itoa(defaultCapacity) + " [3 4 5 6]}"},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromVector(tt.args.Vector, tt.args.lo, tt.args.hi); got.String() != tt.want {
				t.Errorf("NewFromVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 已用容量测试
func Test_Size(t *testing.T) {
	type testCases[T num.Q] struct {
		name     string     // 测试用例名
		Receiver *Vector[T] // 接收对象
		want     int        // 期望结果
	}
	// int 测试
	intTests := []testCases[int]{
		{"int 0", NewFromSlice([]int{}), 0},
		{"int 最小容量", NewFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}), 8},
		{"int 扩展一次", NewFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), 9},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.Size(); got != tt.want {
				t.Errorf("this.Size() = %v, want %v", got, tt.want)
			}
		})
	}
	// uint 测试
	uintTests := []testCases[uint]{
		{"uint 0", NewFromSlice([]uint{}), 0},
		{"uint 最小容量", NewFromSlice([]uint{1, 2, 3, 4, 5, 6, 7, 8}), 8},
		{"uint 拓展一次", NewFromSlice([]uint{1, 2, 3, 4, 5, 6, 7, 8, 9}), 9},
	}
	for _, tt := range uintTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.Size(); got != tt.want {
				t.Errorf("this.Size() = %v, want %v", got, tt.want)
			}
		})
	}
	// float 测试
	floatTests := []testCases[float32]{
		{"float32 0", NewFromSlice([]float32{}), 0},
		{"float32 最小容量", NewFromSlice([]float32{1, 2, 3, 4, 5, 6, 7, 8}), 8},
		{"float32 拓展一次", NewFromSlice([]float32{1, 2, 3, 4, 5, 6, 7, 8, 9}), 9},
	}
	for _, tt := range floatTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.Size(); got != tt.want {
				t.Errorf("this.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 空向量判断
func Test_IsEmpty(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string     // 测试用例名
		Receiver *Vector[T] // 接收对象
		want     bool       // 预期结果
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 空向量测试", NewFromSlice([]int{}), true},
		{"int 短非空", NewFromSlice([]int{1, 2, 3}), false},
		{"int 长非空", NewFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), false},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Receiver.IsEmpty(); got != tt.want {
				t.Errorf("this.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 寻秩访问
func Test_Get(t *testing.T) {
	type testCase[T num.Q] struct {
		name     string     // 测试用例名
		Receiver *Vector[T] // 接收对象
		arg      int        // 单参数，秩
		want     int        // 预期结果
	}
	// int 类型测试
	intTests := []testCase[int]{
		// {"int 空向量测试", NewFromSlice([]int{}), -1, 0},
		{"int 下边界", NewFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}), 0, 1},
		{"int 中间随机", NewFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}), 3, 4},
		{"int 上边界", NewFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}), 7, 8},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.Receiver.Get(tt.arg); got != tt.want {
				t.Errorf("this.Get(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

// 寻址修改
func Test_Put(t *testing.T) {
	type args[T num.Q] struct {
		r int
		e T
	}
	type testCase[T num.Q] struct {
		name     string     // 测试用例名
		Receiver *Vector[T] // 接收对象
		args     args[T]    // 单参数，秩
		want     string     // 预期结果
	}
	// int 类型测试
	intTests := []testCase[int]{
		// {"int 空向量测试", NewFromSlice([]int{}), args[int]{0,1}, 0},
		{"int 下边界", NewFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}), args[int]{0, 0}, "{8 16 [0 2 3 4 5 6 7 8]}"},
		{"int 中间随机", NewFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}), args[int]{5, 0}, "{8 16 [1 2 3 4 5 0 7 8]}"},
		{"int 上边界", NewFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}), args[int]{7, 0}, "{8 16 [1 2 3 4 5 6 7 0]}"},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.Receiver.Put(tt.args.r, tt.args.e); tt.Receiver.String() != tt.want {
				t.Errorf("this.Put(%v, %v) => %v, want %v", tt.args.r, tt.args.e, tt.Receiver.String(), tt.want)
			}
		})
	}
}

// 插入单个元素
func Test_Insert(t *testing.T) {
	type args[T num.Q] struct {
		r int
		e T
	}
	type testCase[T num.Q] struct {
		name     string     // 测试用例名
		Receiver *Vector[T] // 接收对象
		args     args[T]    // 单参数，秩
		want     string     // 预期结果
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 空向量测试", NewFromSlice([]int{}), args[int]{0, 0}, "{1 8 [0]}"},
		{"int 下边界", NewFromSlice([]int{1, 2, 3, 4, 5, 6, 7}), args[int]{0, 0}, "{8 8 [0 1 2 3 4 5 6 7]}"},
		{"int 中间随机", NewFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}), args[int]{5, 0}, "{9 16 [1 2 3 4 5 0 6 7 8]}"},
		{"int 上边界", NewFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}), args[int]{7, 0}, "{9 16 [1 2 3 4 5 6 7 0 8]}"},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.Receiver.Insert(tt.args.r, tt.args.e); tt.Receiver.String() != tt.want {
				t.Errorf("this.Insert(%v, %v) => %v, want %v", tt.args.r, tt.args.e, tt.Receiver.String(), tt.want)
			}
		})
	}
}

// 区间移除
func Test_Remove(t *testing.T) {
	type args struct {
		lo, hi int
	}
	type wants[T num.Q] struct {
		receiver string
		ret      int
	}
	type testCase[T num.Q] struct {
		name     string     // 测试用例名
		receiver *Vector[T] // 接收对象
		args     args       // 区间范围参数
		wants    wants[T]   // 预期结果
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 空向量测试", NewFromSlice([]int{}), args{0, 0}, wants[int]{"{0 8 []}", 0}},
		{"int 下边界", NewFromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}), args{0, 2}, wants[int]{"{7 18 [2 3 4 5 6 7 8]}", 2}},
		{"int 中间随机", NewFromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}), args{3, 6}, wants[int]{"{6 18 [0 1 2 6 7 8]}", 3}},
		{"int 上边界", NewFromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}), args{7, 9}, wants[int]{"{7 18 [0 1 2 3 4 5 6]}", 2}},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.receiver.Remove(tt.args.lo, tt.args.hi); tt.receiver.String() != tt.wants.receiver || got != tt.wants.ret {
				t.Errorf("this.Remove(%v, %v) = %v, Receiver => %v,want %v %v", tt.args.lo, tt.args.hi, got, tt.receiver.String(), tt.wants.ret, tt.wants.receiver)
			}
		})
	}
}

// 单元素移除
func Test_Remove1(t *testing.T) {
	type wants[T num.Q] struct {
		receiver string
		ret      int
	}
	type testCase[T num.Q] struct {
		name     string     // 测试用例名
		receiver *Vector[T] // 接收对象
		arg      int        // 单参数，秩
		wants    wants[T]   // 预期结果
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 空向量测试", NewFromSlice([]int{}), 0, wants[int]{"{0 8 []}", 0}},
		{"int 下边界", NewFromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}), 0, wants[int]{"{8 18 [1 2 3 4 5 6 7 8]}", 0}},
		{"int 中间随机", NewFromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}), 4, wants[int]{"{8 18 [0 1 2 3 5 6 7 8]}", 4}},
		{"int 上边界", NewFromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}), 8, wants[int]{"{8 18 [0 1 2 3 4 5 6 7]}", 8}},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.receiver.Remove1(tt.arg); tt.receiver.String() != tt.wants.receiver || got != tt.wants.ret {
				t.Errorf("this.Remove1(%v) = %v, Receiver => %v,want %v %v", tt.arg, got, tt.receiver.String(), tt.wants.ret, tt.wants.receiver)
			}
		})
	}
}

// 用于计算最小目标容量的函数
func Test_enoughCapacity(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{"0 容量", 0, defaultCapacity},
		{"较小，但没有达到最小容量", defaultCapacity/2 - 1, defaultCapacity},
		{"最小容量", defaultCapacity, defaultCapacity * 2},
		{"大于最小容量", defaultCapacity + 1, (defaultCapacity + 1) * 2},
		{"较大", 800, 800 * 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := enoughCapacity(tt.arg); got != tt.want {
				t.Errorf("enoughCapacity(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}