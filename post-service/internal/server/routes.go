package server

import (
	"blog-service/internal/config"
	"blog-service/internal/controller"
	repo "blog-service/internal/repositories/impl"
	svc "blog-service/internal/services/impl"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	err := config.InitDB()
	if err != nil {
		log.Println("❌ Failed to connect to PostgreSQL database:", err)
	}
	r := gin.Default()

	r.Use(middleware.ResponseFormatterMiddleware())
	r.Use(cors.New(config.CorsConfig()))

	postRepo := repo.NewPostRepository(config.DbInstance)
	postSvc := svc.NewPostService(postRepo)
	postCtrl := controller.NewCommunityController(postSvc)
	postCtrl.RegisterRoutes(r)

	return r
}
