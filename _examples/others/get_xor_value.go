package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%#x", 0xFFD8^0x0D2A)
}

// 参考
// [Go语言fmt包Printf方法详解](https://www.jianshu.com/p/8be8d36e779c)
