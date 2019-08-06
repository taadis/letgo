package main

import "fmt"

func main() {
	num, str, err := foo()
	if err != nil {
		panic(err)
	}
	fmt.Println(num, str)
}

// 返回3个不同类型的值
func foo() (num int, str string, err error) {
	num = 123
	str = "四五六"
	err = nil
	return
}
