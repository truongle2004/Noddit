package impl

import (
	"encoding/json"
	"fmt"
	"time"

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

func (r *RedisServiceImpl) Set(c *gin.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(c.Request.Context(), key, value, expiration).Err()
}

func (r *RedisServiceImpl) Get(c *gin.Context, key string) (string, error) {
	return r.client.Get(c.Request.Context(), key).Result()
}

func (r *RedisServiceImpl) Delete(c *gin.Context, key string) error {
	return r.client.Del(c.Request.Context(), key).Err()
}

func (r *RedisServiceImpl) Exists(c *gin.Context, key string) (bool, error) {
	n, err := r.client.Exists(c.Request.Context(), key).Result()
	return n > 0, err
}

func (r *RedisServiceImpl) SetJSON(c *gin.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(c.Request.Context(), key, data, expiration).Err()
}

func (r *RedisServiceImpl) GetJSON(c *gin.Context, key string, dest interface{}) error {
	data, err := r.client.Get(c.Request.Context(), key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data), dest)
}

func (r *RedisServiceImpl) MatchingKey(c *gin.Context, key string) (bool, error) {
	pattern := fmt.Sprintf("*%s*", key)

	keys, err := r.client.Keys(c.Request.Context(), pattern).Result()
	if err != nil {
		return false, err
	}

	return len(keys) > 0, nil
}
