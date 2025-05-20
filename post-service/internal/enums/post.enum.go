package enums

type EnumPostType string

const (
	PostTypeText  EnumPostType = "text"
	PostTypeImage EnumPostType = "image"
	PostTypeLink  EnumPostType = "link"
	PostTypeVideo EnumPostType = "video"
)

func GetEnumPostType(postType *string) EnumPostType {
	if postType == nil {
		return PostTypeText
	}

	switch EnumPostType(*postType) {
	case PostTypeText:
		return PostTypeText
	case PostTypeImage:
		return PostTypeImage
	case PostTypeLink:
		return PostTypeLink
	case PostTypeVideo:
		return PostTypeVideo
	default:
		return ""
	}
}
