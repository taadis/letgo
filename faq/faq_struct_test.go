package faq

import (
	"testing"
	"unsafe"
)

// 空结构体的内存占用
func TestStruct0_EmptyStruct(t *testing.T) {
	s := struct{}{}
	t.Logf("empty struct sizeof:%d", unsafe.Sizeof(s))
}
