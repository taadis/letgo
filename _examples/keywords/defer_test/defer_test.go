package defer_test

import (
	"sync"
	"testing"
)

var mutex sync.Mutex

func foo() {
	mutex.Lock()
	mutex.Unlock()
}

func deferFoo() {
	mutex.Lock()
	defer mutex.Unlock()
}

func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo()
	}
}

func BenchmarkDeferCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deferFoo()
	}
}

/*
go test -bench=.
goos: windows
goarch: amd64
pkg: github.com/taadis/letgo/_examples/keywords/defer_test
BenchmarkCall-4         100000000               11.9 ns/op
BenchmarkDeferCall-4    34349997                34.2 ns/op
PASS
ok      github.com/taadis/letgo/_examples/keywords/defer_test    4.821s

go 1.13.8 中直接调用比延迟调用快3倍左右
go 1.14.x 中说是有延迟有性能提升, 近乎无格外损耗, 到时候看看.
*/
