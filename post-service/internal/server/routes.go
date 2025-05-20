package server

import (
	"blog-service/internal/config"
	"blog-service/internal/controller"
	repo "blog-service/internal/repositories/impl"
	svc "blog-service/internal/services/impl"
	"github.com/truongle2004/service-context/utils"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/logger"
	"github.com/truongle2004/service-context/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	err := utils.ImageUploadConfig()
	if err != nil {
		logger.Errorf("failed to upload image: %v", err.Error())
	}

	dbErr := config.InitDB()
	if dbErr != nil {
		logger.Errorf("failed to connect to PostgreSQL database: %v", dbErr.Error())
	}

	// redisErr := config.InitRedis()
	// if redisErr != nil {
	// 	logger.Errorf("failed to connect to PostgreSQL database: %v", redisErr.Error())
	// } else {
	// 	logger.Info("redis connected!!")
	// }
	r := gin.Default()

	r.Use(middleware.ResponseFormatterMiddleware())
	r.Use(cors.New(config.CorsConfig()))

	postRepo := repo.NewPostRepository(config.DbInstance)
	postSvc := svc.NewPostService(postRepo)
	postCtrl := controller.NewCommunityController(postSvc)
	postCtrl.RegisterRoutes(r)

	imageCtrl := controller.NewImageController()
	imageCtrl.RegisterRoutes(r)

	return r
}
