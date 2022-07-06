package cron

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/taadis/letgo/store"
)

// Create
func Create(ctx *gin.Context) {
	payload := struct {
		Name    string `form:"name" binding:"required"`
		Value   string `form:"value" binding:"required"`
		Enabled bool   `form:"enabled"`
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

	entity := &store.ScheduleCron{}
	entity.Id = uuid.New().String()
	entity.Name = payload.Name
	entity.Value = payload.Value
	entity.Enabled = payload.Enabled
	// TODO: check cron value
	err = store.Db.Create(entity).Error
	if err != nil {
		log.Println("db.Create error", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    entity,
		"message": "ok",
	})
}
