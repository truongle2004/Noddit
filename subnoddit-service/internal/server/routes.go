package server

import (
	"net/http"
	"subnoddit-service/internal/config"
	"subnoddit-service/internal/controller"
	repo "subnoddit-service/internal/repositories/impl"
	svc "subnoddit-service/internal/services/impl"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	subNodditRepo := repo.NewSubredditRepository(config.DbInstance)
	subNodditSvc := svc.NewCommunityService(subNodditRepo)
	subNodditCtrl := controller.NewCommunityController(subNodditSvc)
	subNodditCtrl.RegisterRoutes(r)

	r.Use(cors.New(config.CorsConfig()))

	return r
}
