package config

import (
	"auth-service/internal/environment"
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	redisOnce   sync.Once
)

// InitRedis initializes the Redis client using singleton pattern
func InitRedis() error {
	var initErr error

	redisOnce.Do(func() {
		ctx := context.Background()
		RedisClient = redis.NewClient(&redis.Options{
			Addr: environment.RedisAddr,
		})

		if err := RedisClient.Ping(ctx).Err(); err != nil {
			initErr = err
		}
	})

	return initErr
}
