package tests

import (
	"testing"
)

// 验证new所说的指向零值的指针
// 相同类型的指针地址是否指向同一个内存地址?
// 结论：
// 不相同，零值不复用。
func TestNew(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		var ptr1 *int = new(int)
		var ptr2 *int = new(int)
		t.Logf("ptr1: %+v, ptr2: %+v", ptr1, ptr2)
		// 这里打印出来的地址不一样
		// output:
		// ptr1: 0xc0000123b0, ptr2: 0xc0000123b8
		// t.Logf("ptr1: %v, ptr2: %v", &ptr1, &ptr2)
	})

	t.Run("1", func(t *testing.T) {
		var ptr1, ptr2 *int
		t.Logf("ptr1: %+v, ptr2: %+v", ptr1, ptr2)
		// 这里打印出来的地址不一样
		// output:
		// ptr1: 0xc0000a82f0, ptr2: 0xc0000a82f8
	})
}
