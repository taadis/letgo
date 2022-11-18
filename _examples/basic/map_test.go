package basic

import "testing"

func TestMap(t *testing.T) {
	m := make(map[string]struct{})
	m["1"] = struct{}{}
	m["2"] = struct{}{}
	m["3"] = struct{}{}
	t.Logf("befor map %+v", m)
	UpdateMap(m)
	t.Logf("after map %+v", m)
}

func UpdateMap(m map[string]struct{}) {
	m["4"] = struct{}{}
}
