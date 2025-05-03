package impl

import (
	"errors"
	"log"
	"net/http"
	"time"

	"auth-service/internal/constant"
	domain "auth-service/internal/domain/models"
	requestDto "auth-service/internal/dto/request"
	"auth-service/internal/dto/response"
	"auth-service/internal/helper"
	"auth-service/internal/repositories"
	"auth-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/truongle2004/service-context/core"
	"gorm.io/gorm"
)

const (
	RefreshTokenKey = "refresh_token"
)

type AuthServiceImpl struct {
	userRepo  repositories.UserRepository
	redisSvc  services.RedisService
	casbinSvc services.CasbinService
}

func NewAuthService(userRepo repositories.UserRepository,
	redisSvc services.RedisService,
	casbinSvc services.CasbinService,
) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepo:  userRepo,
		redisSvc:  redisSvc,
		casbinSvc: casbinSvc,
	}
}

func (u *AuthServiceImpl) Login(c *gin.Context) {
	var loginDto requestDto.LoginDto

	if err := c.ShouldBindJSON(&loginDto); err != nil {
		helper.ResponseServerError(c, "Bin json data failed", err)
		return
	}

	if err := loginDto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError(err.Error()))
		return
	}

	val, err := u.redisSvc.Get(c, loginDto.Email)

	if errors.Is(err, redis.Nil) {
		c.JSON(http.StatusUnauthorized, core.ErrUnauthorized.
			WithError(constant.MessageInvalidCredential).
			WithReason("Invalid credential provided").
			WithDetail("error", "failed_login"))
		return
	} else if val != string(core.ACTIVE) {
		c.JSON(http.StatusForbidden, core.ErrForbidden.
			WithError("Your account is locked. You should contact admin or support").
			WithReason("The account is locked").
			WithDetail("error", "failed_login"))
		return
	}

	user, err := u.userRepo.GetUserByEmail(c.Request.Context(), loginDto.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusUnauthorized, core.ErrUnauthorized.
			WithError(constant.MessageInvalidCredential).
			WithReason("Invalid credential provided").
			WithDetail("error", "failed_login"))
		return
	}

	if !helper.CompareHashPassword(user.Password, user.Salt, loginDto.Password) {
		c.JSON(http.StatusUnauthorized, core.ErrUnauthorized.
			WithReason("Compare hash password failed").
			WithError(constant.MessageInvalidCredential).
			WithDetail("error", "failed_compare_password"))
		return
	}

	if err := u.userRepo.UpdateLastLogin(c.Request.Context(), user.ID); err != nil {
		log.Printf("Failed to update last login: %v", err)
	}

	accessToken, refreshToken, err := helper.GenerateAccessTokenAndRefreshToken(user.ID, user.Email)
	if err != nil {
		helper.ResponseServerError(c, "Generate token failed", err)
		return
	}

	platform := c.GetHeader("X-Client-Platform")

	switch platform {
	case "web":
		c.SetCookie(RefreshTokenKey, refreshToken, 3600*24*30, "/", "localhost", true, true)
		tokenResponseDto := response.TokenResponseDto{
			Token: response.TokenDto{
				AccessToken: accessToken,
				TokenType:   "Bearer",
			},
		}

		c.JSON(http.StatusOK, tokenResponseDto)
		return
	case "mobile":
		tokenResponseDto := response.TokenResponseDto{
			Token: response.TokenDto{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
				TokenType:    "Bearer",
			},
		}
		c.JSON(http.StatusOK, tokenResponseDto)
		return
	default:
		c.JSON(http.StatusBadRequest, core.ErrBadRequest.
			WithError("Invalid platform").
			WithReason("Missing platform header").
			WithDetail("error", "invalid_platform"))
	}
}

func (u *AuthServiceImpl) Register(c *gin.Context) {
	var registerDto requestDto.RegisterDto

	if err := c.ShouldBindJSON(&registerDto); err != nil {
		helper.ResponseServerError(c, "error bind json: "+err.Error(), errors.New("register_failed"))
		return
	}

	if _, err := u.redisSvc.Get(c, registerDto.Email); !errors.Is(err, redis.Nil) {
		c.JSON(http.StatusConflict, core.ErrConflict.
			WithError("Email is already taken"))
		return
	}

	if _, err := u.redisSvc.Get(c, registerDto.Username); !errors.Is(err, redis.Nil) {
		c.JSON(http.StatusConflict, core.ErrConflict.
			WithError("Username is already taken"))
		return
	}

	if err := u.redisSvc.Set(c, registerDto.Email, string(core.ACTIVE), 0); err != nil {
		helper.ResponseServerError(c, "Set email to redis failed", err)
		return
	}

	if err := u.redisSvc.Set(c, registerDto.Username, registerDto.Username, 0); err != nil {
		helper.ResponseServerError(c, "Set username to redis failed", err)
		return
	}

	if err := registerDto.Validate(); err != nil {
		c.JSON(http.StatusUnauthorized, core.ErrUnauthorized.
			WithError(err.Error()))
		return
	}

	salt, err := helper.RandomStr(constant.SaltLimit)
	if err != nil {
		helper.ResponseServerError(c, "Failed to get salt", err)
		return
	}

	hashedPassword, err := helper.HashPassword(salt, registerDto.Password)
	if err != nil {
		helper.ResponseServerError(c, "Failed to hash password", err)
		return
	}

	user := domain.User{
		Username:  registerDto.Username,
		Email:     registerDto.Email,
		Password:  hashedPassword,
		Status:    core.ACTIVE,
		LastLogin: time.Now(),
		Salt:      salt,
	}

	if err := u.userRepo.Create(c.Request.Context(), &user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusConflict, core.ErrConflict.
				WithError("Username or email already exists. Please try another email or username").
				WithReason("Failed to create new user").
				WithDetail("error", err.Error()))
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully. Please login to continue",
	})
}

func (u *AuthServiceImpl) Logout(c *gin.Context) {
	c.SetCookie(RefreshTokenKey, "", -1, "/", "localhost", true, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout success",
	})
}

func (u *AuthServiceImpl) CheckUsername(c *gin.Context) {
	username := c.Param("username")

	if username == "" {
		c.JSON(http.StatusBadRequest, core.ErrBadRequest.
			WithError("Username is required"))
		return
	}

	if _, err := u.redisSvc.Get(c, username); !errors.Is(err, redis.Nil) {
		c.JSON(http.StatusConflict, core.ErrConflict.
			WithError("Username is already exists"))
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}

func (u *AuthServiceImpl) CheckEmail(c *gin.Context) {
	email := c.Param("email")

	if email == "" {
		c.JSON(http.StatusBadRequest, core.ErrBadRequest.
			WithError("Email is required").
			WithDetail("error", "Email is required"))
		return
	}

	if _, err := u.redisSvc.Get(c, email); !errors.Is(err, redis.Nil) {
		c.JSON(http.StatusConflict, core.ErrConflict.
			WithError("Email is already exists"))
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}
