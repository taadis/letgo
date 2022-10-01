package light_switch

import "fmt"

type LightSwitch struct {
	state State
}

func NewLightSwitch() *LightSwitch {
	s := new(LightSwitch)
	s.state = new(OffState) // 初始化及默认状态
	return s
}

func (s *LightSwitch) setState(state State) {
	s.state = state
}

func (s *LightSwitch) switchOn() {
	s.state.SwitchOn(s)
}

func (s *LightSwitch) switchOff() {
	s.state.SwitchOff(s)
}

// State 接口定义了灯状态的2个标准方法
// 注意参数是灯开关LightSwitch本身
type State interface {
	SwitchOn(lightSwitch *LightSwitch)
	SwitchOff(lightSwitch *LightSwitch)
}

// OnState 开状态及方法实现
type OnState struct {
}

func (s *OnState) SwitchOn(*LightSwitch) {
	fmt.Println("error:已经是开状态无需重复打开")
}

func (s *OnState) SwitchOff(lightSwitch *LightSwitch) {
	lightSwitch.setState(new(OffState))
	fmt.Println("灯关了")
}

// OffState 关状态及方法实现
type OffState struct {
}

func (s *OffState) SwitchOn(lightSwitch *LightSwitch) {
	lightSwitch.setState(new(OnState))
	fmt.Println("灯开了")
}

func (s *OffState) SwitchOff(lightSwitch *LightSwitch) {
	fmt.Println("error:已经是关闭状态无需重复关闭")
}
