package faq

import (
	"fmt"
	"testing"
	"time"
)

// 调用方负责调度多方整个流水线,channel的流水线简单示例,
// 实际应用中channel里的int类型数据可以按需替换.
func TestChannelPipeline(t *testing.T) {
	ch1 := producer(1, 3, 5, 7, 9)
	ch2 := compute(ch1)
	done := consumer(ch2)
	<-done
	//time.Sleep(time.Second * 3)
}

// producer 负责生产数据,并把数据写入channel
func producer(nums ...int) <-chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)
		for _, n := range nums {
			time.Sleep(time.Second)
			outCh <- n
		}
	}()
	return outCh
}

// compute 负责计算数据(加法计算),并将结果写入channel
func compute(inCh <-chan int) <-chan int {
	outCh := make(chan int)
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
