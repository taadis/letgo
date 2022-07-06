package auth

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/taadis/letgo/store"
	"github.com/taadis/letgo/util/jwt"
)

// AuthMiddleware
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Authorization header
		tokenString := ctx.GetHeader("Authorization")
		// validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "unauthorized"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := jwt.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "unauthorized"})
			ctx.Abort()
			return
		}

		// userId
		userId := claims.UserId
		log.Println("userId:", userId)
		user := &store.SystemUser{}
		err = store.Db.Where("id=?", userId).First(user).Error
		if err != nil {
			log.Println("db.First error", err.Error())
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "unauthorized"})
			ctx.Abort()
			return
		}

		// set identity to http context
		ctx.Set("identity", user)
		ctx.Next()
	}
}
