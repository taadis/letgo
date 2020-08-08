package shop

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
		Id      string `json:"id" binding:"-"`
		Name    string `json:"name" binding:"-"`
		Enabled string `json:"enabled" binding:"-"`
	}{}
	err := ctx.BindJSON(&payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  ".BindJson error:" + err.Error(),
		})
		return
	}

	var shops []store.BasicShop
	count := 0
	query := store.Db
	if payload.Id != "" {
		query = query.Where("id = ?", payload.Id)
	}
	if payload.Name != "" {
		query = query.Where("name = ?", payload.Name)
	}
	if payload.Enabled != "" {
		query = query.Where("enabled = ?", payload.Enabled)
	}
	err = query.Find(&shops).Count(&count).Error
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
	).Find(&shops).Error
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
		Id      string `json:"id"`
		Name    string `json:"name"`
		Enabled string `json:"enabled"`
	}{}
	for _, v := range shops {
		d := struct {
			Id      string `json:"id"`
			Name    string `json:"name"`
			Enabled string `json:"enabled"`
		}{}
		d.Id = v.Id
		d.Name = v.Name
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
