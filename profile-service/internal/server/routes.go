package server

import (
	"log"
	"net/http"
	"profile-service/internal/config"
	"profile-service/internal/rest"

	repositories "profile-service/internal/repositories/impl"
	services "profile-service/internal/services/impl"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(config.CorsConfig())

	driver, err := config.InitNeo4j()
	if err != nil {
		log.Println("❌ Failed to connect to Neo4j:", err)
	}

	followRepo := repositories.NewFollowRepository(driver)
	followSvc := services.NewFollowService(followRepo)
	profileRepo := repositories.NewProfileRepository(driver)
	profileSvc := services.NewProfileServiceImpl(profileRepo)
	rest.NewProfileController(profileSvc).RegisterRoute(r)
	rest.NewFollowController(followSvc).RegisterRoute(r)

	return r
}
