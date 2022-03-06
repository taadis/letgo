package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomA() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			rand.Seed(time.Now().UnixNano())
			ch <- rand.Int()
		}
	}()
	return ch
}

// GenerateRandomB 生成1000以内的随机数，以做区分
func GenerateRandomB() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			rand.Seed(time.Now().UnixNano())
			ch <- rand.Intn(1000)
		}
	}()
	return ch
}

func GenerateRandom() chan int {
	ch := make(chan int, 20)
	go func() {
		for {
			// 使用select的fan-in来增加生成器的随机源
			select {
			case ch <- <-GenerateRandomA():
			case ch <- <-GenerateRandomB():
			}
		}
	}()
	return ch
}

func main() {
	ch := GenerateRandom()
	for random := range ch {
		fmt.Printf("%d\n", random)
	}
}
