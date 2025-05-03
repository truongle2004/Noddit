package request

import (
	"fmt"
	"strings"
)

type JoinCommunityRequest struct {
	CommunityID string `json:"community_id"`
	UserID      string `json:"user_id"`
}

func (r *JoinCommunityRequest) Validate() error {
	var validationErrors []string

	if strings.TrimSpace(r.CommunityID) == "" {
		validationErrors = append(validationErrors, "community_id is required")
	}
	if strings.TrimSpace(r.UserID) == "" {
		validationErrors = append(validationErrors, "user_id is required")
	}

	if len(validationErrors) > 0 {
		return fmt.Errorf("validation error(s): %s", strings.Join(validationErrors, "; "))
	}

	return nil
}
