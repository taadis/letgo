package security

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taadis/letgo/util/security"
)

func EncryptMD5(ctx *gin.Context) {
	payload := struct {
		Plaintext string `json:"plaintext" binding:"required"`
	}{}
	err := ctx.BindJSON(&payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ciphertext := security.WithMd5(payload.Plaintext)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"plaintext":  payload.Plaintext,
			"ciphertext": ciphertext,
		},
		"message": "ok",
	})
}
