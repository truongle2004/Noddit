package impl

import (
	"blog-service/internal/domain/models"
	"blog-service/internal/dtos"
	"blog-service/internal/utils"
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

func (p *PostRepositoryImpl) GetAllPostsByCommunityId(ctx *gin.Context, communityId string, pagination utils.Pagination) ([]*dtos.PostDto, int64, error) {
	var posts []*dtos.PostDto
	var total int64

	// Count total posts
	const countSQL = "SELECT COUNT(*) FROM posts WHERE community_id = $1"
	if err := p.db.WithContext(ctx).Raw(countSQL, communityId).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := pagination.GetOffset()
	const selectSQL = `SELECT p.id, p.title, p.slug, p.content, p.creator_id, 
						p.community_id, p.tags, p.post_type, p.read_time, 
						p.meta_title, p.meta_desc, p.created_at,
						u.username as creator_name, u.id as creator_id 
					FROM posts p
					JOIN users u on p.creator_id = u.id
					WHERE community_id = $1
					ORDER BY p.created_at DESC LIMIT $2 OFFSET $3`
	if err := p.db.WithContext(ctx).Debug().Raw(selectSQL, communityId, pagination.Limit, offset).Scan(&posts).Error; err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

//
//func (p *PostRepositoryImpl) CheckCreatorIdAndCommunityId(ctx *gin.Context, creatorId, communityId string) (bool, error) {
//
//}
