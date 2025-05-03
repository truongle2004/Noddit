package models

import (
	"blog-service/internal/constant"

	"github.com/truongle2004/service-context/core"
)

type Post struct {
	core.SQLModel
	Title        *string            `gorm:"type:varchar(255);not null"`
	Content      *string            `gorm:"type:text"`
	AuthorID     *string            `gorm:"type:varchar(255);index"`
	CommunityID  *string            `gorm:"type:varchar(255);index"`
	Tags         *string            `gorm:"type:varchar(255)"`
	PostType     *constant.PostType `gorm:"type:varchar(52);not null"`
	MediaURL     *string            `gorm:"type:text"`
	CommentCount *int               `gorm:"default:0"`
	IsEdited     *bool              `gorm:"default:false"`
	IsDeleted    *bool              `gorm:"default:false"`
}
