package map_test

import (
	"testing"
)

// 测试 map 的声常用明写法
func TestMapDefine(t *testing.T) {
	m := map[string]int{}
	t.Log(m)
	m1 := map[string]int{"1": 1, "2": 2, "3": 3}
	t.Log(m1)
	m2 := make(map[string]int, 10)
	t.Log(m2)
}

// 测试 map 的赋值操作
func TestMapSet(t *testing.T) {
	m := map[string]int{}
	t.Log(m)
	m["1"] = 100
	t.Log(m)
	m["1"] = 200
	t.Log(m)
}
