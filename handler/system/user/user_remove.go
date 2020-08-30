package user

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
			"code":    http.StatusBadRequest,
			"message": "Error .BindJson:" + err.Error(),
		})
		return
	}

	err = store.Db.Where("id = ?", payload.Id).Delete(&store.SystemUser{}).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error .Delete: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "删除成功",
	})
}
