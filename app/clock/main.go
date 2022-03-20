package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建ticker,并设置多长时间触发一次
	ticker := time.NewTicker(time.Second * 1)
	for t := range ticker.C {
		//fmt.Printf("\r\033[know time %v", t.Local().String()) // \r\033]k is ok
		fmt.Printf("now time %v\r", t.Local().String()) // \r is ok
	}
}
