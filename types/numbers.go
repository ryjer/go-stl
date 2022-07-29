package types

type Num interface {
	// 测试用临时泛型，一维可比较数字类型
	RationalNumber
}

// 有理数泛型_全程
type RationalNumber interface {
	Integer | Float
}

// 有理数泛型 Q
type Qnum interface {
	Integer | Float
}

// 整数泛型 Z
type Znum interface {
	Integer
}

// 自然数泛型 N
type Nnum interface {
	Uint
}

// 机器复数类型
type Complex interface {
	complex64 | complex128
}

// 机器浮点数泛型
type Float interface {
	~float32 | ~float64
}

// 机器整数泛型
type Integer interface {
	Int | Uint
}

// 机器无符号数泛型
type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// 机器有符号数泛型
type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
