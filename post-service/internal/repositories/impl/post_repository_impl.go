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

func (p *PostRepositoryImpl) CreateNewPost(ctx *gin.Context, post *models.Post) (*models.Post, error) {
	if err := p.db.WithContext(ctx).Create(&post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (p *PostRepositoryImpl) UploadImageUrl(ctx *gin.Context, sql string) error {
	if err := p.db.WithContext(ctx).Exec(sql).Error; err != nil {
		return err
	}
	return nil
}

func (p *PostRepositoryImpl) GetAllImageUrlByPostId(ctx *gin.Context, postId string) ([]models.PostImage, error) {
	var postImages []models.PostImage
	if err := p.db.WithContext(ctx).Where("post_id = ?", postId).Find(&postImages).Error; err != nil {
		return nil, err
	}
	return postImages, nil
}

//
//func (p *PostRepositoryImpl) CheckCreatorIdAndCommunityId(ctx *gin.Context, creatorId, communityId string) (bool, error) {
//
//}
