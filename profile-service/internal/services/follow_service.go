package services

import "github.com/gin-gonic/gin"

type FollowService interface {
	GetFollowers(c *gin.Context)
	Follow(c *gin.Context)
	UnFollow(c *gin.Context)
}
