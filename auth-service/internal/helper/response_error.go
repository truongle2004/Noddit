package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
)

func ResponseServerError(c *gin.Context, reason string, err error) {
	c.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", err.Error()).WithReason(reason))
}
