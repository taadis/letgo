package algorithm

import (
	"testing"
)

// 测试冒泡排序
func TestBubbleSort(t *testing.T) {
	values := []int{100, 1, 9, 8, 23, 5, 8, 7, 6}
	result := BubbleSort(values)
	t.Log(result)
}
