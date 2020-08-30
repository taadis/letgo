package basic

import (
	"testing"
	"unsafe"
)

// 测试数据类型存储大小
func TestSizeof(t *testing.T) {
	var b bool
	var a int
	var f32 float32
	var f64 float64
	t.Log("数据类型 存储大小(单位:字节 byte)")
	t.Log("bool", unsafe.Sizeof(b))
	t.Log("int", unsafe.Sizeof(a))
	t.Log("float32", unsafe.Sizeof(f32))
	t.Log("float64", unsafe.Sizeof(f64))
	// ...
}

// 当前包目录下执行以下命令测试
// go test -v datatype_test.go
