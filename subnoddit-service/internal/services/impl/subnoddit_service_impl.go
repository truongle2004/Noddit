package impl

import (
	"net/http"
	"subnoddit-service/internal/domain/models"
	"subnoddit-service/internal/dtos/request"
	"subnoddit-service/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
)

type CommunityServiceImpl struct {
	db repositories.SubredditRepository
}

func NewCommunityService(db repositories.SubredditRepository) *CommunityServiceImpl {
	return &CommunityServiceImpl{db: db}
}

func (s *CommunityServiceImpl) CreateCommunity(ctx *gin.Context, req *request.CreateCommunityRequest) {

	community := &models.Community{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.db.CreateCommunity(community); err != nil {
		ctx.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", "Failed to create community: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, community)
}

func (s *CommunityServiceImpl) UpdateCommunity(ctx *gin.Context, req *request.UpdateCommunityRequest) {
	community := &models.Community{}

	community.ID = req.ID
	community.Name = *req.Name
	community.Description = *req.Description

	if err := s.db.UpdateCommunity(community); err != nil {
		ctx.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", "Failed to update community: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, community)
}

func (s *CommunityServiceImpl) GetCommunityById(ctx *gin.Context, communityId *string) {
	community, err := s.db.GetCommunityByID(communityId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", "Failed to get community: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, community)
}

func (s *CommunityServiceImpl) ListCommunities(ctx *gin.Context) {
	communities, err := s.db.ListCommunities()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", "Failed to list communities: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, communities)
}

func (s *CommunityServiceImpl) JoinCommunity(ctx *gin.Context, req *request.JoinCommunityRequest) {

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", err.Error()))
		return
	}

	if err := s.db.JoinCommunity(&req.UserID, &req.CommunityID); err != nil {
		ctx.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", "Failed to join community: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Joined community successfully"})
}

func (s *CommunityServiceImpl) LeaveCommunity(ctx *gin.Context, req *request.LeaveCommunityRequest) {
	if err := s.db.LeaveCommunity(&req.UserID, &req.CommunityID); err != nil {
		ctx.JSON(http.StatusInternalServerError,
			core.
				ErrInternalServerError.
				WithDetail("error", "Failed to leave community: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Left community successfully"})
}

func (s *CommunityServiceImpl) GetCommunityMemberCount(ctx *gin.Context, communityId *string) {
	cnt, err := s.db.GetNumberOfMembersInCommunity(communityId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"member_count": cnt})
}

func (s *CommunityServiceImpl) IsUserMember(ctx *gin.Context, req *request.IsUserMemberRequest) {

	isMember, err := s.db.IsUserMember(&req.UserID, &req.CommunityID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, core.ErrInternalServerError.
			WithDetail("error", "Failed to check user member: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"is_member": isMember})
}
