package models

import "github.com/truongle2004/service-context/core"

type Topic struct {
	core.SQLModel
	Name        string       `gorm:"type:varchar(100);not null;unique"`
	Description string       `gorm:"type:text"`
	Communities []*Community `gorm:"many2many:community_topics;"`
}
