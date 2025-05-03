package impl

import (
	"github.com/casbin/casbin/v2"
)

type CasbinServiceImpl struct {
	enforcer *casbin.Enforcer
}

func NewCasbinService(enforcer *casbin.Enforcer) *CasbinServiceImpl {
	return &CasbinServiceImpl{enforcer: enforcer}
}

func (c *CasbinServiceImpl) CheckPermission(userID, obj, act string) (bool, error) {
	return c.enforcer.Enforce(userID, obj, act)
}

func (c *CasbinServiceImpl) AssignRole(userID, role string) (bool, error) {
	ok, err := c.enforcer.AddGroupingPolicy(userID, role)
	if err != nil {
		return false, err
	}
	err = c.enforcer.SavePolicy()
	return ok, err
}

func (c *CasbinServiceImpl) RemoveRoles(userID string) (bool, error) {
	ok, err := c.enforcer.DeleteRolesForUser(userID)
	if err != nil {
		return false, err
	}
	err = c.enforcer.SavePolicy()
	return ok, err
}

func (c *CasbinServiceImpl) GetRoles(userID string) ([]string, error) {
	return c.enforcer.GetRolesForUser(userID)
}
