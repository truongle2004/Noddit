package enums

type EnumPostStatus string

const (
	PostStatusDraft    EnumPostStatus = "draft"
	PostStatusPublish  EnumPostStatus = "publish"
	PostStatusDeleted  EnumPostStatus = "deleted"
	PostStatusArchived EnumPostStatus = "archived"
)
