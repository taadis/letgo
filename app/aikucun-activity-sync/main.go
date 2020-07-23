package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func main() {
	fmt.Println("aikucun-activity-sync")
	c := cron.New()
	spec := "* * * * *"
	entryId, err := c.AddFunc(spec, ActivitiList)
	if err != nil {
		fmt.Println(".AddFunc error:", err.Error())
		return
	}
	fmt.Println("entryId:", entryId)
	c.Start()
	defer c.Stop()
	select {}
}

// ActivitiList
func ActivitiList() {
	fmt.Println("synching")
}
