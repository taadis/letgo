package main

import "log"

type HelloJob struct {
	Name string
}

func (h HelloJob) Run() {
	log.Printf("hello %s", h.Name)
}
