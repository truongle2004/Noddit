package repositories

import (
	"subnoddit-service/internal/domain/models"
)

type TopicRepository interface {
	GetTopics() ([]models.Topic, error)
	CreateTopic(topic *models.Topic) error
	GetTopicByID(id *string) (*models.Topic, error)
	UpdateTopic(topic *models.Topic) error
	DeleteTopic(id *string) error
	GetTopicsByCommunityID(communityID *string) ([]models.Topic, error)
	AddTopicsToCommunity(communityID string, topicIDs []string) error
}
