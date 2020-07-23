package platform

import (
	"log"
	"net/http"

	"github.com/google/uuid"

	"gitee.com/taadis/letgo/store"
	"github.com/gin-gonic/gin"
)

// Add
func Add(ctx *gin.Context) {
	payload := struct {
		Name    string `form:"name" binding:"required"`
		Code    string `form:"code" binding:"required"`
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

	entity := &store.BasicPlatform{}
	entity.Id = uuid.New().String()
	entity.Name = payload.Name
	entity.Code = payload.Code
	entity.Enabled = payload.Enabled
	// TODO: check name/code value
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
		"message": "added platform",
	})
}
