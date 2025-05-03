package services

type RBACService interface {
	CheckPermission(userID, obj, act string) (bool, error)
	GetRoles(userID string) ([]string, error)
}
