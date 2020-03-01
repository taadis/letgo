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

// 测试 map 的取值操作
func TestMapGet(t *testing.T) {
	m := map[string]int{}
	t.Log(m)
	value := m["0"] // 取值时没有对应的key, 则取值为类型对应的默认零值
	t.Log(value)
	// 但有个问题是当我们设置了一个 key, 赋值为零值时,
	// 如何区分是有key 还是没有key呢?
	value, ok := m["0"]
	if ok != true {
		t.Log("not key")
	} else {
		t.Log("has key")
	}
}

// 测试 map 的长度
func TestMapLength(t *testing.T) {
	m := map[string]int{}
	t.Log(m, len(m))
	m1 := map[string]int{"2": 2}
	t.Log(m1, len(m1))
}
