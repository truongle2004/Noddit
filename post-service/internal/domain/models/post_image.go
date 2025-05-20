package models

import "github.com/truongle2004/service-context/core"

type PostImage struct {
	core.SQLModel
	PostID   *string `gorm:"type:varchar(255);index"`
	ImageURL *string `gorm:"type:text"`
}
