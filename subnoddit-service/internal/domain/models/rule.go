package models

import (
	"github.com/truongle2004/service-context/core"
)

type Rule struct {
	core.SQLModel
	CommunityID *string `gorm:"type:varchar(100);not null;index"`
	Title       *string `gorm:"type:varchar(255);not null"`
	Description *string `gorm:"type:text"`
	Position    *int16  `gorm:"type:smallint;default:0"`
}
