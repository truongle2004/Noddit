package server

import (
	"fmt"
	"log"
	"net/http"
	"profile-service/internal/environment"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
}

func NewServer() *http.Server {

	environment.LoadConfig()

	NewServer := &Server{
		port: environment.PORT,
	}

	log.Println("Server is running in port =>", environment.PORT)

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
