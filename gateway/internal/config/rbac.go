package config

import (
	"fmt"
	"gateway/internal/constant"
	"gateway/internal/environment"
	"log"
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

const (
	RoleIndex       = 0
	UrlIndex        = 1
	MethodIndex     = 2
	ParentRoleIndex = 0
	ChildRoleIndex  = 1
)

var (
	EnforcerInstance *casbin.Enforcer
	enforcerOnce     sync.Once
	enforcerErr      error
)

func InitRBAC() error {
	enforcerOnce.Do(func() {
		var a *gormadapter.Adapter
		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s sslmode=disable",
			environment.DbHost,
			environment.DbPort,
			environment.DbUser,
			environment.DbPassword,
		)

		a, err := gormadapter.NewAdapter("postgres", dsn)
		if err != nil {
			return
		}

		EnforcerInstance, enforcerErr = casbin.NewEnforcer(environment.RbacPath, a)
		if enforcerErr != nil {
			return
		}

		if enforcerErr = EnforcerInstance.LoadPolicy(); enforcerErr != nil {
			return
		}

		// Define casbin permissions
		policies := [][]string{
			{"ROLE_USER", "/api/posts", "POST"},
			{"ROLE_USER", "/api/posts/:id/vote", "POST"},
			{"ROLE_VERIFIED", "/api/posts/:id/edit", "PUT"},
			{"ROLE_MODERATOR", "/api/subreddits/:id/moderate", "POST"},
			{"ROLE_CREATOR", "/api/subreddits", "POST"},
			{"ROLE_ADMIN", constant.V1 + "/users/", "GET"},
			{"ROLE_ADMIN", "/users/:id/ban", "POST"},
		}

		for _, p := range policies {
			if _, enforcerErr = EnforcerInstance.AddPolicy(p[RoleIndex], p[UrlIndex], p[MethodIndex]); enforcerErr != nil {
				return
			}
		}

		// Role inheritance
		groupings := [][]string{
			{"ROLE_VERIFIED", "ROLE_USER"},
			{"ROLE_MODERATOR", "ROLE_VERIFIED"},
			{"ROLE_CREATOR", "ROLE_MODERATOR"},
			{"ROLE_ADMIN", "ROLE_CREATOR"},
		}

		for _, g := range groupings {
			if len(g) != 2 {
				enforcerErr = fmt.Errorf("invalid grouping policy: %v", g)
				return
			}
			if _, enforcerErr = EnforcerInstance.AddGroupingPolicy(g[ParentRoleIndex], g[ChildRoleIndex]); enforcerErr != nil {
				return
			}
		}

		if enforcerErr = EnforcerInstance.SavePolicy(); enforcerErr != nil {
			return
		}

		log.Println("✅ Load casbin successfully")
	})

	return enforcerErr
}
