package services

import (
	"github.com/gin-gonic/gin"
)

type SubrodditService interface {
	CreateCommunity(ctx *gin.Context)
	UpdateCommunity(ctx *gin.Context)
	GetCommunityById(ctx *gin.Context)
	ListCommunities(ctx *gin.Context)

	JoinCommunity(ctx *gin.Context)
	LeaveCommunity(ctx *gin.Context)
	GetCommunityMemberCount(ctx *gin.Context)
	IsUserMember(ctx *gin.Context)
}
