package faq

import (
	"fmt"
	"testing"
)

// 调用方负责调度多方整个流水线,channel的流水线简单示例,
// 实际应用中channel里的int类型数据可以按需替换.
// 这里的流水线彼此间还比较同步,可以调整不同channel的缓冲大小以及生产/计算/消费速度来模拟更多实际情况.
func TestChannelPipeline(t *testing.T) {
	// nums := []int{1, 3, 5, 7, 9}
	nums := make([]int, 0)
	for i := 0; i < 10000; i++ {
		nums = append(nums, i)
	}
	ch1 := producer(nums...)
	ch2 := compute(ch1)
	done := consumer(ch2)
	<-done
	//time.Sleep(time.Second * 3)
}

// producer 负责生产数据,并把数据写入channel
func producer(nums ...int) <-chan int {
	outCh := make(chan int, 10)
	go func() {
		defer close(outCh)
		for _, n := range nums {
			//time.Sleep(time.Second)
			outCh <- n
		}
	}()
	return outCh
}

// compute 负责计算数据(加法计算),并将结果写入channel
func compute(inCh <-chan int) <-chan int {
	outCh := make(chan int, 5)
	go func() {
		defer close(outCh)
		for num := range inCh {
			outCh <- num + num
		}
	}()
	return outCh
}

// consumer 负责从channel中取出数据,打印输出
func consumer(inCh <-chan int) <-chan int {
	done := make(chan int)
	go func() {
		defer close(done)
		for num := range inCh {
			fmt.Printf("print num:%d\n", num)
		}
	}()
	return done
}
