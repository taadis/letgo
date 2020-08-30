package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserProfile
type UserProfile struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Alisa  string `json:"alisa"`
	Avatar string `json:"avatar"`
	Email  string `json:"email"`
}

// Profile
func Profile(ctx *gin.Context) {
	// get identity from context
	value, exists := ctx.Get("identity")
	if !exists {
		log.Println("identity not exists")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "identity not exists",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    value,
		"message": "profile",
	})
}
