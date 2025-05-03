package server

import (
	"net/http"
	"notification-service/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(config.CorsConfig()))

	_, err := config.InitDB()
	if err != nil {
		return nil
	}

	return r
}
