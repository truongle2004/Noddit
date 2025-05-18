package helper

import (
	"errors"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// SaveImageToFolder saves the image to the specified folder
func SaveImageToFolder(ctx *gin.Context, image *multipart.FileHeader, dst string) error {
	if err := ctx.SaveUploadedFile(image, dst); err != nil {
		return err
	}
	return nil
}

// CheckExtension checks if the extension of the image is valid
func CheckExtension(ext string) (string, error) {
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		return "", errors.New("type of image should be jpg, png, jpeg")
	}
	return ext, nil
}

func GenerateDst(path, fileName string) string {
	return filepath.Join(path, fileName)
}
