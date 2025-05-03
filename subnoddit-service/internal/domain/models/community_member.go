package models

import (
	"time"
)

type CommunityMember struct {
	UserID      string    `gorm:"primaryKey"`
	CommunityID string    `gorm:"type:uuid;primaryKey"`
	JoinedAt    time.Time `gorm:"autoCreateTime"`
}
