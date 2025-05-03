package services

import (
	"time"

	"github.com/gin-gonic/gin"
)

type RedisService interface {
	Set(c *gin.Context, key string, value interface{}, expiration time.Duration) error
	Get(c *gin.Context, key string) (string, error)
	Delete(c *gin.Context, key string) error
	Exists(c *gin.Context, key string) (bool, error)
	SetJSON(c *gin.Context, key string, value interface{}, expiration time.Duration) error
	GetJSON(c *gin.Context, key string, dest interface{}) error
	MatchingKey(c *gin.Context, key string) (bool, error)
}
