package request

import "errors"

type AccountStatusUpdateRequest struct {
	Status string `json:"status"`
	ID     string `json:"id"`
	Email  string `json:"email"`
}

func (req *AccountStatusUpdateRequest) Validate() error {
	if req.Status == "" {
		return errors.New("status is required")
	}
	if req.ID == "" {
		return errors.New("id is required")
	}
	if req.Email == "" {
		return errors.New("email is required")
	}
	return nil
}
