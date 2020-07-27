package platform

import (
	"log"
	"net/http"

	"gitee.com/taadis/letgo/common"
	"gitee.com/taadis/letgo/store"
	"github.com/gin-gonic/gin"
)

// List
func List(ctx *gin.Context) {
	payload := struct {
		common.Paging
		Code    string      `json:"code" binding:"-"`
		Name    string      `json:"name" binding:"-"`
		Enabled interface{} `json:"enabled" binding:"-"`
	}{}
	err := ctx.BindJSON(&payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  ".BindJson error:" + err.Error(),
		})
		return
	}

	var platforms []store.BasicPlatform
	count := 0
	query := store.Db
	if payload.Code != "" {
		log.Println("payload.Code != empty")
		query = query.Where("code = ?", payload.Code)
	}
	if payload.Name != "" {
		query = query.Where("name = ?", payload.Name)
	}
	if payload.Enabled != nil {
		query = query.Where("enabled = ?", payload.Enabled.(bool))
	}
	err = query.Find(&platforms).Count(&count).Error
	if err != nil {
		log.Print(".Count error:", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  ".Count error:" + err.Error(),
		})
		return
	}
	offset := (payload.Page - 1) * payload.Limit
	err = query.Limit(
		payload.Limit,
	).Offset(
		offset,
	).Find(&platforms).Error
	//err = store.Db.Where("").Find(&platforms).Error
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
		"count": count,
		"msg":   "ok",
	})
}
