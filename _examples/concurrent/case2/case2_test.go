package case2

import "testing"

func TestCase2(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	for i := 0; i < 9; i++ {
		go func(i int) {
			ch1 <- i
		}(i)
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			ch2 <- i
		}(i)
	}

	for {
		select {
		case x1 := <-ch1:
			t.Logf("x1 = %d", x1)
		case x2 := <-ch2:
			t.Logf("x2 = %d", x2)
		default:
			t.Logf("break")
			break
		}
	}
}

// 没有 default 时报错如下:
// fatal error: all goroutines are asleep - deadlock!
// 有 default 时, 执行测试, 内存会一直增加, 而没有任何输出!
// 究其原因:
//
