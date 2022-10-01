package state_pattern

import "fmt"

// TrafficLight 交通信号灯
type TrafficLight struct {
	state State
}

func NewTrafficLight() *TrafficLight {
	l := new(TrafficLight)
	l.state = new(RedState) // 初始化及默认值
	return l
}

func (l *TrafficLight) setState(state State) {
	l.state = state
}

func (l *TrafficLight) switchToGreen() {
	l.state.switchToGreen(l)
}

func (l *TrafficLight) switchToYellow() {
	l.state.switchToYellow(l)
}

func (l *TrafficLight) switchToRed() {
	l.state.switchToRed(l)
}

// State 接口定义了3个标准方法
// 注意每个方法的参数是交通灯TrafficLight
type State interface {
	switchToGreen(light *TrafficLight)
	switchToYellow(light *TrafficLight)
	switchToRed(light *TrafficLight)
}

// RedState 红灯状态及方法实现
type RedState struct {
}

func (r *RedState) switchToGreen(light *TrafficLight) {
	fmt.Println("error:红灯不能直接切换为绿灯")
}

func (r *RedState) switchToYellow(light *TrafficLight) {
	light.setState(new(YellowState))
	fmt.Println("黄灯亮了")
}

func (r *RedState) switchToRed(light *TrafficLight) {
	fmt.Println("error:已经是红灯状态无需重复切换")
}

// YellowState 黄灯状态及方法实现
type YellowState struct {
}

func (r *YellowState) switchToGreen(light *TrafficLight) {
	light.setState(new(GreenState))
	fmt.Println("绿灯亮了")
}

func (r *YellowState) switchToYellow(light *TrafficLight) {
	fmt.Println("error:已经是黄灯状态无需重复切换")
}

func (r *YellowState) switchToRed(light *TrafficLight) {
	light.setState(new(RedState))
	fmt.Println("红色亮了")
}

// GreenState 绿灯状态及方法实现
type GreenState struct {
}

func (r *GreenState) switchToGreen(light *TrafficLight) {
	fmt.Println("error:已经是绿灯状态无需重复切换")
}

func (r *GreenState) switchToYellow(light *TrafficLight) {
	light.setState(new(YellowState))
	fmt.Println("黄灯亮了")
}

func (r *GreenState) switchToRed(light *TrafficLight) {
	fmt.Println("error:绿灯不能直接切换为红灯")
}
