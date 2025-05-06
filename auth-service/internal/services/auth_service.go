package services

import (
	"auth-service/internal/dto/request"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Login(c *gin.Context, loginDto *request.LoginDto)
	Register(c *gin.Context, registerDto *request.RegisterDto)
	CheckEmail(c *gin.Context)
	CheckUsername(c *gin.Context)
	Logout(c *gin.Context)
	RefreshToken(c *gin.Context)
}
