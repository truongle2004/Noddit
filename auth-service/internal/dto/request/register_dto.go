package request

import (
	"errors"
	"regexp"
)

type RegisterDto struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

func CheckPassword(p string) bool {
	lt := len(p)
	return lt > 8 && lt < 30
}

func (r *RegisterDto) Validate() error {
	if r.Username == "" {
		return errors.New("User name is required")
	}

	if r.Email == "" {
		return errors.New("Email is required")
	}

	if IsEmailValid(r.Email) {
		return errors.New("Email must be a valid address like example@domain.com")
	}

	if r.Password == "" {
		return errors.New("Password is required")
	}

	if CheckPassword(r.Password) {
		return errors.New("Password must be between 8 and 30 characters")
	}

	if r.ConfirmPassword == "" {
		return errors.New("Password confirmation is required")
	}

	if r.Password != r.ConfirmPassword {
		return errors.New("Passwords do not match")
	}

	return nil
}
