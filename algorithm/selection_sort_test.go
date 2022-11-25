package algorithm

import (
	"math/rand"
	"testing"
	"time"
)

func TestSelectionSort(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var values []int
		sortedValues := SelectionSort(values)
		t.Log(sortedValues)
	})

	t.Run("2613", func(t *testing.T) {
		values := []int{2, 6, 1, 3}
		sortedValues := SelectionSort(values)
		t.Log(sortedValues)
	})

	t.Run("42713", func(t *testing.T) {
		values := []int{4, 2, 7, 1, 3}
		sortedValues := SelectionSort(values)
		t.Log(sortedValues)
	})

	rand.Seed(time.Now().UnixNano())
	var slice []int
	for i := 0; i < 9; i++ {
		slice = append(slice, rand.Intn(100))
	}
	t.Log("before Selection Sort:", slice)
	slice = SelectionSort(slice)
	t.Log("after Selection Sort:", slice)
}
