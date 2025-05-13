package server

import (
	"fmt"
	"log"
	"net/http"
	"subnoddit-service/internal/config"
	"subnoddit-service/internal/environment"
	"subnoddit-service/internal/seeds"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
}

func LoadConfig() {
	environment.LoadConfig()
	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	if err := config.ImageUploadConfig(); err != nil {
		log.Println("There are errors during creating new upload folder")
	}
	if err := seeds.SeedTopics(config.DbInstance); err != nil {
		log.Fatalf("Failed to seed topics: %v", err)
	}
}

func NewServer() *http.Server {
	LoadConfig()

	log.Println("Server is running on port:", environment.PORT)

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
