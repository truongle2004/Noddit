package controller

import (
	"subnoddit-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
)

type CommunityController struct {
	communitySvc services.SubrodditService
}

func NewCommunityController(subNodditSvc services.SubrodditService) CommunityController {
	return CommunityController{communitySvc: subNodditSvc}
}

func (a *CommunityController) RegisterRoutes(r *gin.Engine) {
	api := r.Group(core.V1 + "/subnoddit-service")

	// Community routes
	community := api.Group("/communities")
	{
		community.POST("/create", a.communitySvc.CreateCommunity)
		community.PUT("/update", a.communitySvc.UpdateCommunity)
		community.GET("/:id", a.communitySvc.GetCommunityById)
		community.GET("", a.communitySvc.ListCommunities)
		community.POST("/join", a.communitySvc.JoinCommunity)
		community.POST("/leave", a.communitySvc.LeaveCommunity)
		community.GET("/:id/member-count", a.communitySvc.GetCommunityMemberCount)
		community.POST("/is-member", a.communitySvc.IsUserMember)
	}
}
