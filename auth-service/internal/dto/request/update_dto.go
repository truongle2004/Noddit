package request

type UpdateUserProfileDto struct {
	FirstName   *string `json:"first_name" validate:"omitempty,max=50"`
	LastName    *string `json:"last_name" validate:"omitempty,max=50"`
	DisplayName *string `json:"display_name" validate:"omitempty,max=50"`
	Bio         *string `json:"bio" validate:"omitempty,max=255"`
	AvatarURL   *string `json:"avatar_url" validate:"omitempty,url"`
	BirthDate   *string `json:"birth_date" validate:"omitempty,datetime=2006-01-02"`
}
