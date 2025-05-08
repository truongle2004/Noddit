package controller

import (
	"subnoddit-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
)

type CommunityController struct {
	subNodditSvc services.SubrodditService
	imageSvc     services.UploadImage
}

func NewCommunityController(subNodditSvc services.SubrodditService, imageSvc services.UploadImage) CommunityController {
	return CommunityController{subNodditSvc: subNodditSvc, imageSvc: imageSvc}
}

func (a *CommunityController) RegisterRoutes(r *gin.Engine) {
	api := r.Group(core.V1 + "/subnoddit-service")

	// Community routes
	community := api.Group("/community")
	{
		community.POST("/create", a.subNodditSvc.CreateCommunity)
		community.PUT("/update", a.subNodditSvc.UpdateCommunity)
		community.GET("/:id", a.subNodditSvc.GetCommunityById)
		community.GET("", a.subNodditSvc.ListCommunities)
		community.POST("/join", a.subNodditSvc.JoinCommunity)
		community.POST("/leave", a.subNodditSvc.LeaveCommunity)
		community.GET("/:id/member-count", a.subNodditSvc.GetCommunityMemberCount)
		community.POST("/is-member", a.subNodditSvc.IsUserMember)
	}

	// Image upload & serving
	image := api.Group("/image")
	{
		image.POST("/upload", a.imageSvc.UploadImage)
		image.GET("/:filename", a.imageSvc.LoadImage)
	}
}
