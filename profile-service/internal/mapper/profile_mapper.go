package mapper

import (
	"profile-service/internal/domain"
	"profile-service/internal/dto/request"
	"profile-service/internal/dto/response"
)

func ToProfileDto(profile *domain.Profile) *request.ProfileDto {
	if profile == nil {
		return nil
	}
	return &request.ProfileDto{
		UserId:      profile.UserID,
		FirstName:   profile.FirstName,
		LastName:    profile.LastName,
		DisplayName: profile.DisplayName,
		Bio:         profile.Bio,
		AvatarURL:   profile.AvatarURL,
		BirthDate:   profile.BirthDate,
	}
}

func ToProfileEntity(profileDto *request.ProfileDto) *domain.Profile {
	if profileDto == nil {
		return nil
	}
	return &domain.Profile{
		UserID:      profileDto.UserId,
		FirstName:   profileDto.FirstName,
		LastName:    profileDto.LastName,
		DisplayName: profileDto.DisplayName,
		Bio:         profileDto.Bio,
		AvatarURL:   profileDto.AvatarURL,
		BirthDate:   profileDto.BirthDate,
	}
}

func ToProfileDtoResponse(profile *domain.Profile) *response.ProfileResponseDto {
	if profile == nil {
		return nil
	}
	return &response.ProfileResponseDto{
		UserId:      profile.UserID,
		DisplayName: profile.DisplayName,
		AvatarUrl:   profile.AvatarURL,
	}
}
