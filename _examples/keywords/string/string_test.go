package string

import (
	"strings"
	"testing"
)

// string's zero value
func TestFoo(t *testing.T) {
	var s string
	ss := ""
	t.Logf("string's zero value is %s %s", s, ss)
}

// 测试字符串分隔
func TestStringSplit(t *testing.T) {
	s := "a-b-c"
	sSlice := strings.Split(s, "-") // 返回的是字符串切片
	t.Log(sSlice)
}

// go test -v string_test.go
