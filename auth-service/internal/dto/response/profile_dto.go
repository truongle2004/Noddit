package response

import (
	"time"
)

type ProfileResponseDTO struct {
	ID          string    `json:"id"`
	UserID      uint      `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DisplayName string    `json:"display_name"`
	Bio         string    `json:"bio"`
	AvatarURL   string    `json:"avatar_url"`
	BirthDate   time.Time `json:"birth_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// func NewProfileResponseDTO(profile *domain.Profile) *ProfileResponseDTO {
// 	return &ProfileResponseDTO{
// 		ID:          profile.ID,
// 		UserID:      profile.UserID,
// 		FirstName:   profile.FirstName,
// 		LastName:    profile.LastName,
// 		DisplayName: profile.DisplayName,
// 		Bio:         profile.Bio,
// 		AvatarURL:   profile.AvatarURL,
// 		BirthDate:   profile.BirthDate,
// 		CreatedAt:   profile.CreatedAt,
// 		UpdatedAt:   profile.UpdatedAt,
// 	}
// }
