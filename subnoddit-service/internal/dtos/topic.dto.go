package dtos

type TopicDto struct {
	ID          string `json:"id"`
	Name        string `json:"name" binding:"required,alphanum,min=3,max=100"`
	Description string `json:"description"`
}
