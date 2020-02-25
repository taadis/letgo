// 字符串转换测试
package string

import (
	"strconv"
	"testing"
)

// 测试字符串转数值
func TestStringAtoi(t *testing.T) {
	s := "10"
	i, err := strconv.Atoi(s)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(i)
}
