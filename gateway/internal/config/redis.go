package config

import (
	"gateway/internal/environment"
	"github.com/truongle2004/service-context/redisclient"
)

func NewRedisClient() error {
	return redisclient.InitRedis(environment.RedisAddr, environment.RedisPassword, environment.RedisDb)
	//return redis.NewClient(&redis.Options{
	//	Addr:     environment.RedisAddr,
	//	Password: environment.RedisPassword,
	//	DB:       environment.RedisDb,
	//})
}
