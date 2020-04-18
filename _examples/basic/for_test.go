package basic

import (
	"testing"
)

// 测试无限循环
func TestFor1(t *testing.T) {
	for {
		t.Log("循环中...")
		break
	}
}

// 测试使用条件循环计算1-100的和
func TestFor2(t *testing.T) {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	t.Logf("sum = %d", sum)
}
