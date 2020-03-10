package case1

import (
	"testing"
	"time"
)

// 现有一个计数器初始值为 0,
// 但有 10000 个人来一起累加.
// 会发生点什么?
func TestCase1(t *testing.T)  {
	counter := 0
	for i := 0; i < 10000; i++{
		go func() {
			counter++
		}()
	}
	time.Sleep(3 * time.Second)
	t.Logf("counter = %d", counter)
}

// 你会发现最后计数器的值总是小于 10000,
// 那么怎么就丢失了一部分数据呢?
// 究其原因:
// 协程并发执行时, 产生了竞态条件, 导致数据的乱入.
// 这也是我们通常所说的非线程安全 (这里用非协程安全或许更适合.)
