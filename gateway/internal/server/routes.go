package server

import (
	"context"
	"gateway/internal/config"
	"gateway/internal/constant"
	"gateway/internal/environment"
	"gateway/internal/middleware"
	"gateway/internal/proxy"
	"gateway/internal/services/impl"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func initRedisConnection() *redis.Client {
	client := config.NewRedisClient()
	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	} else {
		log.Println("✅ Redis connection established.")
	}

	if err := config.InitRBAC(); err != nil {
		log.Println("failed to initialize RBAC:", err)
	}

	return client
}

func (s *Server) RegisterRoutes() http.Handler {

	r := gin.Default()
	r.Use(cors.New(config.CorsConfig()))

	// rbacSvc := impl.NewRBACService(config.EnforcerInstance)
	client := initRedisConnection()
	redisSvc := impl.NewRedisService(client)
	middleware.NewAccountMiddleware(redisSvc)

	authGroup := r.Group(constant.V1)
	{
		authGroup.Any("/auth/*proxyPath",
			proxy.ReserveProxy(environment.AuthServiceRoute))

		// authGroup.Any("/users/*proxyPath",
		// 	middleware.AuthMiddleware(),
		// 	middleware.ValidTokenMiddleware(),
		// 	middleware.RBACMiddleware(rbacSvc),
		// 	proxy.ReserveProxy(environment.AuthServiceRoute))
	}

	subnoddit := r.Group(constant.V1)
	{
		subnoddit.Any("/subnoddit-service/*proxyPath",
			middleware.AuthMiddleware(),
			middleware.ValidTokenMiddleware(),
			//middleware.RBACMiddleware(rbacSvc),
			proxy.ReserveProxy(environment.SubnodditServiceRoute))

	}

	profileGroup := r.Group(constant.V1)
	{
		profileGroup.Any("/profile-service/*proxyPath",
			middleware.AuthMiddleware(),
			middleware.ValidTokenMiddleware(),
			//middleware.RBACMiddleware(rbacSvc),
			proxy.ReserveProxy(environment.ProfileServiceRoute))
	}

	return r
}
