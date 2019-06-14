// defer_parameters.go
// http://go.jsrun.net/BUyKp
package main

import (
	"fmt"
)

func main() {
	foo()
}

func foo() {
	number := 1
	defer fmt.Println(number)
	number = 2
	return
}
