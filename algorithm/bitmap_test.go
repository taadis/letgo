package algorithm

import "testing"

func TestBitmap_AddAndHas(t *testing.T) {
	num := 40 * 10000 * 10000
	bitmap := new(Bitmap)
	bitmap.Add(num)
	has := bitmap.Has(num)
	t.Logf("has:%v", has)
}
