package config

import (
	"gateway/internal/environment"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     environment.RedisAddr,
		Password: environment.RedisPassword,
		DB:       environment.RedisDb,
	})
}
