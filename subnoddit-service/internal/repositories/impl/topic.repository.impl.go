package impl

import (
	"subnoddit-service/internal/domain/models"

	"gorm.io/gorm"
)

type TopicRepository struct {
	db *gorm.DB
}

func NewTopicRepository(db *gorm.DB) *TopicRepository {
	return &TopicRepository{
		db: db,
	}
}

func (r *TopicRepository) GetTopics() ([]models.Topic, error) {
	var topics []models.Topic
	if err := r.db.Find(&topics).Error; err != nil {
		return nil, err
	}
	return topics, nil
}

func (r *TopicRepository) CreateTopic(topic *models.Topic) error {
	return r.db.Create(topic).Error
}

func (r *TopicRepository) GetTopicByID(id *string) (*models.Topic, error) {
	var topic models.Topic
	if err := r.db.First(&topic, "id = ?", *id).Error; err != nil {
		return nil, err
	}
	return &topic, nil
}

func (r *TopicRepository) UpdateTopic(topic *models.Topic) error {
	return r.db.Save(topic).Error
}

func (r *TopicRepository) DeleteTopic(id *string) error {
	return r.db.Delete(&models.Topic{}, "id = ?", *id).Error
}

func (r *TopicRepository) GetTopicsByCommunityID(communityID *string) ([]models.Topic, error) {
	var topics []models.Topic
	err := r.db.
		Joins("JOIN community_topics ct ON ct.topic_id = topics.id").
		Where("ct.community_id = ?", *communityID).
		Find(&topics).Error

	return topics, err
}

func (r *TopicRepository) AddTopicsToCommunity(communityID string, topicIDs []string) error {
	var community models.Community
	if err := r.db.First(&community, "id = ?", communityID).Error; err != nil {
		return err
	}

	var topics []models.Topic
	if err := r.db.Where("id IN ?", topicIDs).Find(&topics).Error; err != nil {
		return err
	}

	return r.db.Model(&community).Association("Topics").Replace(&topics)
}
