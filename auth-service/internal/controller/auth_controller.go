package controller

import (
	"auth-service/internal/constant"
	"auth-service/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthSvc services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{AuthSvc: authService}
}

func (a *AuthController) RegisterRoutes(c *gin.Engine) {
	v1 := c.Group(constant.V1 + "/auth")
	{
		v1.POST("/register", a.AuthSvc.Register)
		v1.POST("/login", a.AuthSvc.Login)
		v1.POST("/logout", a.AuthSvc.Logout)
		v1.GET("/check-email/:email", a.AuthSvc.CheckEmail)
		v1.GET("/check-username/:username", a.AuthSvc.CheckUsername)
		v1.POST("/refresh", a.AuthSvc.RefreshToken)
	}
}
