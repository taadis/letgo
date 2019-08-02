package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	guid := uuid.New()
	fmt.Println(guid)
}
