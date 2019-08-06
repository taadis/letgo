// 数组和切片
package main

import (
	"fmt"
	"runtime"
)

func main() {
	//pc, file, line, ok := runtime.Caller(0)
	_, file, line, ok := runtime.Caller(0)
	fmt.Println(file, line, ok)

	// 几种定义数组的方式
	var arr01 [3]int = [3]int{1, 3, 5}
	fmt.Println("arr01= ", arr01)

	var arr02 = [3]int{2, 4, 6}
	fmt.Println("arr02=", arr02)

	var arr03 = [...]int{3, 6, 9}
	fmt.Println("arr03=", arr03)

	var arr04 = [...]int{1: 4, 0: 2, 2: 8}
	fmt.Println("arr04=", arr04)

	arr05 := [...]string{"jack", "tom", "lee"}
	fmt.Println("arr05=", arr05)
}
