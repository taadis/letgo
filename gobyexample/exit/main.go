package main

import (
	"fmt"
	"os"
)

func main() {
	// 当使用 `os.Exit` 退出程序时, 这里的 `defer` 将不会执行
	// 因此这里的 `fmt.Println("!")` 不会被调用并打印出来
	defer fmt.Println("!")

	// 退出程序,并且指定状态码为3
	os.Exit(3)
}

// ## os.Exit
// 使用 `os.Exit` 可以让程序以指定的状态码退出.
// 通常,状态码0表示成功,非0表示错误.
// 程序会立即终止,因此延迟函数不会执行.
//
// 对了,为了便于移植程序,状态码最好在[0,125]范围内.

// https://go.dev/play/p/Rq6h9DcHoFX
