package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
)

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exist := c.Get("role")
		if !exist {
			c.JSON(http.StatusForbidden, core.ErrForbidden.
				WithReason("Missing role").
				WithError("You don't have permission to access this resource"))
			c.Abort()
			return
		}

		roleStr, ok := userRole.(string)
		if !ok {
			c.JSON(http.StatusForbidden, core.ErrForbidden.
				WithReason("Invalid role").
				WithError("You don't have permission to access this resource"))
			c.Abort()
			return
		}

		for _, role := range allowedRoles {
			if role == roleStr {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, core.ErrForbidden.
			WithReason("You don't have permission to access this resource").
			WithError("You don't have permission to access this resource"))
		c.Abort()
	}
}
