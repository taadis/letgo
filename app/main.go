package main

import (
	"context"
	"fmt"
	"log"

	"github.com/taadis/letgo"
)

func main() {
	fmt.Printf("hello letgo!")
	ctx := context.Background()

	app := letgo.NewApp(
	//letgo.WithConfigName("config.yaml", &cfg),
	//WithConfig("config.yaml")
	)

	var cfg map[string]interface{}
	if err := app.Load(&cfg); err != nil {
		log.Fatal(err)
	}

	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
