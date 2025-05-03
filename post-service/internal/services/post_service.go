package services

import (
	"blog-service/internal/dto/request"
	"github.com/gin-gonic/gin"
)

type PostService interface {
	Create(c *gin.Context, postDto *request.PostCreateDTO)
	// Delete(id string)
	// Update()
	// Get()
	// GetAll()
}
