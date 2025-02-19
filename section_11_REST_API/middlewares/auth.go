package middlewares

import (
	"net/http"
	"practice/Rest/utils"

	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}
	user_id, err := utils.VerifyToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}
	ctx.Set("user_id", user_id)
	ctx.Next()
}
