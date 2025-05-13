package services

import "github.com/gin-gonic/gin"

type TopicService interface {
	GetTopics(ctx *gin.Context)
	CreateTopic(ctx *gin.Context)
	GetTopicByID(ctx *gin.Context)
	UpdateTopic(ctx *gin.Context)
	DeleteTopic(ctx *gin.Context)
	GetTopicsByCommunityID(ctx *gin.Context)
	AddTopicsToCommunity(ctx *gin.Context)
}
