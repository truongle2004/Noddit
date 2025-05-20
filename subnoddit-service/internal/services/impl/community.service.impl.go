package impl

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"path"
	"path/filepath"
	"subnoddit-service/internal/constant"
	"subnoddit-service/internal/domain/models"
	"subnoddit-service/internal/dtos"
	"subnoddit-service/internal/environment"
	"subnoddit-service/internal/helper"
	"subnoddit-service/internal/repositories"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/truongle2004/service-context/core"
	"github.com/truongle2004/service-context/utils"
)

type CommunityServiceImpl struct {
	communityDB repositories.CommunityRepository
	topicDB     repositories.TopicRepository
}

func NewCommunityService(communityDB repositories.CommunityRepository, topicDB repositories.TopicRepository) *CommunityServiceImpl {
	return &CommunityServiceImpl{communityDB: communityDB, topicDB: topicDB}
}

func (s *CommunityServiceImpl) CreateCommunity(ctx *gin.Context) {
	// Extract basic form data
	name := ctx.PostForm("name")
	description := ctx.PostForm("description")
	_type := ctx.PostForm("type")
	creatorId := ctx.PostForm("creator_id")
	topics := ctx.PostForm("topics")

	bannerFile, bannerImageErr := ctx.FormFile("banner_image")
	iconFile, iconImageErr := ctx.FormFile("icon_image")

	newCommunityID := uuid.New().String()

	// Processing topic string
	var topicIds []string
	if topics != "" {
		topicIds = helper.SplitTopicIDs(topics)
	}

	// FIXME: query so many times
	var topicsModel []*models.Topic
	for _, id := range topicIds {
		val, err := s.topicDB.GetTopicByID(&id)
		if err != nil {
			log.Println(fmt.Sprintf("topic with id %s not found", id))
			continue
		}

		topicsModel = append(topicsModel, val)
	}

	community := &models.Community{
		SQLModel: core.SQLModel{
			ID: newCommunityID,
		},
		Name:        &name,
		Description: &description,
		Type:        &_type,
		CreatorID:   &creatorId,
		Topics:      topicsModel,
	}

	var bannerFileName string
	var iconFileName string

	if bannerImageErr == nil {
		ext, err := helper.CheckExtension(filepath.Ext(bannerFile.Filename))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError(err.Error()))
		}
		bannerFileName = fmt.Sprintf("%s_%s%s", "banner", newCommunityID, ext)
		community.BannerImage = &bannerFileName
	}

	if iconImageErr == nil {
		ext, err := helper.CheckExtension(filepath.Ext(bannerFile.Filename))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError(err.Error()))
		}
		iconFileName = fmt.Sprintf("%s_%s%s", "icon", newCommunityID, ext)
		community.IconImage = &iconFileName

	}
	if err := s.communityDB.CreateCommunity(community); err != nil {
		ctx.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", err.Error()))
		return
	}

	// Start processing image
	var wg sync.WaitGroup

	if bannerImageErr == nil && bannerFile != nil && bannerFileName != "" {
		wg.Add(1)
		go func(file *multipart.FileHeader, path string) {
			defer wg.Done() // ensure Done is always called
			dst := utils.GenerateDst(constant.UploadImagePath, bannerFileName)
			if err := utils.SaveImageToFolder(ctx, dst, file); err != nil {
				log.Println("Failed to save banner:", err)
			}
		}(bannerFile, bannerFileName)
	}

	if iconImageErr == nil && iconFile != nil && iconFileName != "" {
		wg.Add(1)
		go func(file *multipart.FileHeader, path string) {
			defer wg.Done()
			dst := utils.GenerateDst(constant.UploadImagePath, iconFileName)
			if err := utils.SaveImageToFolder(ctx, dst, file); err != nil {
				log.Println("Failed to save icon:", err)
			}
		}(iconFile, iconFileName)
	}

	wg.Wait() // Wait for all goroutines to complete

	var topicDtos []dtos.TopicDto
	for _, topic := range topicsModel {
		topicDto := dtos.TopicDto{
			ID:   topic.ID,
			Name: topic.Name,
		}
		topicDtos = append(topicDtos, topicDto)
	}

	communityDto := dtos.CommunityDto{
		ID:          community.ID,
		Name:        community.Name,
		Description: community.Description,
		Type:        community.Type,
		CreatorId:   community.CreatorID,
		BannerImage: &bannerFileName,
		IconImage:   &iconFileName,
		Topics:      topicDtos,
	}

	ctx.JSON(http.StatusOK, communityDto)
}

func (s *CommunityServiceImpl) UpdateCommunity(ctx *gin.Context) {
	var req dtos.UpdateCommunityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Invalid request body"))
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", err.Error()))
		return
	}

	community := &models.Community{}

	if err := s.communityDB.UpdateCommunity(community); err != nil {
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
	community, err := s.communityDB.GetCommunityByID(&id)

	type UserInfo struct {
		Username string
		Email    string
	}

	var userDto *dtos.UserDto

	if community.CreatorID != nil {
		userDto, err = s.communityDB.GetUserInfoByUserID(community.CreatorID)
	}

	var topicDtos []dtos.TopicDto

	if len(community.Topics) > 0 {
		for _, topic := range community.Topics {
			topicDto := dtos.TopicDto{
				ID:          topic.ID,
				Name:        topic.Name,
				Description: topic.Description,
			}
			topicDtos = append(topicDtos, topicDto)
		}
	}

	communityDto := dtos.CommunityDto{
		ID:          community.ID,
		Name:        community.Name,
		Description: community.Description,
		Type:        community.Type,
		CreatorId:   community.CreatorID,
		BannerImage: community.BannerImage,
		IconImage:   community.IconImage,
		CreatedAt:   &community.CreatedAt,
		CreatorName: &userDto.Username,
		Topics:      topicDtos,
	}

	if err != nil {
		ctx.Error(core.ErrInternalServerError.WithDetail("error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, communityDto)
}

func (s *CommunityServiceImpl) ListCommunities(ctx *gin.Context) {
	communities, err := s.communityDB.ListCommunities()
	if err != nil {
		ctx.Error(core.ErrInternalServerError.WithDetail("error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, communities)
}

func (s *CommunityServiceImpl) JoinCommunity(ctx *gin.Context) {
	var req dtos.JoinCommunityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Invalid request body"))
		return
	}
	if err := s.communityDB.JoinCommunity(&req.UserID, &req.CommunityID); err != nil {
		ctx.Error(core.ErrInternalServerError.WithDetail("error", "Failed to join community: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Joined community successfully"})
}

func (s *CommunityServiceImpl) LeaveCommunity(ctx *gin.Context) {
	var req dtos.LeaveCommunityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Invalid request body"))
		return
	}
	if err := s.communityDB.LeaveCommunity(&req.UserID, &req.CommunityID); err != nil {
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
	cnt, err := s.communityDB.GetNumberOfMembersInCommunity(&id)
	if err != nil {
		ctx.Error(core.ErrInternalServerError.WithDetail("error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"member_count": cnt})
}

func (s *CommunityServiceImpl) IsUserMember(ctx *gin.Context) {

	var req dtos.IsUserMemberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Invalid request body"))
		return
	}
	isMember, err := s.communityDB.IsUserMember(&req.UserID, &req.CommunityID)
	if err != nil {
		ctx.Error(core.ErrInternalServerError.WithDetail("error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"is_member": isMember,
	})
}

func (s *CommunityServiceImpl) GetAllCommunityByTopic(ctx *gin.Context) {
	topicId := ctx.Param("id")
	if topicId == "" {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Topic ID is required"))
		return
	}

	communities, err := s.communityDB.GetAllCommunityByTopicId(&topicId)
	if err != nil {
		ctx.Error(core.ErrInternalServerError.WithDetail("error", err.Error()))
		return
	}

	var communityDtos []dtos.CommunityDto
	for _, community := range communities {

		var banner_image string
		var icon_image string

		if community.BannerImage != nil {
			banner_image = path.Join(environment.AppName+core.V1, *community.BannerImage)
		}
		if community.IconImage != nil {
			icon_image = path.Join(environment.AppName+core.V1, *community.IconImage)
		}

		communityDtos = append(communityDtos, dtos.CommunityDto{
			IconImage:   &icon_image,
			BannerImage: &banner_image,
			Name:        community.Name,
			Description: community.Description,
			ID:          community.ID,
			CreatedAt:   &community.CreatedAt,
			UpdatedAt:   &community.UpdatedAt,
		})
	}

	ctx.JSON(http.StatusOK, communityDtos)
}
