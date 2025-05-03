package services

import "github.com/gin-gonic/gin"

type AuthService interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	CheckEmail(c *gin.Context)
	CheckUsername(c *gin.Context)
	Logout(c *gin.Context)
}
