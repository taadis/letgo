package algorithm

import (
	"testing"
)

// TestFibonacci
func TestFibonacci(t *testing.T) {
	for i := 0; i < 10; i++ {
		ret := Fibonacci(i)
		t.Log(ret)
	}
}
