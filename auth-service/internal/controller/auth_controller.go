package controller

import (
	"auth-service/internal/constant"
	"auth-service/internal/dto/request"
	"auth-service/internal/helper"
	"auth-service/internal/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
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
		v1.POST("/refresh", a.RefreshToken)
	}
}

func (a *AuthController) Register(ctx *gin.Context) {

	var registerDto request.RegisterDto

	if err := ctx.ShouldBindJSON(&registerDto); err != nil {
		helper.ResponseServerError(ctx, "error bind json: "+err.Error(), errors.New("register_failed"))
		return
	}

	if err := registerDto.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError(err.Error()))
		return
	}

	a.AuthSvc.Register(ctx, &registerDto)
}

func (a *AuthController) CheckEmail(ctx *gin.Context) {
	a.AuthSvc.CheckEmail(ctx)
}

func (a *AuthController) Login(ctx *gin.Context) {
	var loginDto request.LoginDto

	if err := ctx.ShouldBindJSON(&loginDto); err != nil {
		helper.ResponseServerError(ctx, "Bin json data failed", err)
		return
	}

	if err := loginDto.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError(err.Error()))
		return
	}

	a.AuthSvc.Login(ctx, &loginDto)
}

func (a *AuthController) Logout(ctx *gin.Context) {
	a.AuthSvc.Logout(ctx)
}

func (a *AuthController) CheckUsername(ctx *gin.Context) {
	a.AuthSvc.CheckUsername(ctx)
}

func (a *AuthController) RefreshToken(ctx *gin.Context) {
	a.AuthSvc.RefreshToken(ctx)
}