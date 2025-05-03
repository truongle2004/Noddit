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
	Neo4jUsername string
	Neo4jPassword string
	Neo4jUri      string
)

// LoadConfig Load centralized configuration
func LoadConfig() {
	v := viper.New()

	AppName := os.Getenv("APP_NAME")
	AppProfile := os.Getenv("APP_PROFILE")
	ConfigURL := os.Getenv("CONFIG_URL")

	err := core.LoadConfig(AppName, AppProfile, ConfigURL, v)
	if err != nil {
		log.Println("❌ Load config server failed:", err)
		return
	}

	rootEnv := "app.environment"
	PORT = v.GetInt("server.port")
	Neo4jUsername = v.GetString(rootEnv + ".NEO4J_USERNAME")
	Neo4jPassword = v.GetString(rootEnv + ".NEO4J_PASSWORD")
	Neo4jUri = v.GetString(rootEnv + ".NEO4J_URI")
}
