package faq

import (
	"fmt"
	"testing"
	"time"
)

// channel+select的应用-吃饭睡觉打豆豆
// 如何让协程合理退出?使用time.After来控制超时退出
func TestChannel2_EatSleepBeatDouDou(t *testing.T) {
	eatCh := make(chan string)
	go func(chan string) {
		time.Sleep(time.Second * 2)
		eatCh <- "call eating"
	}(eatCh)
	sleepCh := time.NewTimer(time.Second * 3)

	for {
		select {
		case s := <-eatCh:
			t.Logf("%s, to eat", s)
		case <-sleepCh.C:
			t.Logf("to sleep")
		case <-time.After(time.Second * 5):
			t.Logf("timeout break")
			return
			//default:
			//	t.Logf("beat dou dou...")
		}
	}
}

// 场景1:某逻辑使用了多协程处理,要求主动控制每个协程的超时时间,不能一直等着.
func TestTimeout1(t *testing.T) {
	ch := make(chan string)
	quit := make(chan bool)

	for i := 0; i < 10; i++ {
		go func(i int) {
			// 表示业务逻辑处理越来越耗时耗时
			time.Sleep(time.Second * time.Duration(i))
			ch <- fmt.Sprintf("done%d", i)
		}(i)
	}

	go func() {
		for {
			select {
			case res := <-ch:
				t.Logf(res)
			case <-time.After(time.Second * 3):
				t.Logf("timeout")
				quit <- true
			case <-quit:
				return
			}
		}
	}()

	quit <- true

	t.Logf("func done")
}
