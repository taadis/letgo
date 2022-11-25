package intx

import "testing"

func TestHasDuplicateValue0(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var values []int
		got := HasDuplicateValue0(values)
		if got != false {
			t.Fatalf("want false but got %v", got)
		}
	})

	t.Run("123", func(t *testing.T) {
		values := []int{1, 2, 3}
		got := HasDuplicateValue0(values)
		if got != false {
			t.Fatalf("want false but got %v", got)
		}
	})

	t.Run("010", func(t *testing.T) {
		values := []int{0, 1, 0}
		got := HasDuplicateValue0(values)
		if got != true {
			t.Fatalf("want true but got %v", got)
		}
	})
}

func TestHasDuplicateValue1(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var values []int
		got := HasDuplicateValue1(values)
		if got != false {
			t.Fatalf("want false but got %v", got)
		}
	})

	t.Run("123", func(t *testing.T) {
		values := []int{1, 2, 3}
		got := HasDuplicateValue1(values)
		if got != false {
			t.Fatalf("want false but got %v", got)
		}
	})

	t.Run("010", func(t *testing.T) {
		values := []int{0, 1, 0}
		got := HasDuplicateValue1(values)
		if got != true {
			t.Fatalf("want true but got %v", got)
		}
	})
}
