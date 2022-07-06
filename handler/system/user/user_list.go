package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taadis/letgo/common"
	"github.com/taadis/letgo/store"
)

//
func changePassword(id, oldPassword, newPassword string) error {
	return errors.New("修改密码开发中")
}

// ChangePassword
func ChangePassword(ctx *gin.Context) {
	payload := struct {
		Id          string `json:"id" binding:"required"`
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}{}
	err := ctx.BindJSON(&payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Error .BindJson: " + err.Error(),
		})
		return
	}

	err = changePassword(payload.Id, payload.OldPassword, payload.NewPassword)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Error changePassword: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "密码修改成功",
	})
}

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

	var users []store.SystemUser
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
	err = query.Find(&users).Count(&count).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  ".Count error:" + err.Error(),
		})
		return
	}
	offset := (payload.Page - 1) * payload.Limit
	err = query.Limit(payload.Limit).Offset(offset).Find(&users).Error
	if err != nil {
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
	for _, v := range users {
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
