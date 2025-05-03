package services

import "github.com/gin-gonic/gin"

type NotificationConfigService interface {
	GetNotificationConfig(c *gin.Context)
}
