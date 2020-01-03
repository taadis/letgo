package oper

import (
	"testing"
)

// 测试 Access()
func Test_Access(t *testing.T) {
	err := Access("https://golang.google.cn")
	if err != nil {
		t.Error(err)
	}
}
