package config

import (
	"blog-service/internal/domain/models"
	"blog-service/internal/environment"
	"fmt"
	"log"
	"sync"
	"time"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var (
	DbInstance *gorm.DB
	dbOnce     sync.Once
)

// InitDB initializes the database using singleton pattern
func InitDB() error {
	var initErr error

	dbOnce.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
			environment.DbHost, environment.DbPort, environment.DbUsername, environment.DbPassword, environment.DbName,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			initErr = fmt.Errorf("❌ Failed to connect to PostgreSQL: %v", err)
			return
		}

		if err := db.AutoMigrate(&models.Post{}, &models.PostImage{}); err != nil {
			initErr = fmt.Errorf("❌ Failed to auto migrate: %v", err)
			return
		}

		sqlDB, err := db.DB()
		if err != nil {
			initErr = fmt.Errorf("❌ Failed to get underlying DB: %v", err)
			return
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)

		log.Println("✅ Connected to PostgreSQL database!")
		DbInstance = db
	})

	return initErr
}
