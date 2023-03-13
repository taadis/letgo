package tests

import (
	"log"
	"testing"
)

func hello(c chan string) {
	name := <-c // 从通道获取数据
	log.Printf("Hello %s", name)
}

// TestChannel1 一个简单 channel 示例
func TestChannel1(t *testing.T) {
	// 声明一个字符串类型的变量
	ch := make(chan string)

	// 开启一个 goroutine
	go hello(ch)

	// 发送数据到通道 ch
	ch <- "Gophers"

	// 关闭通道
	close(ch)
}
