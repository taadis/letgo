package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建ticker,并设置多长时间触发一次
	ticker := time.NewTicker(time.Second * 1)
	for t := range ticker.C {
		fmt.Printf("now time %v\n", t.Local().String())
	}
}
