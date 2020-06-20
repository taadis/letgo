// main.go
package main

import (
	"github.com/gin-gonic/gin"
)

// main
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Listening and serving HTTP on :8080
	r.Run()
}
