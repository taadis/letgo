package state_pattern

import "testing"

// 使用测试来模拟如何使用
func TestTrafficLightState(t *testing.T) {
	light := NewTrafficLight()
	light.switchToYellow()
	light.switchToGreen()
	light.switchToYellow()
	light.switchToRed()
}
