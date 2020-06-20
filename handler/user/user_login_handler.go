package user

import (
	"fmt"
	"log"
	"net/http"

	"gitee.com/taadis/letgo/store"
	"gitee.com/taadis/letgo/util/jwt"
	"gitee.com/taadis/letgo/util/security"
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
	user := &store.SystemUser{}
	err = store.Db.Where("name=?", payload.UserName).First(user).Error
	if err != nil {
		log.Println("db.First error", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	// check password
	plaintext := payload.Password + user.Salt
	ciphertext := security.WithMd5(plaintext)
	ok := security.CheckPassword(ciphertext, user.Password)
	if !ok {
		ctx.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"message": "forbidden",
		})
		return
	}

	//
	accessToken, err := jwt.GenereateToken(user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	//
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"access_token":  accessToken,
			"refresh_token": "",
			"expired":       0,
		},
		"message": fmt.Sprint(payload.UserName, payload.Password),
	})
}
