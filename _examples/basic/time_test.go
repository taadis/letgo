package basic

import (
	"testing"
	"time"
)

// 获取时间戳
func TestTime1(t *testing.T) {
	t.Log("时间戳(单位:秒,位数:10):\t", time.Now().Unix())
	t.Log("时间戳(单位:毫秒,位数:13):\t", time.Now().UnixNano()/1e6)
	t.Log("时间戳(单位:纳秒,位数:19):\t", time.Now().UnixNano())
}
