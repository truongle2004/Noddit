package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
	"profile-service/internal/services"
)

type FollowController struct {
	followSvc services.FollowService
}

func NewFollowController(followSvc services.FollowService) *FollowController {
	return &FollowController{
		followSvc: followSvc,
	}
}

func (ctrl *FollowController) RegisterRoute(r *gin.Engine) {
	v1 := r.Group(core.V1 + "/profile-service/follows")
	{
		v1.POST("/create", ctrl.Follow)
		v1.POST("/delete", ctrl.UnFollow)
		v1.GET("/get", ctrl.GetFollowers)
	}
}

func (ctrl *FollowController) Follow(c *gin.Context) {
	ctrl.followSvc.Follow(c)
}

func (ctrl *FollowController) UnFollow(c *gin.Context) {
	ctrl.followSvc.UnFollow(c)
}

func (ctrl *FollowController) GetFollowers(c *gin.Context) {
	ctrl.followSvc.GetFollowers(c)
}
