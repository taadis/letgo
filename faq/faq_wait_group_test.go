package faq

import (
	"net/http"
	"sync"
	"testing"
)

// wg.Add(-1)等价于wg.Done(),是不是还可以翻下源码看看.
func TestSyncWait4_equal(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(3)  // 添加3次计数
	wg.Done()  // 用wg.Done完成1次
	wg.Done()  // 用wg.Done完成2次
	wg.Add(-1) // 用wg.Add(-1)完成3次,这行注释掉,下面的wg.Wait()会一直等待
	wg.Wait()
}

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
