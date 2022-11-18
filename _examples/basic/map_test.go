package basic

import "testing"

func TestMap(t *testing.T) {
	m := make(map[string]struct{})
	m["1"] = struct{}{}
	m["2"] = struct{}{}
	m["3"] = struct{}{}
	t.Logf("0befor map %+v %p", m, &m)
	UpdateMap(t, m)
	t.Logf("0after map %+v %p", m, &m)
}

func UpdateMap(t *testing.T, m map[string]struct{}) {
	t.Logf("1after map %+v %p", m, &m)
	m["4"] = struct{}{}
	t.Logf("1after map %+v %p", m, &m)
}
