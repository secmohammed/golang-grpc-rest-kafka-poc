package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GuestUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)
		if len(fields) != 0 && fields[0] == "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "You are already logged in"})
			return
		}
		ctx.Next()
	}
}
