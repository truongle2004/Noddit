package impl

import (
	"blog-service/internal/domain/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepositoryImpl {
	return &PostRepositoryImpl{
		db: db,
	}
}

func (p *PostRepositoryImpl) Create(ctx *gin.Context, post *models.Post) (*models.Post, error) {
	if err := p.db.WithContext(ctx).Create(&post).Error; err != nil {
		return nil, err
	}
	return post, nil
}
