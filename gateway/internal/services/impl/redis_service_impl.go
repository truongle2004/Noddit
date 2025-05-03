package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type RedisServiceImpl struct {
	client *redis.Client
}

func NewRedisService(client *redis.Client) *RedisServiceImpl {
	return &RedisServiceImpl{
		client: client,
	}
}

func (r *RedisServiceImpl) Get(c *gin.Context, key string) (string, error) {
	return r.client.Get(c.Request.Context(), key).Result()
}
