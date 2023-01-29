package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://go.dev")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}
}

// 释义:
// 向服务端发送一个HTTP GET请求.
// `http.Get()` 本质是创建一个`http.Client`并调用Get方法的快捷方式,
// 它使用了`http.DefaultClient`对象以及默认设置.
//
// 请求完成后,打印HTTP response状态.
// 并打印响应体前5行内容
//
// 在线示例:https://go.dev/play/p/oNxsIz5uq4e
