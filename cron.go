// cron.go
package main

import (
	"log"

	"github.com/robfig/cron"
)

func main() {
	cron := cron.New()
	cron.AddFunc("*/1 * * * * *", func() { log.Println("log message") })
	//cron.Start()
	cron.Run()
	defer cron.Stop()
	//select {}
}
