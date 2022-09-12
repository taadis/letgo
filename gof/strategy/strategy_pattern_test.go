package strategy

import "testing"

func TestCalculator1(t *testing.T) {
	calculator := new(Calculator)

	calculator.Calc = new(CalcAdd)
	c := calculator.Calc.Calc(1, 2)
	t.Log(c)

	calculator.Calc = new(CalcMul)
	c = calculator.Calc.Calc(2, 5)
	t.Log(c)

	// 这里有个问题,调用方使用的时候还是需要通过条件分支语句来判断用哪个Calc策略的实现
	// 这里可以用map映射来封装一下

}

func TestCalculatorWithTableDriven(t *testing.T) {
	// 表驱动
	m := make(map[string]Calc)
	m["+"] = new(CalcAdd)
	m["-"] = new(CalcSub)
	m["*"] = new(CalcMul)
	m["/"] = new(CalcDiv)

	calculator := new(Calculator)
	calculator.Calc = m["+"]
	c := calculator.Calc.Calc(1, 2)
	t.Log(c)

	// 这样还需要调用方来维护这个表,可以再封装到上下文构造函数或Set函数中
}
