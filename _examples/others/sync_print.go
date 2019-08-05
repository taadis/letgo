// Q: 使用两个goroutine交替打印序列，一个goroutinue打印数字，另一个goroutine打印字母，最终结果如下: 12AB34CD56EF78GH910IJ
// 参考: https://www.cnblogs.com/zhangjinfu/p/11288472.html 网上答案很多, 这个是比较简洁明了滴
package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)

	go func() {
		for i := 1; i < 10; i += 2 {
			fmt.Print(i)
			fmt.Print(i + 1)
			c <- true
			<-c
		}
	}()

	for i := 'A'; i < 'J'; i += 2 {
		<-c
		fmt.Printf("%c", i)
		fmt.Printf("%c", i+1)
		c <- true
	}
}
