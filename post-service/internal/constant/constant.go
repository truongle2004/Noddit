package constant

type PostType string

const (
	PostTypeText  PostType = "text"
	PostTypeImage PostType = "image"
	PostTypeLink  PostType = "link"
	PostTypeVideo PostType = "video"
)
