// defer_impact.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo())
}

func foo() (result int) {
	defer func() {
		result++
	}()
	return 0
}

// go run defer_defer.go
// 1
