package faq

import (
	"sync"
	"testing"
)

// 如果wg.Add调用维护的计数更改成了负数,将产生panic
func TestSyncWaitGroupPanic1(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(-1)
}

// 如果wg.Done()的调用维护的计数变成了负数,将产生panic
func TestSyncWaitGroupPanic2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 11; i++ {
		wg.Done()
	}
}
