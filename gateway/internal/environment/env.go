package environment

import (
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/truongle2004/service-context/core"
)

var (
	PORT          int
	RedisAddr     string
	RedisDb       int
	RedisPassword string
	RbacPath      string
)

var (
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
)

var (
	PublicKeyPath string
)

var (
	AuthServiceRoute    string
	ProfileServiceRoute string
)

const (
	rootEnv   = "app.environment"
	rootRoute = "app.routes"
)

// InitConfigServer load the centralized configuration
func InitConfigServer() {
	v := viper.New()

	AppName := os.Getenv("APP_NAME")
	AppProfile := os.Getenv("APP_PROFILE")
	ConfigURL := os.Getenv("CONFIG_URL")

	log.Println("🚀 Start loading config server...")
	if err := core.LoadConfig(AppName, AppProfile, ConfigURL, v); err != nil {
		log.Println("❌ Load config server failed:", err)
		return
	}

	PORT = v.GetInt("server.port")

	// Redis
	RedisAddr = v.GetString(rootEnv + ".REDIS_ADDR")
	RedisDb = v.GetInt(rootEnv + ".REDIS_DB")
	RedisPassword = v.GetString(rootEnv + ".REDIS_PASSWORD")

	// Security
	RbacPath = v.GetString(rootEnv + ".RBAC_PATH")
	PublicKeyPath = v.GetString(rootEnv + ".PUBLIC_KEY_PATH")

	// Routes
	AuthServiceRoute = v.GetString(rootRoute + ".auth-service")
	ProfileServiceRoute = v.GetString(rootRoute + ".profile-service")

	DbHost = v.GetString(rootEnv + ".DB_HOST")
	DbPort = v.GetInt(rootEnv + ".DB_PORT")
	DbUser = v.GetString(rootEnv + ".DB_USER")
	DbPassword = v.GetString(rootEnv + ".DB_PASSWORD")
}
