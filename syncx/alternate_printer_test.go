package syncx

import "testing"

func TestAlternatePrinter(t *testing.T) {
	p := NewAlternatePrinter(100)
	p.Run()
}
