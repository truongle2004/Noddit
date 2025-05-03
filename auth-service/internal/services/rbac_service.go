package services

type CasbinService interface {
	CheckPermission(userID, obj, act string) (bool, error)
	AssignRole(userID, role string) (bool, error)
	RemoveRoles(userID string) (bool, error)
	GetRoles(userID string) ([]string, error)
}
