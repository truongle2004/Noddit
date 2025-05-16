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
	}

	root := "/subnoddit-service"
	subnoddit := r.Group(constant.V1 + root)
	{
		// Create community
		subnoddit.POST("/communities",
			middleware.AuthMiddleware(),
			middleware.ValidTokenMiddleware(),
			proxy.ReserveProxy(environment.SubnodditServiceRoute),
		)

		// Update community (requires ID)
		subnoddit.PUT("/communities/:id",
			middleware.AuthMiddleware(),
			middleware.ValidTokenMiddleware(),
			proxy.ReserveProxy(environment.SubnodditServiceRoute),
		)

		// Get all topics in community
		subnoddit.GET("/communities/:id/topics",
			proxy.ReserveProxy(environment.SubnodditServiceRoute),
		)

		// Get community by ID
		subnoddit.GET("/communities/:id",
			proxy.ReserveProxy(environment.SubnodditServiceRoute),
		)

		// List all communities
		subnoddit.GET("/communities",
			proxy.ReserveProxy(environment.SubnodditServiceRoute),
		)

		// Join community
		subnoddit.POST("/communities/:id/members",
			middleware.AuthMiddleware(),
			middleware.ValidTokenMiddleware(),
			proxy.ReserveProxy(environment.SubnodditServiceRoute),
		)

		// Leave community
		subnoddit.DELETE("/communities/:id/members",
			middleware.AuthMiddleware(),
			middleware.ValidTokenMiddleware(),
			proxy.ReserveProxy(environment.SubnodditServiceRoute),
		)

		// Check membership
		subnoddit.GET("/communities/:id/members/me",
			middleware.AuthMiddleware(),
			middleware.ValidTokenMiddleware(),
			proxy.ReserveProxy(environment.SubnodditServiceRoute),
		)

		// Get member count
		subnoddit.GET("/communities/:id/members/count",
			proxy.ReserveProxy(environment.SubnodditServiceRoute),
		)

		subnoddit.POST("/image/upload",
			middleware.AuthMiddleware(),
			middleware.ValidTokenMiddleware(),
			proxy.ReserveProxy(environment.SubnodditServiceRoute),
		)

		subnoddit.GET("/image/:filename",
			proxy.ReserveProxy(environment.SubnodditServiceRoute))

		// get all topics
		subnoddit.GET("/topics",
			proxy.ReserveProxy(environment.SubnodditServiceRoute))

		// get topic by ID
		subnoddit.GET("/topics/:id",
			proxy.ReserveProxy(environment.SubnodditServiceRoute))

		// create new topic
		subnoddit.POST("/topics",
			middleware.AuthMiddleware(),
			middleware.ValidTokenMiddleware(),
			proxy.ReserveProxy(environment.SubnodditServiceRoute))

		// update topic
		subnoddit.PUT("/topics/:id",
			proxy.ReserveProxy(environment.SubnodditServiceRoute))

		// delete topic
		subnoddit.DELETE("/topics/:id",
			middleware.AuthMiddleware(),
			middleware.ValidTokenMiddleware(),
			proxy.ReserveProxy(environment.SubnodditServiceRoute))

	}

	return r
}
