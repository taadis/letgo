package concurrent

import (
	"sync"
	"testing"
)

// 测试早餐店
func TestBreakfastShop(t *testing.T) {
	t.Log("早餐店开门啦.")
	// 声明一个信道
	ch := make(chan int, 100)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// 老板在后厨造包子
	go func(waitGroup *sync.WaitGroup, channel chan int) {
		for i := 1; i <= 100; i++ {
			t.Logf("老板造出了第 %d 个包子\n", i)
			channel <- i
		}
		waitGroup.Done()
	}(&wg, ch)

	// 老板娘在前台卖包子
	go func(waitGroup *sync.WaitGroup, channel chan int) {
		for i, _ := <-channel; i <= 100; i++ {
			t.Logf("老板娘卖出了第 %d 个包子", i)
		}
		waitGroup.Done()
	}(&wg, ch)

	wg.Wait()
	t.Log("卖光啦, 早餐店收摊了.")
}
