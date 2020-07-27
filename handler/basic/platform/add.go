package platform

import (
	"log"
	"net/http"

	"gitee.com/taadis/letgo/store"
	"github.com/gin-gonic/gin"
)

// Add
func Add(ctx *gin.Context) {
	payload := struct {
		Name    string `json:"name" binding:"required"`
		Code    string `json:"code" binding:"required"`
		Enabled bool   `json:"enabled" binding:"-"`
	}{}
	err := ctx.BindJSON(&payload)
	if err != nil {
		log.Println("ctx.Bind error", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	entity := &store.BasicPlatform{}
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
		"code":    0,
		"data":    entity,
		"message": "added platform",
	})
}
