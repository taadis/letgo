package platform

import (
	"log"
	"net/http"

	"gitee.com/taadis/letgo/store"
	"github.com/gin-gonic/gin"
)

// List
func List(ctx *gin.Context) {
	log.Println(ctx.Request.RequestURI)
	var platforms []store.BasicPlatform
	err := store.Db.Where("").Find(&platforms).Error
	if err != nil {
		log.Print(".Find error:", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  ".Find error:" + err.Error(),
		})
		return
	}

	//
	data := []struct {
		Code    string `json:"code"`
		Name    string `json:"name"`
		Enabled bool   `json:"enabled"`
	}{}
	for i, v := range platforms {
		log.Println(i, v)
		d := struct {
			Code    string `json:"code"`
			Name    string `json:"name"`
			Enabled bool   `json:"enabled"`
		}{}
		d.Code = v.Code
		d.Name = v.Name
		log.Println("v.Enabled", v.Enabled)
		d.Enabled = v.Enabled
		data = append(data, d)
	}

	//
	ctx.JSON(http.StatusOK, gin.H{
		"code":  0,
		"data":  data,
		"count": len(platforms),
		"msg":   "ok",
	})
}
