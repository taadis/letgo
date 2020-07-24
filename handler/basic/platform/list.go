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
	store.Db.Where("").Find(&platforms)

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
