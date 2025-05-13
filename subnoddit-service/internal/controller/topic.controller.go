package controller

import (
	"subnoddit-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
)

type TopicController struct {
	TopicService services.TopicService
}

func NewTopicController(topicService services.TopicService) *TopicController {
	return &TopicController{
		TopicService: topicService,
	}
}

func (a *TopicController) RegisterRoutes(r *gin.Engine) {
	api := r.Group(core.V1 + "/subnoddit-service")

	// Topic general routes
	topic := api.Group("/topics")
	{
		topic.GET("", a.TopicService.GetTopics)
		topic.POST("", a.TopicService.CreateTopic)
		topic.GET("/:id", a.TopicService.GetTopicByID)
		topic.PUT("/:id", a.TopicService.UpdateTopic)
		topic.DELETE("/:id", a.TopicService.DeleteTopic)
	}
}
