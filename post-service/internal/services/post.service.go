package services

import (
	"github.com/gin-gonic/gin"
)

type PostService interface {
	CreateNewPost(c *gin.Context)
	// Delete(id string)
	// Update()
	// Get()
	// GetAll()
}
