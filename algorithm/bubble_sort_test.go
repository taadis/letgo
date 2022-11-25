package algorithm

import (
	"testing"
)

// 测试冒泡排序
func TestBubbleSort(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		var values []int
		result := BubbleSort(values)
		t.Log(result)
	})

	t.Run("BubbleSort", func(t *testing.T) {
		values := []int{100, 1, 9, 8, 23, 5, 8, 7, 6}
		result := BubbleSort(values)
		t.Log(result)
	})

	t.Run("2135", func(t *testing.T) {
		values := []int{2, 1, 3, 5}
		sortedValues := BubbleSort(values)
		t.Log(sortedValues)
	})

	t.Run("42713", func(t *testing.T) {
		values := []int{4, 2, 7, 1, 3}
		sortedValues := BubbleSort(values)
		t.Log(sortedValues)
	})
}
