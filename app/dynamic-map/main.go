package main

import "fmt"

func main() {
	m := map[int]int{
		1: 21,
		2: 22,
		3: 23,
	}
	counter := 0
	for k, v := range m {
		counter++
		m[len(m)+1] = 1
		fmt.Println("current map ", m)
		fmt.Println("current kv ", k, v)
	}
	fmt.Println("counter is ", counter)
}

// TODO:输出结果为什么会波动?
