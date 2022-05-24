package basic

import (
	"testing"
	"time"
)

// 获取时间戳
func TestTime1(t *testing.T) {
	t.Log("时间戳(单位:秒,位数:10):\t", time.Now().Unix())          // 1653360862
	t.Log("时间戳(单位:毫秒,位数:13):\t", time.Now().UnixNano()/1e6) // 1653360862152
	t.Log("时间戳(单位:毫秒,位数:13):\t", time.Now().UnixMilli())    // 1653360862152
	t.Log("时间戳(单位:微秒,位数:16):\t", time.Now().UnixMicro())    // 1653360862152801
	t.Log("时间戳(单位:纳秒,位数:19):\t", time.Now().UnixNano())     // 1653360862152803000
}

// 时间戳转Time类型
func TestTimestampToTime(t *testing.T) {
	ts := time.Now().UnixNano() / 1e6
	s := time.UnixMilli(ts).Format("2006-01-02 15:04-05")
	t.Log(s) // 2022-05-24 10:55-02
}
