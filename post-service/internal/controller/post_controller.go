package controller

import (
	"blog-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
)

type PostController struct {
	PostSvc services.PostService
}

func NewCommunityController(postSvc services.PostService) PostController {
	return PostController{PostSvc: postSvc}
}

func (a *PostController) RegisterRoutes(r *gin.Engine) {
	v1 := r.Group(core.V1 + "/post")
	{
		v1.POST("/create", a.PostSvc.Create)
	}
}
