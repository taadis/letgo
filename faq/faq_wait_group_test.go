package faq

import (
	"net/http"
	"sync"
	"testing"
)

// HTTP请求不同耗时的url,来验证sync.WaitGroup的使用
func TestSyncWait3_urls(t *testing.T) {
	wg := &sync.WaitGroup{}

	urls := []string{
		"https://baidu.com",
		"https://google.com",
		"https://bing.com",
	}

	for _, url := range urls {
		wg.Add(1)
		go func(wg *sync.WaitGroup, url string) {
			defer wg.Done()

			resp, err := http.Get(url)
			if err != nil {
				t.Logf("HTTP GET %s -> error:%v", url, err)
				return
			}
			t.Logf("HTTP GET %s -> resp:%v, error:%v", url, resp.StatusCode, err)
		}(wg, url)
	}

	wg.Wait()
	t.Logf("finished")
}

// 如果wg.Add调用维护的计数更改成了负数,将产生panic
func TestSyncWaitGroupPanic1(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(-1)
}

// 如果wg.Done()的调用维护的计数变成了负数,将产生panic
func TestSyncWaitGroupPanic2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 11; i++ {
		wg.Done()
	}
}
