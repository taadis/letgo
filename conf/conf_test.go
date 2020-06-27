package conf

import (
	"testing"
)

// TestSingleton
func TestSingleton(t *testing.T) {
	a := Database()
	b := Database()
	if a != b {
		t.Fatal("instance is not equal")
	}
}

// TestParallelSingleton
func TestParallelSingleton(t *testing.T) {

}
