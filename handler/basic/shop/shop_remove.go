package shop

import (
	"net/http"

	"gitee.com/taadis/letgo/store"
	"github.com/gin-gonic/gin"
)

//
func Remove(ctx *gin.Context) {
	payload := struct {
		Id string `json:"id" binding:"required"`
	}{}
	err := ctx.BindJSON(&payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  ".BindJson error:" + err.Error(),
		})
		return
	}

	//
	err = store.Db.Where("id = ?", payload.Id).Delete(&store.BasicShop{}).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  ".Delete error:" + err.Error(),
		})
		return
	}

	//
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "removed",
	})
}
