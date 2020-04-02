package basic

import (
	"testing"
)

func TestSlice1(t *testing.T) {
	s1 := []int{1, 2, 3}
	t.Log(s1)

	s2 := ([3]int{1, 2, 3})[:]
	// invalid operation [3]int literal[:] (slice of unaddressable value)
	t.Log(s2)

	s3 := s1[:]
	t.Log(s3)
}
