package request

import (
	"blog-service/internal/constant"
	"errors"
)

type PostCreateDTO struct {
	Title       *string            `json:"title" binding:"required"`
	Content     *string            `json:"content,omitempty"`
	AuthorID    *string            `json:"author_id" binding:"required"`
	CommunityID *string            `json:"community_id" binding:"required"`
	Tags        *string            `json:"tags,omitempty"`
	PostType    *constant.PostType `json:"post_type" binding:"required"`
	MediaURL    *string            `json:"media_url,omitempty"`
}

func (dto *PostCreateDTO) Validate() error {
	if dto.Title == nil || *dto.Title == "" {
		return errors.New("title is required")
	}
	if dto.AuthorID == nil || *dto.AuthorID == "" {
		return errors.New("author_id is required")
	}
	if dto.CommunityID == nil || *dto.CommunityID == "" {
		return errors.New("community_id is required")
	}
	if dto.PostType == nil {
		return errors.New("post_type is required")
	}
	return nil
}
