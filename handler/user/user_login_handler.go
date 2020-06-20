package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserLogin
type UserLogin struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// Login
func Login(ctx *gin.Context) {
	payload := &UserLogin{}
	err := ctx.Bind(payload)
	if err != nil {
		log.Println("ctx.Bind error", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	//
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": fmt.Sprint(payload.UserName, payload.Password),
	})
}
