package models

import (
	"time"

	"github.com/truongle2004/service-context/core"
)

type Post struct {
	core.SQLModel
	Title         *string    `gorm:"type:varchar(255);not null"`
	Slug          *string    `gorm:"type:varchar(255);uniqueIndex"`
	Content       *string    `gorm:"type:text"`
	CreatorID     *string    `gorm:"type:varchar(255);index"`
	CommunityID   *string    `gorm:"type:varchar(255);index"`
	Tags          *string    `gorm:"type:varchar(255)"`
	PostType      *string    `gorm:"type:varchar(52);not null"`
	ViewCount     *int       `gorm:"default:0"`
	LikeCount     *int       `gorm:"default:0"`
	CommentCount  *int       `gorm:"default:0"`
	ShareCount    *int       `gorm:"default:0"`
	ReadTime      *int       `gorm:"default:0"`
	Status        *string    `gorm:"type:varchar(50);default:'draft'"`
	IsEdited      *bool      `gorm:"default:false"`
	IsDeleted     *bool      `gorm:"default:false"`
	IsFeatured    *bool      `gorm:"default:false"`
	FeaturedOrder *int       `gorm:"default:0"`
	PublishedAt   *time.Time `gorm:"index"`
	MetaTitle     *string    `gorm:"type:varchar(255)"`
	MetaDesc      *string    `gorm:"type:varchar(255)"`
}
