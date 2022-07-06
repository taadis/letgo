package platform

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taadis/letgo/store"
)

//
func Detail(ctx *gin.Context) {
	payload := struct {
		Code string `json:"code" binding:"required"`
	}{}
	err := ctx.BindJSON(&payload)
	if err != nil {
		log.Println(".BindJson error:", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  ".BindJson error:" + err.Error(),
		})
		return
	}

	//
	var platform store.BasicPlatform
	err = store.Db.Where("code=?", payload.Code).Find(&platform).Error
	if err != nil {
		log.Println(".Find error:", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  ".Find error:" + err.Error(),
		})
		return
	}
	log.Println(".Find success")

	//
	data := struct {
		Code    string `json:"code"`
		Name    string `json:"name"`
		Enabled bool   `json:"enabled"`
	}{}
	data.Code = platform.Code
	data.Name = platform.Name
	data.Enabled = platform.Enabled
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
		"msg":  "removed",
	})
}
