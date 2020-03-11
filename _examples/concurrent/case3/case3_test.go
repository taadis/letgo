package case3

import (
	"testing"
	"time"
)

// 先来看一个简单的数值累加运算,
// 用我们最常用的 for 直接循环累加.
func TestSum1(t *testing.T) {
	startTime := time.Now()
	sum := 0
	for i, n := 0, 1*10000*10000; i < n; i++ {
		sum += i
	}
	cost := time.Since(startTime)
	t.Logf("sum = %d", sum)
	t.Logf("cost = %d", cost.Microseconds())
}

// 如何用 go 来实现并行计算来提高性能?
func TestSum2(t *testing.T) {
	startTime := time.Now()
	sum := 0
	// ...
	cost := time.Since(startTime)
	t.Logf("sum = %d", sum)
	t.Logf("cost = %d", cost.Microseconds())
}
