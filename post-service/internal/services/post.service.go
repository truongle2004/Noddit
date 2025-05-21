package services

import (
	"github.com/gin-gonic/gin"
)

type PostService interface {
	CreateNewPost(c *gin.Context)
	GetAllPostsByCommunityId(c *gin.Context)
}
