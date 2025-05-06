package request

import (
	"testing"
)

func TestRegisterDto_Validate(t *testing.T) {
	tests := []struct {
		name    string
		dto     RegisterDto
		wantErr bool
		errMsg  string
	}{
		{
			name:    "missing username",
			dto:     RegisterDto{Email: "test@example.com", Password: "password123", ConfirmPassword: "password123"},
			wantErr: true,
			errMsg:  "user name is required",
		},
		{
			name:    "missing email",
			dto:     RegisterDto{Username: "john", Password: "password123", ConfirmPassword: "password123"},
			wantErr: true,
			errMsg:  "email is required",
		},
		{
			name:    "invalid email format",
			dto:     RegisterDto{Username: "john", Email: "bademail", Password: "password123", ConfirmPassword: "password123"},
			wantErr: true,
			errMsg:  "email must be a valid address like example@domain.com",
		},
		{
			name:    "missing password",
			dto:     RegisterDto{Username: "john", Email: "test@example.com", ConfirmPassword: "password123"},
			wantErr: true,
			errMsg:  "password is required",
		},
		{
			name:    "short password",
			dto:     RegisterDto{Username: "john", Email: "test@example.com", Password: "short", ConfirmPassword: "short"},
			wantErr: true,
			errMsg:  "password must be between 8 and 30 characters",
		},
		{
			name:    "missing confirm password",
			dto:     RegisterDto{Username: "john", Email: "test@example.com", Password: "password123"},
			wantErr: true,
			errMsg:  "password confirmation is required",
		},
		{
			name:    "password mismatch",
			dto:     RegisterDto{Username: "john", Email: "test@example.com", Password: "password123", ConfirmPassword: "pass123"},
			wantErr: true,
			errMsg:  "passwords do not match",
		},
		{
			name:    "valid registration",
			dto:     RegisterDto{Username: "john", Email: "test@example.com", Password: "password123", ConfirmPassword: "password123"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.dto.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr = %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.errMsg {
				t.Errorf("expected error message: %q, got: %q", tt.errMsg, err.Error())
			}
		})
	}
}

func TestIsEmailValid(t *testing.T) {
	validEmails := []string{
		"user@example.com",
		"user.name+tag+sorting@example.com",
		"admin@mail.co.uk",
	}

	invalidEmails := []string{
		"plainaddress",
		"@missingusername.com",
		"user@.com",
	}

	for _, email := range validEmails {
		if !IsEmailValid(email) {
			t.Errorf("IsEmailValid(%q) = false, want true", email)
		}
	}

	for _, email := range invalidEmails {
		if IsEmailValid(email) {
			t.Errorf("IsEmailValid(%q) = true, want false", email)
		}
	}
}

func TestCheckPassword(t *testing.T) {
	tests := []struct {
		password string
		valid    bool
	}{
		{"12345678", false},
		{"123456789", true},
		{"thisisaverylongpasswordthatisinvalidbecauseitistoolong", false},
		{"validPass123", true},
	}

	for _, tt := range tests {
		if got := CheckPassword(tt.password); got != tt.valid {
			t.Errorf("CheckPassword(%q) = %v, want %v", tt.password, got, tt.valid)
		}
	}
}
