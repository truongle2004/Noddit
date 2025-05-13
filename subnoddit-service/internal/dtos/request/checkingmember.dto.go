package request

import "fmt"

type IsUserMemberRequest struct {
	CommunityID string `uri:"community_id"`
	UserID      string `uri:"user_id"`
}

func (r *IsUserMemberRequest) Validate() error {
	if r.CommunityID == "" {
		return fmt.Errorf("community_id is required")
	}
	if r.UserID == "" {
		return fmt.Errorf("user_id is required")
	}
	return nil
}
