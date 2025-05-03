package response

type ProfileResponseDto struct {
	UserId      string `json:"user_id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	AvatarUrl   string `json:"avatar_url,omitempty"`
}
