package domain

import (
	"time"

	"github.com/truongle2004/service-context/core"
)

type Profile struct {
	core.SQLModel
	UserID      string    `gorm:"type:bigint unsigned;unique;not null"`
	FirstName   string    `gorm:"type:varchar(50)"`
	LastName    string    `gorm:"type:varchar(50)"`
	DisplayName string    `gorm:"type:varchar(100)"`
	Bio         string    `gorm:"type:text"`
	AvatarURL   string    `gorm:"type:varchar(255)"`
	BirthDate   time.Time `gorm:"type:date"`
}
