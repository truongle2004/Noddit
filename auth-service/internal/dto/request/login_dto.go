package request

import (
	"errors"
)

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginDto) Validate() error {
	if l.Email == "" {
		return errors.New("email is required")
	}

	if l.Password == "" {
		return errors.New("password is required")
	}

	return nil
}
