package main

import (
	"log"

	"github.com/astaxie/beego/logs"
	"github.com/kardianos/service"
)

func init() {
	logs.SetLogger(logs.AdapterConsole, `{"level":7,"color":true}`)
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/xxx.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
}

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	// Do work here
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "CustomService",
		DisplayName: "Custom Service DisplayName",
		Description: "Custom Service Description",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
