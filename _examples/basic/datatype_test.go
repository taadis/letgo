package basic

import (
	"math"
	"testing"
	"unsafe"
)

// 测试数据类型存储大小
func TestSizeof(t *testing.T) {
	var b bool
	var by byte
	var ru rune
	var s string
	var a int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var f32 float32
	var f64 float64
	var ui uint
	var cx64 complex64
	var cx128 complex128

	t.Log("数据类型 存储大小(单位:字节 byte) | 值区间")
	t.Log("bool", unsafe.Sizeof(b))
	t.Log("byte", unsafe.Sizeof(by))
	t.Log("ru", unsafe.Sizeof(ru))
	t.Log("string", unsafe.Sizeof(s))
	t.Log("int", unsafe.Sizeof(a))
	t.Log("int8", unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)
	t.Log("int16", unsafe.Sizeof(i16), math.MinInt16, math.MaxInt16)
	t.Log("int32", unsafe.Sizeof(i32), math.MinInt32, math.MaxInt32)
	t.Log("int64", unsafe.Sizeof(i64), math.MinInt64, math.MaxInt64)
	t.Log("uint", unsafe.Sizeof(ui))
	t.Log("float32", unsafe.Sizeof(f32))
	t.Log("float64", unsafe.Sizeof(f64))
	t.Log("complex64", unsafe.Sizeof(cx64))
	t.Log("complex128", unsafe.Sizeof(cx128))
	// ...
}

// 当前包目录下执行以下命令测试
// go test -v datatype_test.go
