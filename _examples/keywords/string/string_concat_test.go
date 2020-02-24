package string

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// using +=
func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	s := ""
	for i := 0; i < b.N; i++ {
		s += "s"
	}
	b.StopTimer()
}

// using Sprintf
func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	s := "s"
	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("%v", s)
	}
	b.StopTimer()
}

// using strings.Builder
func BenchmarkStringsBuilder(b *testing.B) {
	b.ResetTimer()
	var stringsBuilder strings.Builder
	s := "s"
	for i := 0; i < b.N; i++ {
		stringsBuilder.WriteString(s)
	}
	_ = stringsBuilder.String()
	b.StopTimer()
}

// using bytes.Buffer
func BenchmarkBytesBuffer(b *testing.B) {
	b.ResetTimer()
	s := "s"
	var bytesBuffer bytes.Buffer
	for i := 0; i < b.N; i++ {
		bytesBuffer.WriteString(s)
	}
	_ = bytesBuffer.String()
	b.StopTimer()
}

// > go test -v -bench=.
// output:
/*
goos: windows
goarch: amd64
pkg: github.com/taadis/letgo/_examples/keywords/string
BenchmarkStringAdd-4              600320             48977 ns/op
BenchmarkSprintf-4               9096511               129 ns/op
BenchmarkStringsBuilder-4       211771195                5.40 ns/op
BenchmarkBytesBuffer-4          100000000               11.8 ns/op
PASS
ok      github.com/taadis/letgo/_examples/keywords/string       36.108s

*/
