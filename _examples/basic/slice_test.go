package basic

import (
	"testing"
)

func TestSlice1(t *testing.T) {
	s1 := []int{1, 2, 3}
	t.Log(s1)

	//s2 := ([3]int{1, 2, 3})[:]
	// invalid operation [3]int literal[:] (slice of unaddressable value)
	//t.Log(s2)

	s3 := s1[:]
	t.Log(s3)
}

func TestSlice(t *testing.T) {
	s := make([]int, 0, 10)
	s = append(s, 0)
	s = append(s, 1)
	s = append(s, 2)
	t.Logf("befor slice %+v %p", s, &s)
	AppendSlice(t, s)
	t.Logf("after slice %+v %p %d %d", s, &s, len(s), cap(s))
}

func AppendSlice(t *testing.T, s []int) {
	t.Logf("befor append slice %+v %p %d %d", s, &s, len(s), cap(s))
	s = append(s, 3)
	s = append(s, 4)
	s = append(s, 5)
	t.Logf("after append slice %+v %p %d %d", s, &s, len(s), cap(s))
}

func TestSlice2(t *testing.T) {
	s := make([]int, 0)
	s = append(s, 0)
	s = append(s, 1)
	s = append(s, 2)
	t.Logf("befor slice %+v", s)
	AppendSlice2(&s)
	t.Logf("after slice %+v", s)
}

func AppendSlice2(s *[]int) {
	*s = append(*s, 3)
}
