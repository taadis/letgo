package cron

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	payload := struct {
		UserName string `form:"username" binding:"required"`
	}{}
	err := ctx.Bind(&payload)
	if err != nil {
		log.Println("ctx.Bind error", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    payload,
		"message": "ok",
	})
}
