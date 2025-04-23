package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	// 快速步进限制
	seqFastStepLimit = 200
)

var (
	// 命令行选项
	formatStr  string
	separator  string
	equalWidth bool
)

func init() {
	// 定义命令行参数
	flag.StringVar(&formatStr, "f", "", "使用printf风格的浮点数格式")
	flag.StringVar(&formatStr, "format", "", "使用printf风格的浮点数格式")
	flag.StringVar(&separator, "s", "\n", "使用指定字符串分隔数字")
	flag.StringVar(&separator, "separator", "\n", "使用指定字符串分隔数字")
	flag.BoolVar(&equalWidth, "w", false, "使用前导零使所有数字等宽")
	flag.BoolVar(&equalWidth, "equal-width", false, "使用前导零使所有数字等宽")
}

// 操作数结构体
type operand struct {
	value     float64 // 数值
	width     int     // 打印宽度
	precision int     // 小数点后精度
}

// 解析命令行参数为操作数
func parseArg(arg string) (operand, error) {
	ret := operand{
		width:     0,
		precision: math.MaxInt32,
	}

	// 去除前导空格和加号
	arg = strings.TrimLeft(arg, " +")

	// 解析数值
	val, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		return ret, fmt.Errorf("无效的浮点数参数: %s", arg)
	}
	ret.value = val

	// 检查是否为NaN
	if math.IsNaN(val) {
		return ret, fmt.Errorf("无效的not-a-number参数: %s", arg)
	}

	// 对于整数,不需要精度
	if !strings.Contains(arg, ".") && !strings.Contains(arg, "p") {
		ret.precision = 0
	}

	// 自动设置宽度和精度
	if !strings.ContainsAny(arg, "xX") && !math.IsInf(val, 0) && !math.IsNaN(val) {
		ret.width = len(arg)
		if idx := strings.Index(arg, "."); idx >= 0 {
			fractionLen := len(arg) - idx - 1
			if strings.ContainsAny(arg, "eE") {
				fractionLen = strings.IndexAny(arg[idx+1:], "eE")
			}
			if fractionLen <= math.MaxInt32 {
				ret.precision = fractionLen
			}
		}
	}

	return ret, nil
}

// 获取默认格式
func getDefaultFormat(first, step, last operand) string {
	prec := first.precision
	if step.precision > prec {
		prec = step.precision
	}

	if prec != math.MaxInt32 && last.precision != math.MaxInt32 {
		if equalWidth {
			firstWidth := first.width + (prec - first.precision)
			lastWidth := last.width + (prec - last.precision)
			width := firstWidth
			if lastWidth > width {
				width = lastWidth
			}
			return fmt.Sprintf("%%0%d.%df", width, prec)
		}
		return fmt.Sprintf("%%.%df", prec)
	}

	return "%g"
}

// 打印数字序列
func printNumbers(format string, first, step, last float64) {
	outOfRange := (step < 0 && first < last) || (step > 0 && last < first)
	if outOfRange {
		return
	}

	for x := first; ; x += step {
		fmt.Printf(format, x)

		if (step < 0 && x <= last) || (step > 0 && x >= last) {
			break
		}
		fmt.Print(separator)
	}
	fmt.Println()
}

// 快速打印整数序列
func seqFast(start, end int64, step uint64) {
	for i := start; i <= end; i += int64(step) {
		fmt.Print(i)
		if i+int64(step) <= end {
			fmt.Print(separator)
		}
	}
	fmt.Println()
}

func main() {
	flag.Parse()

	// 解析位置参数
	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "缺少操作数")
		os.Exit(1)
	}
	if len(args) > 3 {
		fmt.Fprintf(os.Stderr, "多余的操作数 %s\n", args[3])
		os.Exit(1)
	}

	// 设置默认值
	first := operand{value: 1}
	step := operand{value: 1}
	var last operand
	var err error

	// 解析参数
	switch len(args) {
	case 1:
		last, err = parseArg(args[0])
	case 2:
		first, err = parseArg(args[0])
		if err == nil {
			last, err = parseArg(args[1])
		}
	case 3:
		first, err = parseArg(args[0])
		if err == nil {
			step, err = parseArg(args[1])
		}
		if err == nil {
			if step.value == 0 {
				fmt.Fprintf(os.Stderr, "无效的零增量值: %s\n", args[1])
				os.Exit(1)
			}
			last, err = parseArg(args[2])
		}
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// 检查格式和等宽是否冲突
	if formatStr != "" && equalWidth {
		fmt.Fprintln(os.Stderr, "打印等宽字符串时不能指定格式字符串")
		os.Exit(1)
	}

	// 尝试使用快速整数方法
	if formatStr == "" && len(separator) == 1 && !equalWidth &&
		first.precision == 0 && step.precision == 0 && last.precision == 0 &&
		step.value > 0 && step.value <= seqFastStepLimit {
		startInt := int64(first.value)
		endInt := int64(last.value)
		stepInt := uint64(step.value)
		if float64(startInt) == first.value && float64(endInt) == last.value {
			seqFast(startInt, endInt, stepInt)
			return
		}
	}

	// 使用浮点数方法
	format := formatStr
	if format == "" {
		format = getDefaultFormat(first, step, last)
	}

	printNumbers(format, first.value, step.value, last.value)
}
