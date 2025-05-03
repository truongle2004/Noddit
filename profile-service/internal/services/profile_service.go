package services

import (
	"github.com/gin-gonic/gin"
)

type ProfileService interface {
	GetProfile(c *gin.Context)

	// UpdateProfile updates the profile if it exists, otherwise create a new profile
	UpdateProfile(c *gin.Context)
}
