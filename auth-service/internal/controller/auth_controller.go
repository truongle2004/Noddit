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
		v1.POST("/register", a.Register)
		v1.POST("/login", a.Login)
		v1.POST("/logout", a.Logout)
		v1.GET("/check-email/:email", a.CheckEmail)
		v1.GET("/check-username/:username", a.CheckUsername)
	}
}

func (a *AuthController) Register(ctx *gin.Context) {
	a.AuthSvc.Register(ctx)
}

func (a *AuthController) CheckEmail(ctx *gin.Context) {
	a.AuthSvc.CheckEmail(ctx)
}

func (a *AuthController) Login(ctx *gin.Context) {
	a.AuthSvc.Login(ctx)
}

func (a *AuthController) Logout(ctx *gin.Context) {
	a.AuthSvc.Logout(ctx)
}

func (a *AuthController) CheckUsername(ctx *gin.Context) {
	a.AuthSvc.CheckUsername(ctx)
}
