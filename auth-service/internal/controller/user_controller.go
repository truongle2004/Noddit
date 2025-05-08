package controller

import (
	"github.com/gin-gonic/gin"

	"auth-service/internal/constant"
	"auth-service/internal/services"
)

type UserController struct {
	userSvc services.UserService
}

func NewUserController(userSvc services.UserService) *UserController {
	return &UserController{
		userSvc: userSvc,
	}
}

func (ctrl *UserController) RegisterRoutes(r *gin.Engine) {
	v1 := r.Group(constant.V1 + "/users")
	{
		// v1.GET("/email/:email", ctrl.GetUserByEmail)
		v1.DELETE("/", ctrl.userSvc.DeleteUserById)
	}
}
