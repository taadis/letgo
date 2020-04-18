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
