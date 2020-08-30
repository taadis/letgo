package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/robfig/cron/v3"
)

// main
func main() {
	c := cron.New()
	spec := "* * * * *"
	entryId, err := c.AddFunc(spec, Send)
	if err != nil {
		log.Fatalln("cron.AddFunc error:", err.Error())
	}
	log.Println("entryId:", entryId)
	c.Start()
	defer c.Stop()
	select {}
}

// print
func print() {
	log.Println("waterd")
}

// Send
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
