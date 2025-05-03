package environment

import (
	"github.com/spf13/viper"
	"github.com/truongle2004/service-context/core"
	"log"
	"os"
)

var (
	PORT int
)

var (
	DbName     string
	DbUsername string
	DbPassword string
	DbHost     string
	DbPort     int
)

// LoadConfig load centralized configuration for the application
func LoadConfig() {
	v := viper.New()

	AppName := os.Getenv("APP_NAME")
	AppProfile := os.Getenv("APP_PROFILE")
	ConfigURL := os.Getenv("CONFIG_URL")

	if err := core.LoadConfig(AppName, AppProfile, ConfigURL, v); err != nil {
		log.Println("Error loading config:", err)
	}

	rootEnv := "app.environment"
	PORT = v.GetInt("server.port")

	DbUsername = v.GetString(rootEnv + ".DB_USER")
	DbPassword = v.GetString(rootEnv + ".DB_PASSWORD")
	DbHost = v.GetString(rootEnv + ".DB_HOST")
	DbPort = v.GetInt(rootEnv + ".DB_PORT")
	DbName = v.GetString(rootEnv + ".DB_NAME")
}
