package dtos

import "time"

type PostDto struct {
	ID          *string    `json:"id,omitempty"`
	Title       *string    `json:"title"`
	Slug        *string    `json:"slug,omitempty"`
	Content     *string    `json:"content,omitempty"`
	CreatorId   *string    `json:"creator_id"`
	CreatorName *string    `json:"creator_name"`
	CommunityID *string    `json:"community_id"`
	Tags        *string    `json:"tags,omitempty"`
	PostType    *string    `json:"post_type"`
	ReadTime    *int       `json:"read_time,omitempty"`
	MetaTitle   *string    `json:"meta_title,omitempty"`
	MetaDesc    *string    `json:"meta_desc,omitempty"`
	ImageUrls   *[]string  `json:"image_urls,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
}
