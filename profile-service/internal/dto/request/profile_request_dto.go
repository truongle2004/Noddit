package request

import (
	"errors"
	"net/url"
	"time"
)

var (
	ErrFirstNameRequired   = errors.New("first name is required")
	ErrLastNameRequired    = errors.New("last name is required")
	ErrDisplayNameRequired = errors.New("display name is required")
	ErrBioTooLong          = errors.New("bio is too long (max 500 characters)")
	ErrInvalidAvatarURL    = errors.New("avatar URL is invalid")
	ErrBirthDateInFuture   = errors.New("birth date cannot be in the future")
)

type ProfileDto struct {
	UserId      string    `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DisplayName string    `json:"display_name"`
	Bio         string    `json:"bio"`
	AvatarURL   string    `json:"avatar_url"`
	BirthDate   time.Time `json:"birth_date"`
}

func (u *ProfileDto) Validate() error {
	if u.FirstName == "" {
		return ErrFirstNameRequired
	}
	if u.LastName == "" {
		return ErrLastNameRequired
	}
	if u.DisplayName == "" {
		return ErrDisplayNameRequired
	}
	if len(u.Bio) > 500 {
		return ErrBioTooLong
	}
	if u.AvatarURL != "" {
		if _, err := url.ParseRequestURI(u.AvatarURL); err != nil {
			return ErrInvalidAvatarURL
		}
	}
	if u.BirthDate.After(time.Now()) {
		return ErrBirthDateInFuture
	}
	return nil
}
