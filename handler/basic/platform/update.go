package platform

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taadis/letgo/store"
)

//
func Update(ctx *gin.Context) {
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

	// TODO: check name/code value
	err = store.Db.Model(
		&store.BasicPlatform{},
	).Where(
		"code=?",
		payload.Code,
	).Update(
		"name",
		payload.Name,
	).Update(
		"enabled",
		payload.Enabled,
	).Error

	if err != nil {
		log.Println("db.Update error:", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "db.Update error:" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"data":    "",
		"message": "updated",
	})
}
