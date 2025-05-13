package server

import (
	"auth-service/internal/config"
	"auth-service/internal/controller"
	"auth-service/internal/helper"
	repository "auth-service/internal/repositories/impl"
	service "auth-service/internal/services/impl"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(middleware.ResponseFormatterMiddleware())
	r.Use(cors.New(config.CorsConfig()))

	config.InitRBAC()
	helper.InitJwtHelper()

	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := config.InitRedis(); err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}

	if err := config.InitRBAC(); err != nil {
		log.Fatalf("Failed to initialize RBAC: %v", err)
	}

	redisSvc := service.NewRedisService(config.RedisClient)
	casbinSvc := service.NewCasbinService(config.RbacEnforcer)
	userRepo := repository.NewUserRepository(config.DbInstance)
	authSvc := service.NewAuthService(userRepo, redisSvc, casbinSvc)
	userSvc := service.NewUserService(userRepo, redisSvc)

	authController := controller.NewAuthController(authSvc)
	userController := controller.NewUserController(userSvc)

	authController.RegisterRoutes(r)
	userController.RegisterRoutes(r)

	return r
}
