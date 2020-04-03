package concurrent

import (
	"testing"
	"time"
)

func TestCall(t *testing.T) {
	ch := make(chan string, 1)
	go callBaidu(ch)
	go callBing(ch)
	go callGoogle(ch)
	result := <-ch
	t.Log(result)
}

func callBaidu(ch chan string) {
	time.Sleep(time.Second)
	ch <- "from baidu"
}

func callBing(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "from bing"
}

func callGoogle(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from google"
}
