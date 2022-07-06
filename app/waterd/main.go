package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
)

// main
func main() {
	// [spec]func
	// e.g.[* * * * * *]print
	jobs := make(map[string]func())
	jobs["* * * * * *"] = print
	jobs["*/5 * * * * *"] = fooPanic
	jobs["*/5 * * * * *"] = fooSkip

	logger := cron.VerbosePrintfLogger(log.Default())
	c := cron.New(
		cron.WithChain(
			cron.Recover(logger),
			cron.SkipIfStillRunning(logger),
		),
		cron.WithSeconds(),
		cron.WithLogger(logger),
	)

	for spec, job := range jobs {
		entryId, err := c.AddFunc(spec, job)
		if err != nil {
			log.Fatalln("cron.AddFunc error:", err.Error())
		}
		log.Println("added cron func entryId:", entryId)
		c.Start()
	}

	defer c.Stop()

	waitExit()
}

// print
func print() {
	log.Println("waterd")
}

// cron调度时上一个job没执行完,下一个又开始跑了?跳过
func fooSkip() {
	log.Println("fooSkip start...")
	time.Sleep(10 * time.Second)
	log.Println("fooSkip ...end")
}

// cron调度的goroutine发生panic时会导致整个程序挂掉?
// 可以通过cron.Recover统一来处理
func fooPanic() {
	panic("some panic")
	log.Println("panic log")
}

func waitExit() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case s := <-signalCh:
			switch s {
			case syscall.SIGINT:
				fallthrough
			case syscall.SIGTERM:
				fallthrough
			default:
				log.Print("stop signal received, stop all scheduler")
				log.Print("gracefully exit")
				return
			}
		}
	}
}

func Send() {
	webhook := "https://oapi.dingtalk.com/robot/send?access_token=xxx"
	url := webhook
	body := strings.NewReader(`{"msgtype": "text","text": {"content": "我就是我, 是不一样的烟火"}}`)
	req, err := http.NewRequest(
		http.MethodPost, // method string,
		url,             // url string,
		body,            // body io.Reader,
	)
	if err != nil {
		log.Fatalln("http.NewRequest error:", err.Error())
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("http.DefaultClient error:", err.Error())
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("ioutil.ReadAll error:", err.Error())
	}
	log.Println("respBody:", string(respBody[:]))
}
