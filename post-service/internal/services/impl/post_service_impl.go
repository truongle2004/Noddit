package impl

import (
	"blog-service/internal/domain/models"
	"blog-service/internal/dto/request"
	"blog-service/internal/repositories"
	"net/http"

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

func (p *PostServiceImpl) Create(ctx *gin.Context, postDto *request.PostCreateDTO) {
	post := models.Post{
		Title:        postDto.Title,
		Content:      postDto.Content,
		AuthorID:     postDto.AuthorID,
		CommunityID:  postDto.CommunityID,
		Tags:         postDto.Tags,
		PostType:     postDto.PostType,
		MediaURL:     postDto.MediaURL,
		CommentCount: utils.Ptr(0),
		IsEdited:     utils.Ptr(false),
		IsDeleted:    utils.Ptr(false),
	}

	result, err := p.postRepo.Create(ctx, &post)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}
