package config

import (
	"auth-service/internal/environment"
	"fmt"
	"log"
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var (
	RbacEnforcer *casbin.Enforcer
	rbacOnce     sync.Once
)

// InitRBAC initializes Casbin enforcer using singleton pattern
func InitRBAC() error {
	var initErr error

	rbacOnce.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s sslmode=disable",
			environment.DbHost,
			environment.DbPort,
			environment.DbUser,
			environment.DbPassword,
		)

		a, err := gormadapter.NewAdapter("postgres", dsn)
		if err != nil {
			initErr = fmt.Errorf("❌ Failed to create Casbin adapter: %v", err)
			return
		}

		e, err := casbin.NewEnforcer(environment.RbacPath, a)
		if err != nil {
			initErr = fmt.Errorf("❌ Failed to create Casbin enforcer: %v", err)
			return
		}

		if err := e.LoadPolicy(); err != nil {
			initErr = fmt.Errorf("❌ Failed to load Casbin policy: %v", err)
			return
		}

		if err := e.SavePolicy(); err != nil {
			initErr = fmt.Errorf("❌ Failed to save Casbin policy: %v", err)
			return
		}

		log.Println("✅ Loaded Casbin policies successfully")

		RbacEnforcer = e
	})

	return initErr
}
