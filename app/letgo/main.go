// main.go
package main

import (
	"gitee.com/taadis/letgo/handler/basic/platform"
	"gitee.com/taadis/letgo/handler/cron"
	userHandler "gitee.com/taadis/letgo/handler/user"
	"gitee.com/taadis/letgo/middleware/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// setupRouter
func setupRouter() *gin.Engine {
	r := gin.Default()

	// cors
	//r.Use(cors.Default())
	corsConf := cors.DefaultConfig()
	corsConf.AddAllowHeaders("Authorization")
	corsConf.AllowAllOrigins = true
	r.Use(cors.New(corsConf))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.POST("/user/register", userHandler.Register)
	r.POST("/user/login", userHandler.Login)
	r.POST("/user/profile", auth.AuthMiddleware(), userHandler.Profile)

	r.POST("/cron/create", cron.Create)

	basicGroup := r.Group("/basic")
	{
		basicGroup.POST("/platform/add", platform.Add)
		basicGroup.GET("/platform/list", platform.List)
		basicGroup.POST("/platform/remove", platform.Remove)
	}

	return r
}

// main
func main() {
	r := setupRouter()

	// Listening and serving HTTP on :8080
	r.Run()
}
