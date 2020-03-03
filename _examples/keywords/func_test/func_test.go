package func_test

import (
	"testing"
)

// 定义一个可变长参数函数
func sum(op ...int) int {
	ret := 0
	for _, value := range op {
		ret += value
	}
	return ret
}

// 测试可变长参数函数
func TestSum(t *testing.T) {
	t.Log(sum(1), sum(1, 2), sum(1, 2, 3))
}
