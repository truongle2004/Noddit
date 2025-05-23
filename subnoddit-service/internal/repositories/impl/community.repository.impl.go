package impl

import (
	"subnoddit-service/internal/domain/models"
	"subnoddit-service/internal/dtos"

	"gorm.io/gorm"
)

type CommunityRepositoryImpl struct {
	db *gorm.DB
}

func NewCommunityRepository(db *gorm.DB) *CommunityRepositoryImpl {
	return &CommunityRepositoryImpl{db: db}
}

func (r *CommunityRepositoryImpl) CreateCommunity(community *models.Community) error {
	return r.db.Create(community).Error
}

func (r *CommunityRepositoryImpl) UpdateCommunity(community *models.Community) error {
	return r.db.Save(community).Error
}

func (r *CommunityRepositoryImpl) GetCommunityByID(id *string) (*models.Community, error) {
	var community models.Community
	err := r.db.Preload("Topics").Where("id = ?", &id).First(&community).Error
	return &community, err
}

func (r *CommunityRepositoryImpl) ListCommunities() ([]models.Community, error) {
	var communities []models.Community
	err := r.db.Find(&communities).Error
	return communities, err
}

func (r *CommunityRepositoryImpl) JoinCommunity(userID, communityID *string) error {
	communityMember := models.CommunityMember{
		UserID:      *userID,
		CommunityID: *communityID,
	}
	return r.db.Create(&communityMember).Error
}

func (r *CommunityRepositoryImpl) LeaveCommunity(userID, communityID *string) error {
	return r.db.
		Where("user_id = ? AND community_id = ?", userID, communityID).
		Delete(&models.CommunityMember{}).Error
}

func (r *CommunityRepositoryImpl) IsUserMember(userID, communityID *string) (bool, error) {
	var count int64
	err := r.db.Model(&models.CommunityMember{}).
		Where("user_id = ? AND community_id = ?", userID, communityID).
		Count(&count).Error
	return count > 0, err
}

func (r *CommunityRepositoryImpl) GetNumberOfMembersInCommunity(communityID *string) (int64, error) {
	var count int64
	err := r.db.Model(&models.CommunityMember{}).
		Where("community_id = ?", communityID).
		Count(&count).Error
	return count, err
}

func (r *CommunityRepositoryImpl) GetAllCommunityByTopicId(id *string) ([]*models.Community, error) {
	var topics models.Topic
	err := r.db.Preload("Communities.Topics").First(&topics, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	var communities []*models.Community
	for _, community := range topics.Communities {
		communities = append(communities, community)
	}

	return communities, nil
}

func (r *CommunityRepositoryImpl) GetUserInfoByUserID(userID *string) (*dtos.UserDto, error) {
	var userInfo dtos.UserDto

	err := r.db.Raw("SELECT u.id, u.username, u.email FROM users u WHERE u.id = ?", userID).Scan(&userInfo).Error
	return &userInfo, err

}
