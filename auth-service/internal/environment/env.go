package environment

import (
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/truongle2004/service-context/core"
)

var (
	PORT int
)

var (
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string
)

var (
	PrivateKeyPath string
	PublicKeyPath  string
	RbacPath       string
)

var (
	AccessTokenDuration  string
	RefreshTokenDuration string
)

var (
	RedisAddr string
)

// InitConfig Load centralized configuration
func InitConfig() {
	v := viper.New()

	AppName := os.Getenv("APP_NAME")
	AppProfile := os.Getenv("APP_PROFILE")
	ConfigURL := os.Getenv("CONFIG_URL")

	if err := core.LoadConfig(AppName, AppProfile, ConfigURL, v); err != nil {
		log.Println("failed to load config:", err)
	}

	rootEnv := "app.environment."

	PORT = v.GetInt("server.port")

	DbHost = v.GetString(rootEnv + "DB_HOST")
	DbPort = v.GetInt(rootEnv + "DB_PORT")
	DbUser = v.GetString(rootEnv + "DB_USER")
	DbPassword = v.GetString(rootEnv + "DB_PASSWORD")
	DbName = v.GetString(rootEnv + "DB_NAME")

	AccessTokenDuration = v.GetString(rootEnv + "ACCESS_TOKEN_DURATION")
	RefreshTokenDuration = v.GetString(rootEnv + "REFRESH_TOKEN_DURATION")

	RedisAddr = v.GetString(rootEnv + "REDIS_ADDR")

	PrivateKeyPath = v.GetString(rootEnv + "PRIVATE_KEY_PATH")
	PublicKeyPath = v.GetString(rootEnv + "PUBLIC_KEY_PATH")
	RbacPath = v.GetString(rootEnv + "RBAC_PATH")

}
