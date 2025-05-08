package dtos

import (
	"errors"
	"fmt"
	"time"
)

type CommunityDto struct {
	Name         string    `json:"name" binding:"required,alphanum,min=3,max=100"`
	Title        string    `json:"title" binding:"omitempty,max=255"`
	Description  string    `json:"description"`
	Rules        []RuleDTO `json:"rules"`
	Type         string    `json:"type" binding:"omitempty,oneof=public private restricted"`
	BannerImage  string    `json:"banner_image"`
	ProfileImage string    `json:"profile_image"`
	CreatorId    string    `json:"creator_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	// TODO: add topic dto
}

func (r *CommunityDto) Validate() error {
	if r.Name == "" {
		return errors.New("name is required")
	}
	if len(r.Name) < 3 || len(r.Name) > 100 {
		return fmt.Errorf("name must be between 3 and 100 characters (got %d)", len(r.Name))
	}
	if r.Title != "" && len(r.Title) > 255 {
		return fmt.Errorf("title must be less than or equal to 255 characters (got %d)", len(r.Title))
	}
	if r.Type != "" && r.Type != "public" && r.Type != "private" && r.Type != "restricted" {
		return errors.New("type must be one of: public, private, restricted")
	}
	if r.BannerImage != "" && len(r.BannerImage) > 255 {
		return fmt.Errorf("banner_image must be less than or equal to 255 characters (got %d)", len(r.BannerImage))
	}
	if r.ProfileImage != "" && len(r.ProfileImage) > 255 {
		return fmt.Errorf("profile_image must be less than or equal to 255 characters (got %d)", len(r.ProfileImage))
	}

	if len(r.Rules) > 0 && len(r.Rules) < 9 {
		for _, rule := range r.Rules {
			if err := rule.Validate(); err != nil {
				return err
			}
		}
	} else {
		return fmt.Errorf("A community must have at least 1 and at most 8 rules")
	}
	return nil
}
