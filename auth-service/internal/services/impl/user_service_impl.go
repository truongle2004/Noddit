package impl

import (
	"auth-service/internal/constant"
	dto "auth-service/internal/dto/request"
	"auth-service/internal/dto/response"
	"auth-service/internal/helper"
	"auth-service/internal/repositories"
	"auth-service/internal/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	db    repositories.UserRepository
	redis services.RedisService
}

func NewUserService(userRepo repositories.UserRepository, redis services.RedisService) *UserServiceImpl {
	return &UserServiceImpl{
		db:    userRepo,
		redis: redis,
	}
}

//func (u *UserServiceImpl) GetUserById(c *gin.Context) {
//	id := c.GetHeader(constant.XAuthUserID)
//
//	if id == "" {
//		c.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError("Id is required"))
//		return
//	}
//
//	user, err := u.db.GetUserByID(c.Request.Context(), id)
//	if err != nil {
//		// TODO: do something
//	}
//
//	rolesDto := []response.RoleDto{}
//
//	userDto := response.UserDto{
//		Id:       user.ID,
//		Email:    user.Email,
//		Username: user.Username,
//		Role:     rolesDto,
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"data": userDto,
//	})
//}

func (u *UserServiceImpl) DeleteUserById(c *gin.Context) {
	id := c.GetHeader(constant.XAuthUserID)

	if id == "" {
		c.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError("Id is required"))
		return
	}

	err := u.db.DeleteUser(c.Request.Context(), id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, core.ErrNotFound.
			WithError("User you want to delete does not exist").
			WithReason("User not found").
			WithDetail("error", err.Error()))
		return
	} else if err != nil {
		helper.ResponseServerError(c, "Failed to delete user", err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "Delete user successfully",
	})
}

func (u *UserServiceImpl) GetAllUser(c *gin.Context) {
	users, err := u.db.GetAllUser(c.Request.Context())
	if err.Error() == core.ErrRecordNotFound.Error() || err != nil {
		c.JSON(http.StatusNotFound, core.ErrNotFound.
			WithError("Failed to get all user").
			WithReason(err.Error()).
			WithDetail("error", "get_all_user_failed"))
		return
	}

	userResponseDto := []response.UserDto{}
	for _, user := range users {

		rolesDto := []response.RoleDto{}

		userResponseDto = append(userResponseDto, response.UserDto{
			Id:       user.ID,
			Email:    user.Email,
			Username: user.Username,
			Role:     rolesDto,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userResponseDto,
	})
}

func (u *UserServiceImpl) CreateUser(c *gin.Context) {
	// TODO: create user without verification email
}

func (u *UserServiceImpl) BlockUser(c *gin.Context) {
	id := c.GetHeader(constant.XAuthUserID)

	if id == "" {
		c.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError("Id is required"))
		return
	}
	if err := u.db.SetUserAccountStatus(c.Request.Context(), id, string(core.LOCKED)); err != nil {
		if err.Error() == core.ErrRecordNotFound.Error() {
			c.JSON(http.StatusNotFound, core.ErrNotFound.
				WithError("User account not found").
				WithReason("Set user account status failed").
				WithDetail("error", err.Error()))
			return
		}
		helper.ResponseServerError(c, "Failed to block user", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User blocked successfully",
	})
}

func (u *UserServiceImpl) UnBlockUser(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError("Id is required"))
		return
	}
	if err := u.db.SetUserAccountStatus(c.Request.Context(), id, string(core.UNLOCK)); err != nil {
		helper.ResponseServerError(c, "Failed to unblock user", err)
		return
	}
}

func (u *UserServiceImpl) UpdateAccountStatus(c *gin.Context) {
	var updateAccountRequest dto.AccountStatusUpdateRequest

	if err := c.ShouldBindJSON(&updateAccountRequest); err != nil {
		c.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError("Invalid request body"))
		return
	}

	if err := updateAccountRequest.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError(err.Error()))
		return
	}

	if err := u.db.SetUserAccountStatus(c.Request.Context(), updateAccountRequest.ID, updateAccountRequest.Status); err != nil {
		helper.ResponseServerError(c, "Failed to update user account status", err)
		return
	}

	// Update status in redis
	if err := u.redis.Set(c, updateAccountRequest.Email, updateAccountRequest.Status, 0); err != nil {
		helper.ResponseServerError(c, "Failed to update user account status in redis", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User account status updated successfully",
	})
}

func (u *UserServiceImpl) GetAllAccountStatus(c *gin.Context) {

}
