package users

import "github.com/gin-gonic/gin"

type UserRestHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}
