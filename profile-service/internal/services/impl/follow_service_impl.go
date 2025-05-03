package impl

import (
	"net/http"
	"profile-service/internal/dto/request"
	"profile-service/internal/dto/response"
	"profile-service/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
)

type FollowServiceImpl struct {
	db repositories.FollowRepository
}

func NewFollowService(db repositories.FollowRepository) *FollowServiceImpl {
	return &FollowServiceImpl{
		db: db,
	}
}

func (f *FollowServiceImpl) Follow(c *gin.Context) {
	var followDto request.FollowRequestDto

	if err := c.ShouldBindJSON(&followDto); err != nil {
		c.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError(err.Error()).WithDetail("error", "Invalid request body"))
		return
	}
	if err := followDto.ValidateFollowRequest(); err != nil {

		c.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError(err.Error()))
		return
	}

	if err := f.db.Follow(c, followDto.FollowerId, followDto.FolloweeId); err != nil {
		c.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{})

}

func (f *FollowServiceImpl) UnFollow(c *gin.Context) {
	var followDto request.FollowRequestDto

	if err := c.ShouldBindJSON(&followDto); err != nil {
		c.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError(err.Error()).WithDetail("error", "Invalid request body"))
		return
	}

	if err := followDto.ValidateFollowRequest(); err != nil {
		c.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError(err.Error()))
		return
	}

	if err := f.db.UnFollow(c, followDto.FollowerId, followDto.FolloweeId); err != nil {
		c.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithError(err.Error()))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func (f *FollowServiceImpl) GetFollowers(c *gin.Context) {
	userId := c.GetHeader("X-Auth-User-Id")

	if userId == "" {
		c.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError("X-Auth-User-Id header is required"))
		return
	}

	listProfileFollower, err := f.db.GetFollows(c, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithError(err.Error()))
		return
	}

	if len(listProfileFollower) == 0 {
		c.JSON(http.StatusOK, []response.ProfileResponseDto{})
		return
	}

	var listProfileFollowerDto []response.ProfileResponseDto

	for _, profile := range listProfileFollower {
		profileDto := response.ProfileResponseDto{
			UserId:      profile.UserID,
			DisplayName: profile.DisplayName,
			AvatarUrl:   profile.AvatarURL,
		}
		listProfileFollowerDto = append(listProfileFollowerDto, profileDto)
	}

	c.JSON(http.StatusOK, listProfileFollowerDto)
}
