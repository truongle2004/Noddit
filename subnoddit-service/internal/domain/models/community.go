package models

import (
	"github.com/truongle2004/service-context/core"
)

type Community struct {
	core.SQLModel
	Name         *string `gorm:"type:varchar(100);unique;not null"`
	Title        *string `gorm:"type:varchar(255)"`
	Description  *string `gorm:"type:text"`
	Rules        []Rule  `gorm:"foreignKey:CommunityID;constraint:OnDelete:CASCADE"`
	Type         *string `gorm:"type:varchar(20);default:public;not null"`
	BannerImage  *string `gorm:"type:text"`
	ProfileImage *string `gorm:"type:text"`
	CreatorID    *string `gorm:"type:varchar(100);not null"`
	MemberCount  *int64  `gorm:"type:int;default:0;not null"`
}

// JSONB = custom type for JSON fields (rules)
type JSONB map[string]interface{}

// GormDataType tells GORM to treat JSONB as jsonb
func (j JSONB) GormDataType() string {
	return "jsonb"
}
