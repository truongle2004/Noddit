package domain

import (
	"time"

	"github.com/truongle2004/service-context/core"
)

type User struct {
	core.SQLModel
	Username  string          `gorm:"type:varchar(100);not null;unique"`
	Email     string          `gorm:"type:varchar(100);uniqueIndex"`
	Salt      string          `gorm:"type:varchar(255)"`
	Password  string          `gorm:"type:varchar(255)"`
	Status    core.UserStatus `gorm:"type:user_status;default:'ACTIVE'"`
	LastLogin time.Time       `gorm:"type:timestamp"`
}

func (user *User) IsAccountActive() bool {
	return user.Status == core.ACTIVE
}

