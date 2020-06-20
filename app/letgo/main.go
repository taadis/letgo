// main.go
package main

import (
	"github.com/gin-gonic/gin"
)

// setupRouter
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

// main
func main() {
	r := setupRouter()

	// Listening and serving HTTP on :8080
	r.Run()
}
