package middleware

import (
	"gateway/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
)

func RBACMiddleware(rbacSvc services.RBACService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, _ := c.Get("user_id")
		obj := c.Request.URL.Path
		act := c.Request.Method

		if userId == nil {
			c.JSON(http.StatusUnauthorized, core.ErrUnauthorized.WithReason("User id not found"))
			return
		}

		isAllowerd, err := rbacSvc.CheckPermission(userId.(string), obj, act)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		if !isAllowerd {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
