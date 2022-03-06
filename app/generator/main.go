package main

import (
	"fmt"
	"math/rand"
)

func GenerateRandomA() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func main() {
	ch := GenerateRandomA()
	fmt.Printf("%d\n", <-ch)
	fmt.Printf("%d\n", <-ch)
}
