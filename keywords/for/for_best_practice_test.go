// for_best_practice_test.go
package main

import (
	"fmt"
	"testing"
)

// func main() {
// 	list := [5]int{1, 3, 5, 7, 9}
// 	// for i := 0; i < len(list); i++ {
// 	// 	fmt.Println(list[i])
// 	// }
// 	forInLen(&list)
// 	forRange(&list)
// }

func TestForInLen(t *testing.T) {
	list := [5]int{1, 3, 5, 7, 9}
	for i := 0; i < len(list); i++ {
		//fmt.Println(list[i])
		t.Log(list[i])
	}
}

func TestForRange(t *testing.T) {
	list := [5]int{1, 3, 5, 7, 9}
	for _, value := range list {
		fmt.Println(t.Name(), value)
	}
}

// go run for_best_practice_test.go
//
