package case5

import (
	"runtime"
	"testing"
	"time"
)

func TestFirstResult(t *testing.T) {
	ch := make(chan string, 5)
	go pingBaidu(ch)
	go pingBing(ch)
	go pingGoogle(ch)
	t.Logf("NumGoroutine: %d", runtime.NumGoroutine())
	t.Log(<-ch)
	t.Logf("NumGoroutine: %d", runtime.NumGoroutine())
}

func pingBaidu(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "from baidu"
}

func pingBing(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "from bing"
}

func pingGoogle(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from google"
}
