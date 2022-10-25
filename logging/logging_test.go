package logging

import (
	"testing"
	"time"
)

func TestLogging(t *testing.T) {
	t.Logf("t-log")
	Infof("test-error log %s", time.Now().String())
	Errorf("test-error log %s", time.Now().String())
	// 封装后的日志是文件信息是logging程序包的位置,如何定位到调用方的位置[已修正]?
	// 实现时候改用函数func (l *Logger) Output(calldepth int, s string)
}
