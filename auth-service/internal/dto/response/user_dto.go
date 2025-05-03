package response

type RoleDto struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type UserDto struct {
	Id       string    `json:"id,omitempty"`
	Email    string    `json:"email,omitempty"`
	Username string    `json:"username,omitempty"`
	Role     []RoleDto `json:"roles,omitempty"`
}
