package impl

import (
	"github.com/casbin/casbin/v2"
)

type CasbinServiceImpl struct {
	enforcer *casbin.Enforcer
}

func NewRBACService(enforcer *casbin.Enforcer) *CasbinServiceImpl {
	return &CasbinServiceImpl{enforcer: enforcer}
}

func (c *CasbinServiceImpl) CheckPermission(userID, obj, act string) (bool, error) {
	return c.enforcer.Enforce(userID, obj, act)
}

func (c *CasbinServiceImpl) GetRoles(userID string) ([]string, error) {
	return c.enforcer.GetRolesForUser(userID)
}
