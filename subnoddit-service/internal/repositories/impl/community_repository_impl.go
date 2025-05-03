package impl

import (
	"subnoddit-service/internal/domain/models"

	"gorm.io/gorm"
)

type SubredditRepositoryImpl struct {
	db *gorm.DB
}

func NewSubredditRepository(db *gorm.DB) *SubredditRepositoryImpl {
	return &SubredditRepositoryImpl{db: db}
}

func (r *SubredditRepositoryImpl) CreateCommunity(community *models.Community) error {
	return r.db.Create(community).Error
}

func (r *SubredditRepositoryImpl) UpdateCommunity(community *models.Community) error {
	return r.db.Save(community).Error
}

func (r *SubredditRepositoryImpl) GetCommunityByID(id *string) (*models.Community, error) {
	var community models.Community
	err := r.db.Where("id = ?", &id).First(&community).Error
	return &community, err
}

func (r *SubredditRepositoryImpl) ListCommunities() ([]models.Community, error) {
	var communities []models.Community
	err := r.db.Find(&communities).Error
	return communities, err
}

func (r *SubredditRepositoryImpl) JoinCommunity(userID, communityID *string) error {
	communityMember := models.CommunityMember{
		UserID:      *userID,
		CommunityID: *communityID,
	}
	return r.db.Create(&communityMember).Error
}

func (r *SubredditRepositoryImpl) LeaveCommunity(userID, communityID *string) error {
	return r.db.
		Where("user_id = ? AND community_id = ?", userID, communityID).
		Delete(&models.CommunityMember{}).Error
}

func (r *SubredditRepositoryImpl) IsUserMember(userID, communityID *string) (bool, error) {
	var count int64
	err := r.db.Model(&models.CommunityMember{}).
		Where("user_id = ? AND community_id = ?", userID, communityID).
		Count(&count).Error
	return count > 0, err
}

func (r *SubredditRepositoryImpl) GetNumberOfMembersInCommunity(communityID *string) (int64, error) {
	var count int64
	err := r.db.Model(&models.CommunityMember{}).
		Where("community_id = ?", communityID).
		Count(&count).Error
	return count, err
}
