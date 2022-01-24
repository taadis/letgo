package faq

import (
	"fmt"
	"testing"
	"unsafe"
)

// 空结构体的使用-实现集合Set
func TestEmptyStruct_Set(t *testing.T) {
	set := make(Set)
	set.Add("a")
	set.Add("b")
	set.Add("c")
	t.Logf("set has a:%v", set.Has("a"))
	t.Logf("set has b:%v", set.Has("b"))
	t.Logf("set has d:%v", set.Has("d"))
}

type Set map[string]struct{}

func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Add(key string) {
	s[key] = struct{}{}
}

func (s Set) Remove(key string) {
	delete(s, key)
}

// 空结构体的使用1-仅含方法的空结构体
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
