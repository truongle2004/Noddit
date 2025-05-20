package controller

import (
	"blog-service/internal/constant"
	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
	"github.com/truongle2004/service-context/logger"
	"path/filepath"
)

type ImageController struct {
}

func NewImageController() *ImageController {
	return &ImageController{}
}

func (a *ImageController) RegisterRoutes(r *gin.Engine) {
	v1 := r.Group(core.V1 + "/post-service")
	{
		v1.GET("/image/:filename", func(ctx *gin.Context) {
			filename := ctx.Param("filename")
			logger.Info(filename)
			ctx.File(filepath.Join(constant.ImagePath, filename))
		})
	}
}
