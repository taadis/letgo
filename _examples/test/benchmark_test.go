// benchmark_test.go
package main

import (
	"testing"
)

func Benchmark_For_A(b *testing.B) {
	l := 0
	for i, n := 0, b.N; i < n; i++ {
		l++
	}
}

func Benchmark_For_B(b *testing.B) {
	l := 0
	for i := 0; i < b.N; i++ {
		l++
	}
}

/*
go test -v -bench=.
goos: windows
goarch: amd64
pkg: go-example/test
Benchmark_For_A-4       2000000000               0.29 ns/op
Benchmark_For_B-4       2000000000               0.58 ns/op
PASS
ok      go-example/test 3.157s
*/
