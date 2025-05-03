package rest

import (
	"profile-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/truongle2004/service-context/core"
)

type ProfileController struct {
	profileSvc services.ProfileService
}

func NewProfileController(profileSvc services.ProfileService) *ProfileController {
	return &ProfileController{
		profileSvc: profileSvc,
	}
}

func (ctrl *ProfileController) RegisterRoute(r *gin.Engine) {
	v1 := r.Group(core.V1 + "/profile-service")
	{
		v1.POST("/update", ctrl.UpdateProfile)
		v1.GET("/get", ctrl.GetProfile)

	}
}

func (ctrl *ProfileController) UpdateProfile(c *gin.Context) {
	ctrl.profileSvc.UpdateProfile(c)
}

func (ctrl *ProfileController) GetProfile(c *gin.Context) {
	ctrl.profileSvc.GetProfile(c)
}
