package main

import (
	"github.com/astaxie/beego/logs"
)

func init() {
	logs.SetLogger(logs.AdapterConsole, `{"level":7,"color":true}`)
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/xxx.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
}

func main() {
	logs.Info("beego logs.Info() message")
}
