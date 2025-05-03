package services

import (
	"github.com/gin-gonic/gin"
)

type UserService interface {
	DeleteUserById(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateAccountStatus(c *gin.Context)
	GetAllAccountStatus(c *gin.Context)
}
