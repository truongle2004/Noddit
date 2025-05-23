package server

import (
	"blog-service/internal/environment"
	"fmt"
	"github.com/truongle2004/service-context/logger"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
}

func NewServer() *http.Server {
	environment.LoadConfig()

	logger.Infof("Server is running in post: %d", environment.PORT)

	NewServer := &Server{
		port: environment.PORT,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
