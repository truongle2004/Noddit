package repositories

import (
	"subnoddit-service/internal/domain/models"
)

type SubredditRepository interface {
	CreateCommunity(community *models.Community) error
	UpdateCommunity(community *models.Community) error
	GetCommunityByID(id *string) (*models.Community, error)
	ListCommunities() ([]models.Community, error)

	GetNumberOfMembersInCommunity(communityId *string) (int64, error)
	JoinCommunity(userID, communityID *string) error
	LeaveCommunity(userID, communityID *string) error
	IsUserMember(userID, communityID *string) (bool, error)
}
