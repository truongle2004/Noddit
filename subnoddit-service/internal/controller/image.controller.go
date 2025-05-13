package controller

import (
	"subnoddit-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
)

type ImageController struct {
	imageSvc services.ImageService
}

func NewImageController(imageSvc services.ImageService) ImageController {
	return ImageController{imageSvc: imageSvc}
}

func (i *ImageController) RegisterRoutes(r *gin.Engine) {
	api := r.Group(core.V1 + "/subnoddit-service")

	image := api.Group("/image")
	{
		image.POST("/upload", i.imageSvc.UploadImage)
		image.GET("/:filename", i.imageSvc.LoadImage)
	}
}
