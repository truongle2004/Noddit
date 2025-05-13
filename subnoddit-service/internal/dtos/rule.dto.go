package dtos

import "fmt"

type RuleDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Position    int16  `json:"position"`
}

func (r *RuleDTO) Validate() error {
	if r.Title == "" {
		return fmt.Errorf("title is required")
	}
	if len(r.Title) > 255 {
		return fmt.Errorf("title must be less than or equal to 255 characters (got %d)", len(r.Title))
	}
	if r.Description != "" && len(r.Description) > 1000 {
		return fmt.Errorf("description must be less than or equal to 1000 characters (got %d)", len(r.Description))
	}
	return nil
}
