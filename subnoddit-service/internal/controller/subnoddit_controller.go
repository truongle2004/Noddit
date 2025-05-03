package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
	"subnoddit-service/internal/dtos/request"
	"subnoddit-service/internal/services"
)

type CommunityController struct {
	subNodditSvc services.SubrodditService
}

func NewCommunityController(subNodditSvc services.SubrodditService) CommunityController {
	return CommunityController{subNodditSvc: subNodditSvc}
}

func (a *CommunityController) RegisterRoutes(r *gin.Engine) {
	v1 := r.Group(core.V1 + "/auth/community")
	{
		v1.POST("/create", a.CreateCommunity)
		v1.PUT("/update", a.UpdateCommunity)
		v1.GET("/:id", a.GetCommunityById)
		v1.GET("/", a.ListCommunities)
		v1.POST("/join", a.JoinCommunity)
		v1.POST("/leave", a.LeaveCommunity)
		v1.GET("/:id/member-count", a.GetCommunityMemberCount)
		v1.POST("/is-member", a.IsUserMember)
	}
}

func (a *CommunityController) CreateCommunity(ctx *gin.Context) {
	var req request.CreateCommunityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Invalid request body"))
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", err.Error()))
		return
	}
	a.subNodditSvc.CreateCommunity(ctx, &req)
}

func (a *CommunityController) UpdateCommunity(ctx *gin.Context) {
	var req request.UpdateCommunityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Invalid request body"))
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", err.Error()))
		return
	}
	a.subNodditSvc.UpdateCommunity(ctx, &req)
}

func (a *CommunityController) GetCommunityById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Community ID is required"))
		return
	}
	a.subNodditSvc.GetCommunityById(ctx, &id)
}

func (a *CommunityController) ListCommunities(ctx *gin.Context) {
	a.subNodditSvc.ListCommunities(ctx)
}

func (a *CommunityController) JoinCommunity(ctx *gin.Context) {
	var req request.JoinCommunityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Invalid request body"))
		return
	}
	a.subNodditSvc.JoinCommunity(ctx, &req)
}

func (a *CommunityController) LeaveCommunity(ctx *gin.Context) {
	var req request.LeaveCommunityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Invalid request body"))
		return
	}
	a.subNodditSvc.LeaveCommunity(ctx, &req)
}

func (a *CommunityController) GetCommunityMemberCount(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Community ID is required"))
		return
	}
	a.subNodditSvc.GetCommunityMemberCount(ctx, &id)
}

func (a *CommunityController) IsUserMember(ctx *gin.Context) {
	var req request.IsUserMemberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Invalid request body"))
		return
	}
	a.subNodditSvc.IsUserMember(ctx, &req)
}
