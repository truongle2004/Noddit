package repositories

import (
	"blog-service/internal/domain/models"

	"github.com/gin-gonic/gin"
)

type PostRepository interface {
	CreateNewPost(ctx *gin.Context, post *models.Post) (*models.Post, error)
	UploadImageUrl(ctx *gin.Context, sql string) error
	GetAllImageUrlByPostId(ctx *gin.Context, postId string) ([]models.PostImage, error)
	// GetByID(id uint) (*domain.Post, error)
	// GetAll() ([]domain.Post, error)
	// Update(post *domain.Post) error
	// Delete(id uint) error
}
