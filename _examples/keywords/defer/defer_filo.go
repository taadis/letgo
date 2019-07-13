// defer_filo.go
// http://go.jsrun.net/Y9yKp
package main

import (
	"fmt"
)

func main() {
	foo()
}

func foo() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

// go run defer_filo.go
// 3
// 2
// 1
