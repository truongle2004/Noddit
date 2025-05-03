package server

import (
	"auth-service/internal/config"
	"auth-service/internal/controller"
	repository "auth-service/internal/repositories/impl"
	service "auth-service/internal/services/impl"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func LoadConfig(ctx *gin.Context) {
	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := config.InitRedis(ctx); err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}

	if err := config.InitRBAC(); err != nil {
		log.Fatalf("Failed to initialize RBAC: %v", err)
	}
}

func initServerServices(r *gin.Engine) {

	redisSvc := service.NewRedisService(config.RedisClient)
	casbinSvc := service.NewCasbinService(config.RbacEnforcer)
	userRepo := repository.NewUserRepository(config.DbInstance)
	authSvc := service.NewAuthService(userRepo, redisSvc, casbinSvc)
	userSvc := service.NewUserService(userRepo, redisSvc)

	authController := controller.NewAuthController(authSvc)
	userController := controller.NewUserController(userSvc)

	authController.RegisterRoutes(r)
	userController.RegisterRoutes(r)

}

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(config.CorsConfig()))

	r.Use(func(ctx *gin.Context) {
		LoadConfig(ctx)
		ctx.Next()
	})

	initServerServices(r)
	return r
}
