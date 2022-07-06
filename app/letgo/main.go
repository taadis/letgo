// main.go
package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/taadis/letgo/handler/basic/platform"
	"github.com/taadis/letgo/handler/basic/shop"
	"github.com/taadis/letgo/handler/cron"
	"github.com/taadis/letgo/handler/security"
	systemUser "github.com/taadis/letgo/handler/system/user"
	userHandler "github.com/taadis/letgo/handler/user"
	"github.com/taadis/letgo/middleware/auth"
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

	systemGroup := r.Group("/system")
	{
		systemGroup.POST("/user/list", systemUser.List)
		systemGroup.POST("/user/change_password/", systemUser.ChangePassword)
		systemGroup.POST("/user/remove", systemUser.Remove)
	}

	basicGroup := r.Group("/basic")
	{
		basicGroup.POST("/platform/add", platform.Add)
		basicGroup.POST("/platform/list", platform.List)
		basicGroup.POST("/platform/detail", platform.Detail)
		basicGroup.POST("/platform/update", platform.Update)
		basicGroup.POST("/platform/remove", platform.Remove)

		basicGroup.POST("/shop/list", shop.List)
		basicGroup.POST("/shop/remove", shop.Remove)
	}

	//
	securityGroup := r.Group("/security")
	{
		securityGroup.POST("/encrypt/md5", security.EncryptMD5)
	}

	return r
}

// main
func main() {
	r := setupRouter()

	// Listening and serving HTTP on :8080
	r.Run()
}
