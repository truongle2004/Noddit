package config

import (
	"blog-service/internal/environment"

	"github.com/truongle2004/service-context/redisclient"
)

func InitRedis() error {
	return redisclient.InitRedis(environment.RedisAddr, environment.RedisPass, environment.RedisDB)
}
