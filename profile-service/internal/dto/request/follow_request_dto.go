package request

import "errors"

type FollowRequestDto struct {
	FollowerId string `json:"follower_id"`
	FolloweeId string `json:"followee_id"`
}

func (f *FollowRequestDto) ValidateFollowRequest() error {
	if f.FollowerId == "" {
		return errors.New("follower_id is required")
	}
	if f.FolloweeId == "" {
		return errors.New("followee_id is required")
	}
	return nil
}
