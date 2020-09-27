package faq99

import "testing"

func TestFaq99(t *testing.T)  {
	var a uint32 = 10
	var b uint32 = 100
	v := a - b
	t.Log(v)
	// output: 4294967206
}

// a:
// uint32 相减为负数时, 用 := 的话会自动把 v 当作 uint32 类型
// 系统会把高位的 1 当成最高进制来计算, 于是就废了.
// 可改为:
// v := int(a) - int(b)
