package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/truongle2004/service-context/core"
	"github.com/truongle2004/service-context/redisclient"
	"log"
	"net/http"
)

func AccountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		email, _ := c.Get("email")
		log.Println("email: ", email)

		accountStatus, err := redisclient.Get(c, email.(string))
		if errors.Is(err, redis.Nil) {
			c.AbortWithStatusJSON(http.StatusForbidden, core.ErrForbidden.WithError("Account user not found. Please contact admin or try again later"))
			return
		}

		// The account status should be active
		if accountStatus != string(core.ACTIVE) {
			c.AbortWithStatusJSON(http.StatusForbidden, core.ErrForbidden.WithError("Account user is not active. Please contact admin or try again later"))
			return
		}

		c.Next()
	}
}
