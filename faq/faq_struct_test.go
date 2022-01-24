package faq

import (
	"fmt"
	"testing"
	"unsafe"
)

// 空结构的使用1-仅含方法的空结构体
func TestStruct2_EmptyStructOnlyMethods(t *testing.T) {
	foo := new(Foo)
	foo.F1()
	foo.F2()
}

type Foo struct{}

func (f *Foo) F1() {
	fmt.Printf("empty struct f1\n")
}
func (f *Foo) F2() {
	fmt.Printf("empty struct f2\n")
}

// 空结构体的内存占用
func TestStruct0_EmptyStruct(t *testing.T) {
	s := struct{}{}
	t.Logf("empty struct sizeof:%d", unsafe.Sizeof(s))
}
