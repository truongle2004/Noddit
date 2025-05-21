package repositories

import (
	"blog-service/internal/domain/models"
	"blog-service/internal/dtos"
	"blog-service/internal/utils"

	"github.com/gin-gonic/gin"
)

type PostRepository interface {
	CreateNewPost(ctx *gin.Context, post *models.Post) (*models.Post, error)
	UploadImageUrl(ctx *gin.Context, sql string) error
	GetAllImageUrlByPostId(ctx *gin.Context, postId string) ([]models.PostImage, error)
	GetAllPostsByCommunityId(ctx *gin.Context, communityId string, pagination utils.Pagination) ([]*dtos.PostDto, int64, error)
	// GetByID(id uint) (*domain.Post, error)
	// GetAll() ([]domain.Post, error)
	// Update(post *domain.Post) error
	// Delete(id uint) error
}
