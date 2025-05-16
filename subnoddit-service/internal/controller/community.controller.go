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

	community := api.Group("/communities")
	{
		community.POST("", a.communitySvc.CreateCommunity)
		community.PUT("/", a.communitySvc.UpdateCommunity)
		community.GET("", a.communitySvc.ListCommunities)
		community.GET("/:id", a.communitySvc.GetCommunityById)
		community.GET("/:id/topics", a.communitySvc.GetAllCommunityByTopic)
		community.POST("/:id/members", a.communitySvc.JoinCommunity)
		community.DELETE("/:id/members", a.communitySvc.LeaveCommunity)
		community.GET("/:id/members/me", a.communitySvc.IsUserMember)
		community.GET("/:id/members/count", a.communitySvc.GetCommunityMemberCount)
	}
}
