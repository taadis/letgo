// defer_defer.go
// http://go.jsrun.net/k9yKp
package main

import (
	"fmt"
)

func main() {
	foo()
}

func foo() {
	fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

// go run defer_defer.go
// 1
// 3
// 2
