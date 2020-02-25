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

// 测试数值转字符串
func TestItoa(t *testing.T) {
	i := 10
	s := strconv.Itoa(i)
	t.Log(s)
}

// 接口对象转字符串
func TestStringToInterface(t *testing.T) {
	var a interface{}
	a = "str"
	s := a.(string)
	t.Log(s)
}
