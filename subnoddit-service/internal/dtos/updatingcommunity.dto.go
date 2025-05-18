package dtos

import (
	"fmt"
	"strings"
)

type UpdateCommunityRequest struct {
	ID           string                  `json:"id"`
	Name         *string                 `json:"name,omitempty" binding:"omitempty,alphanum,min=3,max=100"`
	Title        *string                 `json:"title,omitempty" binding:"omitempty,max=255"`
	Description  *string                 `json:"description,omitempty"`
	Rules        *map[string]interface{} `json:"rules,omitempty"`
	Type         *string                 `json:"type,omitempty" binding:"omitempty,oneof=public private restricted"`
	BannerImage  *string                 `json:"banner_image,omitempty"`
	ProfileImage *string                 `json:"profile_image,omitempty"`
}

func (r *UpdateCommunityRequest) Validate() error {
	var validationErrors []string

	if strings.TrimSpace(r.ID) == "" {
		validationErrors = append(validationErrors, "id is required")
	}

	if r.Name != nil {
		if len(*r.Name) < 3 || len(*r.Name) > 100 {
			validationErrors = append(validationErrors, "name must be between 3 and 100 characters")
		}
	}

	if r.Title != nil && len(*r.Title) > 255 {
		validationErrors = append(validationErrors, "title must be less than 255 characters")
	}

	if r.Type != nil && (*r.Type != "public" && *r.Type != "private" && *r.Type != "restricted") {
		validationErrors = append(validationErrors, "type must be one of: public, private, restricted")
	}

	if r.BannerImage != nil && len(*r.BannerImage) > 255 {
		validationErrors = append(validationErrors, "banner_image must be less than 255 characters")
	}

	if len(validationErrors) > 0 {
		return fmt.Errorf("validation error(s): %s", strings.Join(validationErrors, "; "))
	}

	return nil
}
