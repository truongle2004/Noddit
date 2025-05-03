package services

import (
	"github.com/gin-gonic/gin"
)

type RedisService interface {
	Get(c *gin.Context, key string) (string, error)
}
