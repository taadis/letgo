package concurrent

import (
	"sync"
	"testing"
)

// TestMutex1:
func TestMutex1(t *testing.T) {
	mu := sync.Mutex{}
	mu.Lock()
	mu.Unlock()
}

// TestMutex2:
// Unlock 未加锁的 Mutex 会 panic
// fatal error: sync: unlock of unlocked mutex
func TestMutex2(t *testing.T) {
	mu := sync.Mutex{}
	//mu.Lock()
	mu.Unlock()
}


