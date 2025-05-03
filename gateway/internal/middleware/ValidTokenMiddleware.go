package middleware

import (
	"gateway/internal/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
)

func ValidTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, _ := c.Get("token")

		publicKey, _ := helper.LoadPublicKey()

		claims, err := helper.ValidateToken(token.(string), publicKey)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		exp := (*claims)["exp"].(float64)
		iat := (*claims)["iat"].(float64)
		if iat > exp {
			c.JSON(http.StatusUnauthorized, core.ErrUnauthorized.WithError("Token is expired"))
			c.Abort()
			return
		}

		// role := (*claims)["role"]
		email := (*claims)["email"]
		userId := (*claims)["id"]

		// storing role and email is to next middleware reuse it
		// c.Set("role", role)
		c.Set("user_id", userId)
		c.Set("email", email)

		c.Next()
	}
}
