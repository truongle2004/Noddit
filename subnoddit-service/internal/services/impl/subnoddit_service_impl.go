package impl

import (
	"net/http"
	"subnoddit-service/internal/domain/models"
	"subnoddit-service/internal/dtos"
	"subnoddit-service/internal/dtos/request"
	"subnoddit-service/internal/repositories"
	"subnoddit-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
)

type CommunityServiceImpl struct {
	db       repositories.SubredditRepository
	imageSvc services.UploadImage
}

func NewCommunityService(db repositories.SubredditRepository, imageSvc services.UploadImage) *CommunityServiceImpl {
	return &CommunityServiceImpl{db: db, imageSvc: imageSvc}
}

func (s *CommunityServiceImpl) CreateCommunity(ctx *gin.Context) {
	// TODO: check creatorId
	//TODO: check the size of image file
	var req dtos.CommunityDto
	creatorID := ctx.GetHeader("X-Creator-ID")

	if creatorID == "" {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Creator ID is required"))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Invalid request body"))
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", err.Error()))
		return
	}

	var rules []models.Rule
	for _, rule := range req.Rules {
		rules = append(rules, models.Rule{
			CommunityID: &creatorID,
			Title:       &rule.Title,
			Description: &rule.Description,
			Position:    &rule.Position,
		})
	}
	community := &models.Community{
		Name:         &req.Name,
		Title:        &req.Title,
		Description:  &req.Description,
		Rules:        rules,
		Type:         &req.Type,
		BannerImage:  &req.BannerImage,
		ProfileImage: &req.BannerImage,
		CreatorID:    &creatorID,
	}

	if err := s.db.CreateCommunity(community); err != nil {
		ctx.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", "Failed to create community: "+err.Error()))
		return
	}

	communityDto := &dtos.CommunityDto{
		Name:         *community.Name,
		Title:        *community.Title,
		Description:  *community.Description,
		Rules:        req.Rules,
		Type:         *community.Type,
		BannerImage:  *community.BannerImage,
		ProfileImage: *community.ProfileImage,
		CreatorId:    *community.CreatorID,
		CreatedAt:    community.CreatedAt,
		UpdatedAt:    community.UpdatedAt,
	}

	ctx.JSON(http.StatusCreated, communityDto)
}

func (s *CommunityServiceImpl) UpdateCommunity(ctx *gin.Context) {
	var req request.UpdateCommunityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Invalid request body"))
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", err.Error()))
		return
	}

	community := &models.Community{}

	if err := s.db.UpdateCommunity(community); err != nil {
		ctx.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", "Failed to update community: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, community)
}

func (s *CommunityServiceImpl) GetCommunityById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Community ID is required"))
		return
	}
	community, err := s.db.GetCommunityByID(&id)
	if err != nil {
		ctx.Error(core.ErrInternalServerError.WithDetail("error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, community)
}

func (s *CommunityServiceImpl) ListCommunities(ctx *gin.Context) {
	communities, err := s.db.ListCommunities()
	if err != nil {
		ctx.Error(core.ErrInternalServerError.WithDetail("error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, communities)
}

func (s *CommunityServiceImpl) JoinCommunity(ctx *gin.Context) {
	var req request.JoinCommunityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Invalid request body"))
		return
	}
	if err := s.db.JoinCommunity(&req.UserID, &req.CommunityID); err != nil {
		ctx.Error(core.ErrInternalServerError.WithDetail("error", "Failed to join community: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Joined community successfully"})
}

func (s *CommunityServiceImpl) LeaveCommunity(ctx *gin.Context) {
	var req request.LeaveCommunityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Invalid request body"))
		return
	}
	if err := s.db.LeaveCommunity(&req.UserID, &req.CommunityID); err != nil {
		ctx.Error(core.ErrInternalServerError.WithDetail("error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Left community successfully"})
}

func (s *CommunityServiceImpl) GetCommunityMemberCount(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Community ID is required"))
		return
	}
	cnt, err := s.db.GetNumberOfMembersInCommunity(&id)
	if err != nil {
		ctx.Error(core.ErrInternalServerError.WithDetail("error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"member_count": cnt})
}

func (s *CommunityServiceImpl) IsUserMember(ctx *gin.Context) {

	var req request.IsUserMemberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Invalid request body"))
		return
	}
	isMember, err := s.db.IsUserMember(&req.UserID, &req.CommunityID)
	if err != nil {
		ctx.Error(core.ErrInternalServerError.WithDetail("error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"is_member": isMember,
	})
}
