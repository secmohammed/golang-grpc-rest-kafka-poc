package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "github.com/secmohammed/golang-kafka-grpc-poc/app/usecase/user"
    "github.com/secmohammed/golang-kafka-grpc-poc/config"
    "github.com/secmohammed/golang-kafka-grpc-poc/utils"
    "net/http"
    "strings"
)

func AuthUser(uc user.UseCase, c config.Repository) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var token string

        authorizationHeader := ctx.Request.Header.Get("Authorization")
        fields := strings.Fields(authorizationHeader)

        if len(fields) != 0 && fields[0] == "Bearer" {
            token = fields[1]
        }

        if token == "" {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
            return
        }
        secret, err := c.GetString("app.jwt.secret")
        if err != nil {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
            return
        }
        sub, err := utils.ValidateToken(token, secret)
        if err != nil {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
            return
        }
        user, err := uc.GetUserByID(uuid.MustParse(sub.(string)))
        if err != nil {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
            return
        }
        ctx.Set("user", user)
        ctx.Next()
    }
}
