package services

import "github.com/gin-gonic/gin"

type ImageService interface {
	// UploadImage return banner url and icon url
	UploadImage(c *gin.Context)
	LoadImage(c *gin.Context)
}
