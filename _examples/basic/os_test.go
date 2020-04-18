package basic

import (
	"os"
	"testing"
)

// 查看环境变量列表
func TestEnv1(t *testing.T) {
	for i, env := range os.Environ() {
		t.Log(i, "=", env)
	}
}
