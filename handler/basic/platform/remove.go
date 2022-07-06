package platform

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taadis/letgo/store"
)

//
func Remove(ctx *gin.Context) {
	log.Println(ctx.Request.RequestURI)

	//
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
	err = store.Db.Where("code=?", payload.Code).Delete(&store.BasicPlatform{}).Error
	if err != nil {
		log.Println(".Delete error:", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  ".Delete error:" + err.Error(),
		})
		return
	}
	log.Println(".Delete success")

	//
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "removed",
	})
}
