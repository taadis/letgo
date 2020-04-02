package case5

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestFirstResult(t *testing.T) {
	t.Logf("NumGoroutine: %d", runtime.NumGoroutine())
	t.Log(ping())
	//time.Sleep(3 * time.Second)
	t.Logf("NumGoroutine: %d", runtime.NumGoroutine())
}

func ping() string {
	ch := make(chan string, 1)
	go pingBaidu(ch)
	go pingBing(ch)
	go pingGoogle(ch)
	return <-ch
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

func runTask(i int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("the result is from %d", i)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			task := runTask(i)
			ch <- task
		}(i)
	}
	return <-ch
}

func TestFirst(t *testing.T) {
	fmt.Println("Go num: ", runtime.NumGoroutine())
	FirstResponse()
	time.Sleep(5 * time.Second)
	fmt.Println("Go num: ", runtime.NumGoroutine())
}
