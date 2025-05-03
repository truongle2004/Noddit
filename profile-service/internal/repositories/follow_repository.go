package repositories

import (
	"github.com/gin-gonic/gin"
	"profile-service/internal/domain"
)

type FollowRepository interface {
	GetFollows(c *gin.Context, userId string) ([]domain.Profile, error)
	Follow(c *gin.Context, followerId, followeeId string) error
	UnFollow(c *gin.Context, followerId, followeeId string) error
}
