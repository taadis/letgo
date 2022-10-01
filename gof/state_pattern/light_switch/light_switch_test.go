package light_switch

import "testing"

// 通过测试来模拟如何使用
func TestLightSwitch(t *testing.T) {
	s := NewLightSwitch()
	s.switchOn()
	s.switchOn()
	s.switchOff()
}
