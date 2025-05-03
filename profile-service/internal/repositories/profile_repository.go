package repositories

import (
	"profile-service/internal/domain"

	"github.com/gin-gonic/gin"
)

type ProfileRepository interface {
	GetProfileById(c *gin.Context, userId string) (*domain.Profile, error)
	UpdateProfile(c *gin.Context, profile *domain.Profile, userId string) (*domain.Profile, error)
	// CreateProfile(c *gin.Context, profile *domain.Profile) (*domain.Profile, error)
}
