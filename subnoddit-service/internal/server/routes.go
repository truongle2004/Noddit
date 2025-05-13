package server

import (
	"net/http"
	"subnoddit-service/internal/config"
	"subnoddit-service/internal/controller"
	repo "subnoddit-service/internal/repositories/impl"
	svc "subnoddit-service/internal/services/impl"

	"github.com/gin-contrib/cors"
	"github.com/truongle2004/service-context/middleware"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(middleware.ResponseFormatterMiddleware())
	r.Use(cors.New(config.CorsConfig()))

	topicRepo := repo.NewTopicRepository(config.DbInstance)
	communityRepo := repo.NewCommunityRepository(config.DbInstance)

	imageSvc := svc.NewImageService()
	imageCtrl := controller.NewImageController(imageSvc)
	imageCtrl.RegisterRoutes(r)

	communitySvc := svc.NewCommunityService(communityRepo, topicRepo)
	communityCtrl := controller.NewCommunityController(communitySvc)
	communityCtrl.RegisterRoutes(r)

	topicSerivce := svc.NewTopicService(topicRepo)
	topicCtrl := controller.NewTopicController(topicSerivce)
	topicCtrl.RegisterRoutes(r)

	return r
}
