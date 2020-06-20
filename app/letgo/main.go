// main.go
package main

import (
	userHandler "gitee.com/taadis/letgo/handler/user"
	"gitee.com/taadis/letgo/middleware/auth"
	"github.com/gin-gonic/gin"
)

// setupRouter
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.POST("/user/register", userHandler.Register)
	r.POST("/user/login", userHandler.Login)
	r.POST("/user/profile", auth.AuthMiddleware(), userHandler.Profile)
	return r
}

// main
func main() {
	r := setupRouter()

	// Listening and serving HTTP on :8080
	r.Run()
}
