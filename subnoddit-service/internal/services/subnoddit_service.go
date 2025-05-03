package services

import (
	"github.com/gin-gonic/gin"
	"subnoddit-service/internal/dtos/request"
)

type SubrodditService interface {
	CreateCommunity(ctx *gin.Context, req *request.CreateCommunityRequest)
	UpdateCommunity(ctx *gin.Context, req *request.UpdateCommunityRequest)
	GetCommunityById(ctx *gin.Context, communityId *string)
	ListCommunities(ctx *gin.Context)

	JoinCommunity(ctx *gin.Context, req *request.JoinCommunityRequest)
	LeaveCommunity(ctx *gin.Context, req *request.LeaveCommunityRequest)
	GetCommunityMemberCount(ctx *gin.Context, communityId *string)
	IsUserMember(ctx *gin.Context, req *request.IsUserMemberRequest)
}
