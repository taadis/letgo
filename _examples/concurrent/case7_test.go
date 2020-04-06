package concurrent

import (
	"fmt"
	"sync"
	"testing"
)

// 测试早餐店
func TestBreakfastShop(t *testing.T) {
	t.Log("早餐店开门啦.")
	// 声明一个信道
	ch := make(chan int, 1000)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go Boss(&wg, ch)
	go BossWife(&wg, ch)
	wg.Wait()
	t.Log("卖光啦, 早餐店收摊了.")
}

// 老板
// 在后厨造包子
func Boss(wg *sync.WaitGroup, ch chan int) {
	for i := 0; i < 1000; i++ {
		fmt.Printf("老板造出了第 %d 个包子\n", i)
		ch <- i
	}
	wg.Done()
}

// 老板娘
// 在前台卖包子
func BossWife(wg *sync.WaitGroup, ch chan int) {
	for i, _ := <-ch; i < 1000; i++ {
		fmt.Printf("老板娘卖出了第 %d 个包子\n", i)
	}
	wg.Done()
}
