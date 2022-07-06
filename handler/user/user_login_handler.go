package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taadis/letgo/store"
	"github.com/taadis/letgo/util/jwt"
	"github.com/taadis/letgo/util/security"
)

// UserLogin
type UserLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login
func Login(ctx *gin.Context) {
	payload := &UserLogin{}
	err := ctx.Bind(payload)
	if err != nil {
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
		"code": 0,
		"data": gin.H{
			"access_token":  accessToken,
			"token_type":    "bearer",
			"refresh_token": "",
			"expired":       0,
			"scope":         "api",
		},
		"message": fmt.Sprint(payload.UserName, payload.Password),
	})
}
