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

	imageSvc := svc.NewImageService()
	subNodditRepo := repo.NewSubredditRepository(config.DbInstance)
	subNodditSvc := svc.NewCommunityService(subNodditRepo, imageSvc)
	subNodditCtrl := controller.NewCommunityController(subNodditSvc, imageSvc)
	subNodditCtrl.RegisterRoutes(r)

	return r
}
