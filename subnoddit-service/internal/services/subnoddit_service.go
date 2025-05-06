package services

import (
	"subnoddit-service/internal/dtos"
	"subnoddit-service/internal/dtos/request"

	"github.com/gin-gonic/gin"
)

type SubrodditService interface {
	UploadCommunityImage(ctx *gin.Context)
	CreateCommunity(ctx *gin.Context, req *dtos.CommunityDto, createId *string)
	UpdateCommunity(ctx *gin.Context, req *request.UpdateCommunityRequest)
	GetCommunityById(ctx *gin.Context, communityId *string)
	ListCommunities(ctx *gin.Context)

	JoinCommunity(ctx *gin.Context, req *request.JoinCommunityRequest)
	LeaveCommunity(ctx *gin.Context, req *request.LeaveCommunityRequest)
	GetCommunityMemberCount(ctx *gin.Context, communityId *string)
	IsUserMember(ctx *gin.Context, req *request.IsUserMemberRequest)
}
