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

	app := letgo.NewApp()

	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
