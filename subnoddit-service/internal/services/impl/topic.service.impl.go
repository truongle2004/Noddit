package impl

import (
	"net/http"
	"subnoddit-service/internal/domain/models"
	"subnoddit-service/internal/dtos"
	"subnoddit-service/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
)

type TopicServiceImpl struct {
	topicRepo repositories.TopicRepository
}

func NewTopicService(topicRepo repositories.TopicRepository) *TopicServiceImpl {
	return &TopicServiceImpl{
		topicRepo: topicRepo,
	}
}

func (s *TopicServiceImpl) GetTopics(ctx *gin.Context) {
	topics, err := s.topicRepo.GetTopics()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get topics"})
		return
	}

	var topicsDTO []dtos.TopicDto
	for _, topic := range topics {
		topicsDTO = append(topicsDTO, dtos.TopicDto{
			ID:          topic.ID,
			Name:        topic.Name,
			Description: topic.Description,
		})
	}

	ctx.JSON(http.StatusOK, topicsDTO)
}

func (s *TopicServiceImpl) CreateTopic(ctx *gin.Context) {
	var topic models.Topic
	if err := ctx.ShouldBindJSON(&topic); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := s.topicRepo.CreateTopic(&topic); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create topic"})
		return
	}
	ctx.JSON(http.StatusCreated, topic)
}

func (s *TopicServiceImpl) GetTopicByID(ctx *gin.Context) {
	id := ctx.Param("id")
	topic, err := s.topicRepo.GetTopicByID(&id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Topic not found"})
		return
	}
	ctx.JSON(http.StatusOK, topic)
}

func (s *TopicServiceImpl) UpdateTopic(ctx *gin.Context) {
	id := ctx.Param("id")
	var topic models.Topic
	if err := ctx.ShouldBindJSON(&topic); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Invalid request body"))
		return
	}
	topic.ID = id
	if err := s.topicRepo.UpdateTopic(&topic); err != nil {
		ctx.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", "Failed to update topic: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, topic)
}

func (s *TopicServiceImpl) DeleteTopic(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := s.topicRepo.DeleteTopic(&id); err != nil {
		ctx.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", "Failed to delete topic: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Topic deleted"})
}

func (s *TopicServiceImpl) GetTopicsByCommunityID(ctx *gin.Context) {
	communityID := ctx.Param("community_id")
	topics, err := s.topicRepo.GetTopicsByCommunityID(&communityID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", "Failed to get topics: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, topics)
}

func (s *TopicServiceImpl) AddTopicsToCommunity(ctx *gin.Context) {
	var payload struct {
		CommunityID string   `json:"community_id"`
		TopicIDs    []string `json:"topic_ids"`
	}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithDetail("error", "Invalid request body"))
		return
	}
	if err := s.topicRepo.AddTopicsToCommunity(payload.CommunityID, payload.TopicIDs); err != nil {
		ctx.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", "Failed to add topics to community: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Topics added to community"})
}
