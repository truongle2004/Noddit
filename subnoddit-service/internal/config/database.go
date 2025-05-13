package config

import (
	"fmt"
	"log"
	"subnoddit-service/internal/domain/models"
	"subnoddit-service/internal/environment"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	CommunityDescription = `
			COMMENT ON TABLE communities IS 'Table for subreddit-style communities';
			COMMENT ON COLUMN communities.name IS 'Unique URL slug (e.g., "golang", "movies")';
			COMMENT ON COLUMN communities.description IS 'Longer description about the community';
			COMMENT ON COLUMN communities.rules IS 'Optional posting rules stored as JSONB';
			COMMENT ON COLUMN communities.member_count IS 'The number of members in the community.';
		`
	CommunityMemberDescription = `
			COMMENT ON TABLE community_members IS 'Table that records user memberships (follows) of communities';
			COMMENT ON COLUMN community_members.joined_at IS 'Timestamp when the user joined the community';
		`

	CommunityRule = `
			COMMENT ON TABLE rules IS 'Table that records rules for communities';
			COMMENT ON COLUMN rules.title IS 'Rule title';
			COMMENT ON COLUMN rules.description IS 'Rule description';
			COMMENT ON COLUMN rules.position IS 'Rule position';
	`

	TopicDescription = `
			COMMENT ON TABLE topics IS 'Table that stores content tags or categories used to classify posts';
			COMMENT ON COLUMN topics.name IS 'Unique name identifying the topic (e.g., "programming", "science")';
			COMMENT ON COLUMN topics.description IS 'Optional longer description for the topic';
`
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

		if err := db.AutoMigrate(&models.Community{}, &models.CommunityMember{}, &models.Rule{}, &models.Topic{}); err != nil {
			initErr = fmt.Errorf("❌ Failed to auto migrate: %v", err)
			return
		}

		db.Exec(CommunityDescription)
		db.Exec(CommunityMemberDescription)
		db.Exec(CommunityRule)
		db.Exec(TopicDescription)

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
