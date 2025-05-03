package request

import (
	"testing"
	"time"
)

func TestProfileDto_Validate(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name    string
		dto     ProfileDto
		wantErr error
	}{
		{
			name: "valid profile",
			dto: ProfileDto{
				UserId:      "123",
				FirstName:   "John",
				LastName:    "Doe",
				DisplayName: "johndoe",
				Bio:         "This is a short bio.",
				AvatarURL:   "https://example.com/avatar.jpg",
				BirthDate:   now.AddDate(-20, 0, 0), // 20 years ago
			},
			wantErr: nil,
		},
		{
			name:    "missing first name",
			dto:     ProfileDto{LastName: "Doe", DisplayName: "johndoe"},
			wantErr: ErrFirstNameRequired,
		},
		{
			name:    "missing last name",
			dto:     ProfileDto{FirstName: "John", DisplayName: "johndoe"},
			wantErr: ErrLastNameRequired,
		},
		{
			name:    "missing display name",
			dto:     ProfileDto{FirstName: "John", LastName: "Doe"},
			wantErr: ErrDisplayNameRequired,
		},
		{
			name: "bio too long",
			dto: ProfileDto{
				FirstName:   "John",
				LastName:    "Doe",
				DisplayName: "johndoe",
				Bio:         makeLongString(501),
			},
			wantErr: ErrBioTooLong,
		},
		{
			name: "invalid avatar url",
			dto: ProfileDto{
				FirstName:   "John",
				LastName:    "Doe",
				DisplayName: "johndoe",
				AvatarURL:   "htp:/invalid-url",
			},
			wantErr: ErrInvalidAvatarURL,
		},
		{
			name: "birth date in future",
			dto: ProfileDto{
				FirstName:   "John",
				LastName:    "Doe",
				DisplayName: "johndoe",
				BirthDate:   now.AddDate(1, 0, 0), // 1 year in future
			},
			wantErr: ErrBirthDateInFuture,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.dto.Validate()
			if err != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// helper function to make a string of a specific length
func makeLongString(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "a"
	}
	return s
}
