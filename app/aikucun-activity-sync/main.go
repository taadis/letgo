package main

import (
	"log"

	"github.com/robfig/cron/v3"
)

func main() {
	log.Println("aikucun-activity-sync")
	c := cron.New()
	spec := "* * * * *"
	entryId, err := c.AddFunc(spec, ActivitiList)
	if err != nil {
		log.Println(".AddFunc error:", err.Error())
		return
	}
	log.Println("entryId:", entryId)
	c.Start()
	defer c.Stop()
	select {}
}

// ActivitiList
func ActivitiList() {
	log.Println("synching")
}
