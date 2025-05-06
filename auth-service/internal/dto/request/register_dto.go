package request

import (
	"errors"
	"regexp"
)

type RegisterDto struct {
	Username        string `json:"username,omitempty"`
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`
}

func IsEmailValid(e string) bool {
	emailRegex := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(e)
}

func CheckPassword(p string) bool {
	lt := len(p)
	return lt > 8 && lt < 30
}

func (r *RegisterDto) Validate() error {
	if r.Username == "" {
		return errors.New("user name is required")
	}

	if r.Email == "" {
		return errors.New("email is required")
	}

	if !IsEmailValid(r.Email) {
		return errors.New("email must be a valid address like example@domain.com")
	}

	if r.Password == "" {
		return errors.New("password is required")
	}

	if !CheckPassword(r.Password) {
		return errors.New("password must be between 8 and 30 characters")
	}

	if r.ConfirmPassword == "" {
		return errors.New("password confirmation is required")
	}

	if r.Password != r.ConfirmPassword {
		return errors.New("passwords do not match")
	}

	return nil
}
