package impl

import (
	"blog-service/internal/constant"
	"blog-service/internal/domain/models"
	"blog-service/internal/dtos"
	"blog-service/internal/enums"
	"blog-service/internal/repositories"
	utils2 "blog-service/internal/utils"
	"github.com/gosimple/slug"
	"github.com/truongle2004/service-context/core"
	"github.com/truongle2004/service-context/logger"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/utils"
)

type PostServiceImpl struct {
	postRepo repositories.PostRepository
}

func NewPostService(postRepository repositories.PostRepository) *PostServiceImpl {
	return &PostServiceImpl{
		postRepo: postRepository,
	}
}

func (p *PostServiceImpl) CreateNewPost(ctx *gin.Context) {
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	creatorID := ctx.PostForm("creator_id")
	communityID := ctx.PostForm("community_id")
	tags := ctx.PostForm("tags")
	postType := ctx.PostForm("post_type")
	metaTitle := ctx.PostForm("meta_title")
	metaDesc := ctx.PostForm("meta_desc")
	readTimeStr := ctx.PostForm("read_time")
	readTime, err := strconv.Atoi(readTimeStr)

	titleSlug := slug.Make(title)

	postId := uuid.New().String()

	post := models.Post{
		SQLModel: core.SQLModel{
			ID: postId,
		},
		Title:       &title,
		Slug:        &titleSlug,
		MetaTitle:   &metaTitle,
		MetaDesc:    &metaDesc,
		ReadTime:    &readTime,
		PostType:    &postType,
		Content:     &content,
		CreatorID:   &creatorID,
		CommunityID: &communityID,
		Tags:        &tags,
	}

	result, err := p.postRepo.CreateNewPost(ctx, &post)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			ctx.JSON(http.StatusConflict, core.ErrConflict.WithDetail("error", err.Error()))
		} else {
			logger.Errorf("failed to create new post: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", err.Error()))
		}
		return
	}

	postDto := dtos.PostDto{
		ID:          &result.ID,
		Title:       result.Title,
		Slug:        result.Slug,
		Content:     result.Content,
		CreatorId:   result.CreatorID,
		CommunityID: result.CommunityID,
		Tags:        result.Tags,
		PostType:    result.PostType,
		ReadTime:    result.ReadTime,
		MetaTitle:   result.MetaTitle,
		MetaDesc:    result.MetaDesc,
	}

	var listImageUrl []string

	switch enums.GetEnumPostType(&postType) {
	case enums.PostTypeImage:
		form, _ := ctx.MultipartForm()
		files := form.File["file"]

		for _, file := range files {
			dst := "/image_" + string(uuid.New().String()) + postId + "_" + file.Filename
			listImageUrl = append(listImageUrl, dst)
			if err := utils.SaveImageToFolder(ctx, constant.ImagePath+dst, file); err != nil {
				logger.Errorf("failed to upload image: %v", err.Error())
				ctx.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", err.Error()))
				return
			}
		}

		if len(listImageUrl) > 0 {
			sql := "INSERT INTO post_images (id, post_id, image_url, created_at, updated_at) VALUES "

			for _, imageUrl := range listImageUrl {
				sql += "('" + uuid.New().String() + "', '" + postId + "', '" + imageUrl + "', NOW(), NOW()),"
			}

			if err := p.postRepo.UploadImageUrl(ctx, sql[:len(sql)-1]); err != nil {
				logger.Errorf("failed to upload image: %v", err.Error())
				ctx.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", err.Error()))
				return
			}
		}

		postDto.ImageUrls = &listImageUrl
		break
	case enums.PostTypeText:
		// do something
		break
	case enums.PostTypeVideo:
		// do something
		break
	}

	ctx.JSON(http.StatusOK, postDto)
}

func (p *PostServiceImpl) GetAllPostsByCommunityId(ctx *gin.Context) {
	communityId := ctx.Param("id")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	if communityId == "" {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "community_id is required"))
		return
	}

	pa := utils2.Pagination{
		Page:  page,
		Limit: limit,
	}

	postDtos, total, err := p.postRepo.GetAllPostsByCommunityId(ctx, communityId, pa)
	if err != nil {
		logger.Errorf("failed to get all posts: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", err.Error()))
		return
	}

	for _, post := range postDtos {
		urls, _ := p.postRepo.GetAllImageUrlByPostId(ctx, *post.ID)
		var listImageUrl []string
		for _, imageUrl := range urls {
			urlProcessing := "http://localhost:8081/api/v1/post-service/image" + *imageUrl.ImageURL
			listImageUrl = append(listImageUrl, urlProcessing)
		}

		post.ImageUrls = &listImageUrl
	}

	ctx.JSON(http.StatusOK, gin.H{
		"page":  pa.Page,
		"limit": pa.Limit,
		"total": total,
		"posts": postDtos,
	})
}
