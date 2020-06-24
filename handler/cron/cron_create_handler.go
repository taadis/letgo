package cron

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	var payload struct {
		a string
		b string
		c string
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    payload,
		"message": "ok",
	})
}
