package middleware

import (
	"gateway/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
)

type AccountMiddleware struct {
	redis services.RedisService
}

func NewAccountMiddleware(redis services.RedisService) *AccountMiddleware {
	return &AccountMiddleware{
		redis: redis,
	}
}

func (am *AccountMiddleware) AccountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		email, _ := c.Get("email")

		accountStatus, err := am.redis.Get(c, email.(string))
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		// The account status should be active
		if accountStatus != string(core.ACTIVE) {
			c.AbortWithStatusJSON(403, gin.H{"error": "Your account is not active or blocked by admin, please contact admin for more information"})
			return
		}

		c.Next()
	}
}
