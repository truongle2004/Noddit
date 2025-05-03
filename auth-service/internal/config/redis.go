package config

import (
	"auth-service/internal/environment"
	"github.com/gin-gonic/gin"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	redisOnce   sync.Once
)

// InitRedis initializes the Redis client using singleton pattern
func InitRedis(ctx *gin.Context) error {
	var initErr error

	redisOnce.Do(func() {
		RedisClient = redis.NewClient(&redis.Options{
			Addr: environment.RedisAddr,
		})

		if err := RedisClient.Ping(ctx).Err(); err != nil {
			initErr = err
		}
	})

	return initErr
}
