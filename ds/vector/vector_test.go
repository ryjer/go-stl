package vector

import (
	"strconv"
	"testing"

	num "gitee.com/ryjer/go-generic/number"
)

// 无参构造
func TestNew(t *testing.T) {
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
func TestNewFromSlice(t *testing.T) {
	type testCase[T num.Q] struct {
		name string
		arg  []T
		want string
	}
	// int 类型测试
	intTests := []testCase[int]{
		{"int 0测试", []int{}, "{0 " + strconv.Itoa(defaultCapacity) + " []}"},
		{"int 最小容量", []int{1, 2, 3, 4, 5, 6, 7, 8}, "{8 " + strconv.Itoa(defaultCapacity) + " [1 2 3 4 5 6 7 8]}"},
		{"int 扩容1次", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, "{9 " + strconv.Itoa(defaultCapacity*2) + " [1 2 3 4 5 6 7 8 9]}"},
		{"int 扩容2次", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, "{17 " + strconv.Itoa(defaultCapacity*4) + " [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17]}"},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromSlice(tt.arg); got.String() != tt.want {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
	// uint 类型测试
	uintTests := []testCase[uint]{
		{"uint 0测试", []uint{}, "{0 8 []}"},
		{"uint 最小容量", []uint{1, 2, 3, 4, 5, 6, 7, 8}, "{8 " + strconv.Itoa(defaultCapacity) + " [1 2 3 4 5 6 7 8]}"},
		{"uint 扩容1次", []uint{1, 2, 3, 4, 5, 6, 7, 8, 9}, "{9 " + strconv.Itoa(defaultCapacity*2) + " [1 2 3 4 5 6 7 8 9]}"},
		{"uint 扩容2次", []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, "{17 " + strconv.Itoa(defaultCapacity*4) + " [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17]}"},
	}
	for _, tt := range uintTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromSlice(tt.arg); got.String() != tt.want {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
	// float32 类型测试
	float32Tests := []testCase[float32]{
		{"float32 0测试", []float32{}, "{0 8 []}"},
		{"float32 最小容量", []float32{1, 2, 3, 4, 5, 6, 7, 8}, "{8 " + strconv.Itoa(defaultCapacity) + " [1 2 3 4 5 6 7 8]}"},
		{"float32 扩容1次", []float32{1, 2, 3, 4, 5, 6, 7, 8, 9}, "{9 " + strconv.Itoa(defaultCapacity*2) + " [1 2 3 4 5 6 7 8 9]}"},
		{"float32 扩容2次", []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, "{17 " + strconv.Itoa(defaultCapacity*4) + " [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17]}"},
	}
	for _, tt := range float32Tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromSlice(tt.arg); got.String() != tt.want {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 方法测试，有返回值，无参数
func TestSize(t *testing.T) {
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

func Test_enoughSize(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{"0 测试", 0, defaultCapacity},
		{"最小容量", defaultCapacity, defaultCapacity},
		{"拓展1次", defaultCapacity + 1, defaultCapacity * 2},
		{"拓展2次", defaultCapacity*2 + 1, defaultCapacity * 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := enoughSize(tt.arg); got != tt.want {
				t.Errorf("enoughSize(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}
