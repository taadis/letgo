package algorithm

import (
	"math/rand"
	"testing"
	"time"
)

// 测试插入排序
func TestInsertSort(t *testing.T) {
	// 生成随机数组
	rand.Seed(time.Now().UnixNano())
	var slice []int
	for i := 0; i < 9; i++ {
		slice = append(slice, rand.Intn(100))
	}
	t.Log("insert sort:", slice)
	slice = InsertSort(slice)
	t.Log("insert sort:", slice)
}
