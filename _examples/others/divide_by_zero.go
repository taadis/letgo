package main

import (
	"fmt"
)

func main() {
	a, b := 1, 0
	c := a / b
	fmt.Println(a, b, c)
}

// 被除数为0时, 程序会直接崩溃, 我们应该提前判断被除数是否等于0来规避可能的问题.
