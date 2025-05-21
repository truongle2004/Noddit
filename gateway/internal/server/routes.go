package server

import (
	"gateway/internal/config"
	"gateway/internal/constant"
	"gateway/internal/environment"
	"gateway/internal/middleware"
	"gateway/internal/proxy"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {

	r := gin.Default()
	r.Use(cors.New(config.CorsConfig()))
	//r.Use(servicecontext.ResponseFormatterMiddleware()) // FIXME: ResponseFormatterMiddleware causes dragging response

	// rbacSvc := impl.NewRBACService(config.EnforcerInstance)
	if err := config.NewRedisClient(); err != nil {
		panic(err)
	}

	SubnodditServiceRoutes(r)
	AuthServiceRoutes(r)
	PostServiceRoutes(r)

	return r
}

func AuthServiceRoutes(r *gin.Engine) {
	root := "/auth-service"
	auth := r.Group(constant.V1 + root)
	{
		auth.POST("/login",
			proxy.ReserveProxy(environment.AuthServiceRoute),
		)

		auth.POST("/register",
			proxy.ReserveProxy(environment.AuthServiceRoute),
		)

		auth.GET("/refresh-token",
			middleware.AuthMiddleware(),
			middleware.ValidTokenMiddleware(),
			proxy.ReserveProxy(environment.AuthServiceRoute),
		)

		auth.GET("/logout",
			middleware.AuthMiddleware(),
			middleware.ValidTokenMiddleware(),
			proxy.ReserveProxy(environment.AuthServiceRoute),
		)
	}
}

func SubnodditServiceRoutes(r *gin.Engine) {
	root := "/subnoddit-service"
	subnoddit := r.Group(constant.V1 + root)
	{
		// Create community
		subnoddit.POST("/communities",
			middleware.AuthMiddleware(),
			middleware.ValidTokenMiddleware(),
			middleware.AccountMiddleware(),
			proxy.ReserveProxy(environment.SubnodditServiceRoute),
		)

		// Update community (requires ID)
		subnoddit.PUT("/communities/:id",
			middleware.AuthMiddleware(),
			middleware.ValidTokenMiddleware(),
			middleware.AccountMiddleware(),
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
			middleware.AccountMiddleware(),
			proxy.ReserveProxy(environment.SubnodditServiceRoute),
		)

		// Leave community
		subnoddit.DELETE("/communities/:id/members",
			middleware.AuthMiddleware(),
			middleware.ValidTokenMiddleware(),
			middleware.AccountMiddleware(),
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
			middleware.AccountMiddleware(),
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

		// TODO: do it later for admin
		//// create new topic
		//subnoddit.POST("/topics",
		//	middleware.AuthMiddleware(),
		//	middleware.ValidTokenMiddleware(),
		//	proxy.ReserveProxy(environment.SubnodditServiceRoute))
		//
		//// update topic
		//subnoddit.PUT("/topics/:id",
		//	proxy.ReserveProxy(environment.SubnodditServiceRoute))
		//
		//// delete topic
		//subnoddit.DELETE("/topics/:id",
		//	middleware.AuthMiddleware(),
		//	middleware.ValidTokenMiddleware(),
		//	proxy.ReserveProxy(environment.SubnodditServiceRoute))

	}
}

func PostServiceRoutes(r *gin.Engine) {
	root := "/post-service"
	post := r.Group(constant.V1 + root)
	{
		post.POST("/create",
			middleware.AuthMiddleware(),
			middleware.ValidTokenMiddleware(),
			middleware.AccountMiddleware(),
			proxy.ReserveProxy(environment.PostServiceRoute))

		post.GET("/image/:filename",
			proxy.ReserveProxy(environment.PostServiceRoute))

		post.GET("/:id", proxy.ReserveProxy(environment.PostServiceRoute))
	}
}
