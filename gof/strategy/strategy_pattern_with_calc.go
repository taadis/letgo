package strategy

// 一般思路,根据传递参数的不通，使用条件分支语句进行穷举操作实现
func calc(a int, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("invalid op")
	}
	// 但是当需要扩展更多运算时,需要在这里不断的补充条件分支语句
	// 此逻辑会越来越长,并且会越来越复杂难以维护
	return 0
}

type Calc interface {
	Calc(a int, b int) int
}

type Calculator struct {
	Op string
	Calc
}

type CalcAdd struct{}

func (c *CalcAdd) Calc(a int, b int) int {
	return a + b
}

type CalcSub struct{}

func (c *CalcSub) Calc(a int, b int) int {
	return a - b
}

type CalcMul struct{}

func (c *CalcMul) Calc(a int, b int) int {
	return a * b
}

type CalcDiv struct{}

func (c *CalcDiv) Calc(a int, b int) int {
	return a / b
}
