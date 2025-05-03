package impl

import (
	"net/http"
	"profile-service/internal/dto/request"
	"profile-service/internal/mapper"
	"profile-service/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
)

type ProfileServiceImpl struct {
	db repositories.ProfileRepository
}

func NewProfileServiceImpl(db repositories.ProfileRepository) *ProfileServiceImpl {
	return &ProfileServiceImpl{
		db: db,
	}
}

func (p *ProfileServiceImpl) GetProfile(c *gin.Context) {
	userId := c.GetHeader("X-Auth-User-Id")

	profile, err := p.db.GetProfileById(c, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", err.Error()))
		return
	}

	if profile == nil {
		c.JSON(http.StatusNotFound, core.ErrNotFound.WithDetail("error", "Profile not found"))
		return
	}

	profileDto := mapper.ToProfileDto(profile)

	c.JSON(http.StatusOK, profileDto)
}

func (p *ProfileServiceImpl) UpdateProfile(c *gin.Context) {
	userId := c.GetHeader("X-Auth-User-Id")

	var profileDto request.ProfileDto

	if err := c.ShouldBindJSON(&profileDto); err != nil {
		c.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", err.Error()))
		return
	}

	if err := profileDto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError(err.Error()).WithDetail("error", err.Error()))
		return
	}

	profile := mapper.ToProfileEntity(&profileDto)

	if _, err := p.db.UpdateProfile(c, profile, userId); err != nil {
		c.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
	})
}
