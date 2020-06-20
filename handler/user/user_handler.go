package user

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"

	"gitee.com/taadis/letgo/store"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserRegister
type UserRegister struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
	//Email    string `binding:"email"`
}

func Register(ctx *gin.Context) {
	payload := &UserRegister{}
	err := ctx.Bind(payload)
	if err != nil {
		log.Println("ctx.Bind error", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"mesaage": fmt.Sprint("bad request ", err.Error()),
		})
		return
	}

	//
	systemUser := &store.SystemUser{}
	systemUser.Id = uuid.New().String()
	systemUser.Name = payload.Name
	systemUser.Salt = uuid.New().String()
	// password is md5(password + salt)
	systemUser.Password = fmt.Sprintf("%x", md5.Sum([]byte(payload.Password+systemUser.Salt)))
	err = store.Db.Create(systemUser).Error
	if err != nil {
		log.Println("Db.Create error", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": fmt.Sprint(err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "register success",
	})

	/*
		//如果名称为空给一个随机字符串
		if len(name) == 0 {
			name = util.RandomString(10)
		}
		if isTelephoneExist(DB, telephone) {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
			return
		}
		hasePassowrd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
			return
		}
		newUser := model.User{
			Name:      name,
			Telephone: telephone,
			Password:  string(hasePassowrd),
		}
		DB.Create(&newUser)
		//发送token
		token, err := common.ReleaseToken(newUser)
		if err != nil {
			response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
			log.Printf("token generate error:%v", err)
			return
		}

		//返回结果
		response.Success(ctx, gin.H{"token": token}, "注册成功")
	*/
}
