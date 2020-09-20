package algorithm

import (
	"math/rand"
	"testing"
	"time"
)

func TestSelectionSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var slice []int
	for i := 0; i < 9; i++{
		slice = append(slice, rand.Intn(100))
	}
	t.Log("Befor Selection Sort:", slice)
	slice = SelectionSort(slice)
	t.Log("After Selection Sort:", slice)
}
