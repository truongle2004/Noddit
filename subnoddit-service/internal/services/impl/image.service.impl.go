package impl

import (
	"fmt"
	"net/http"
	"path/filepath"
	"subnoddit-service/internal/constant"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
)

type ImageServiceImpl struct{}

func NewImageService() *ImageServiceImpl {
	return &ImageServiceImpl{}
}

func (u *ImageServiceImpl) UploadImage(ctx *gin.Context) {
	communityId := ctx.GetHeader(string(constant.XCommunityID))
	fields := []string{"banner", "icon"}

	uploaded := map[string]string{}

	for _, field := range fields {
		file, err := ctx.FormFile(field)
		if err != nil {
			if err == http.ErrMissingFile {
				continue
			}
			ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError("Failed to retrieve file"))
			return
		}

		ext := filepath.Ext(file.Filename)
		if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
			ctx.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError("Type of image should be jpg, png, jpeg"))
			return
		}

		fileName := fmt.Sprintf("%s_%s%s", field, communityId, ext)
		dst := filepath.Join(constant.UploadImagePath, fileName)
		if err := ctx.SaveUploadedFile(file, dst); err != nil {
			ctx.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithDetail("error", err.Error()))
			return
		}

		// Save relative or public path for return
		uploaded[field+"_url"] = fileName
	}

	ctx.JSON(http.StatusOK, uploaded)
}

func (u *ImageServiceImpl) LoadImage(ctx *gin.Context) {
	filename := ctx.Param("filename")
	ctx.File(filepath.Join(constant.UploadImagePath, filename))
}
