package controller

import (
	"myadmin/app/service"
	"myadmin/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PrivilegeController struct {
	service service.PrivilegeService
}

func NewPrivilegeController(s service.PrivilegeService) PrivilegeController {
	return PrivilegeController{
		service: s,
	}
}

// @Summary		Get list Privileges
// @Description	REST API for Privileges
// @Tags		Privilege
// @Accept		json
// @Produce		json
// @Security    BearerAuth
// @Success		200	{object}	helper.Response
// @Failure		400	{object}	helper.Response
// @Failure		404	{object}	helper.Response
// @Failure		500	{object}	helper.Response
// @Router		/privileges [get]
func (controller PrivilegeController) Index(ctx *gin.Context) {
	privileges, err := controller.service.List()

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Privilege not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Privilege Found", http.StatusOK, privileges)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Get one Privileges
// @Description	REST API for Privileges
// @Tags		Privilege
// @Accept		json
// @Produce		json
// @Security    BearerAuth
// @Param		id	path		string	false	"Privileges ID"
// @Success		200	{object}	helper.Response
// @Failure		400	{object}	helper.Response
// @Failure		404	{object}	helper.Response
// @Failure		500	{object}	helper.Response
// @Router		/privileges/{id} [get]
func (controller PrivilegeController) Show(ctx *gin.Context) {
	ID := ctx.Param("id") // Get Param ID

	privilege, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Privilege not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.ErrorJSON(ctx, "Privilege Found", http.StatusOK, privilege)

	ctx.JSON(http.StatusOK, resp)
}
