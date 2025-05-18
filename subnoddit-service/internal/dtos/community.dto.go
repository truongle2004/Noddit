package dtos

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type CommunityDto struct {
	ID            string     `json:"id"`
	Name          *string    `json:"name,omitempty"`
	Description   *string    `json:"description,omitempty"`
	Rules         []RuleDTO  `json:"rules,omitempty"`
	Type          *string    `json:"type,omitempty"`
	BannerImage   *string    `json:"banner_image,omitempty"`
	IconImage     *string    `json:"icon_image,omitempty"`
	CreatorId     *string    `json:"creator_id,omitempty"`
	CreatorName   *string    `json:"creator_name,omitempty"`
	CreatorAvatar *string    `json:"creator_avatar,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	Topics        []TopicDto `json:"topics,omitempty"`
}

func (r *CommunityDto) Validate() error {
	if r.Name == nil || strings.TrimSpace(*r.Name) == "" {
		return errors.New("name is required")
	}
	if len(*r.Name) < 3 || len(*r.Name) > 100 {
		return fmt.Errorf("name must be between 3 and 100 characters (got %d)", len(*r.Name))
	}

	if r.Type != nil {
		validTypes := map[string]bool{
			"public":     true,
			"private":    true,
			"restricted": true,
		}
		if !validTypes[*r.Type] {
			return errors.New("type must be one of: public, private, restricted")
		}
	}

	if r.BannerImage != nil && len(*r.BannerImage) > 255 {
		return fmt.Errorf("banner_image must be less than or equal to 255 characters (got %d)", len(*r.BannerImage))
	}
	if r.IconImage != nil && len(*r.IconImage) > 255 {
		return fmt.Errorf("icon_image must be less than or equal to 255 characters (got %d)", len(*r.IconImage))
	}

	ruleCount := len(r.Rules)
	if ruleCount < 1 || ruleCount > 8 {
		return fmt.Errorf("a community must have at least 1 and at most 8 rules (got %d)", ruleCount)
	}
	for _, rule := range r.Rules {
		if err := rule.Validate(); err != nil {
			return err
		}
	}

	return nil
}
