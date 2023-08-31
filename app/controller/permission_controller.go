package controller

import (
	"myadmin/app/request"
	"myadmin/app/service"
	"myadmin/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PermissionController struct {
	service service.PermissionService
}

func NewPermissionController(s service.PermissionService) PermissionController {
	return PermissionController{
		service: s,
	}
}

// @Summary		Get list permission
// @Description	REST API for permission
// @Tags		Permissions
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Success		200	{object}	helper.Response
// @Failure		400	{object}	helper.Response
// @Failure		404	{object}	helper.Response
// @Failure		500	{object}	helper.Response
// @Router			/permissions [get]
func (controller PermissionController) Index(ctx *gin.Context) {
	permissions, err := controller.service.List()

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Permission not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Permission Found", http.StatusOK, permissions)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Get one Permissions
// @Description	REST API for Permissions
// @Tags		Permissions
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		id	path		string	false	"Permissions ID"
// @Success		200	{object}	helper.Response
// @Failure		400	{object}	helper.Response
// @Failure		404	{object}	helper.Response
// @Failure		500	{object}	helper.Response
// @Router			/permissions/{id} [get]
func (controller PermissionController) Show(ctx *gin.Context) {
	ID := ctx.Param("id") // Get Param ID

	permission, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Permission not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.ErrorJSON(ctx, "Permission Found", http.StatusOK, permission)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Insert Permissions
// @Description	REST API for Permissions
// @Tags		Permissions
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		input	body		request.PermissionRequest	true	"Input data in JSON format"
// @Success		200		{object}	helper.Response
// @Failure		400		{object}	helper.Response
// @Failure		404		{object}	helper.Response
// @Failure		500		{object}	helper.Response
// @Router			/permissions [post]
func (controller PermissionController) Store(ctx *gin.Context) {
	var permissionReq request.PermissionRequest

	if err := ctx.ShouldBindJSON(&permissionReq); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create permission", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	permission, err := controller.service.Insert(permissionReq)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create permission", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Created Permission", http.StatusOK, permission)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Update Permissions
// @Description	REST API for Permissions
// @Tags		Permissions
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		id		query		string						false	"Permissions ID"
// @Param		input	body		request.PermissionRequest	true	"Input data in JSON format"
// @Success		200		{object}	helper.Response
// @Failure		400		{object}	helper.Response
// @Failure		404		{object}	helper.Response
// @Failure		500		{object}	helper.Response
// @Router			/permissions/{id} [put]
func (controller PermissionController) Update(ctx *gin.Context) {
	ID := ctx.Param("id") // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Permission not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var updateReq request.PermissionRequest

	if err := ctx.ShouldBindJSON(&updateReq); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update permission", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	permission, err := controller.service.Update(updateReq, ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update permission", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.SuccessJSON(ctx, "Successfully to Update Permission", http.StatusOK, permission)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Delete Permissions
// @Description	REST API for Permissions
// @Tags		Permissions
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		id	path		string	false	"Permissions ID"
// @Success		200	{object}	helper.Response
// @Failure		400	{object}	helper.Response
// @Failure		404	{object}	helper.Response
// @Failure		500	{object}	helper.Response
// @Router			/permissions/{id} [delete]
func (controller PermissionController) Delete(ctx *gin.Context) {
	ID := ctx.Param("id") // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Permission not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	permission, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to delete permission", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Delete Permission", http.StatusOK, permission)

	ctx.JSON(http.StatusOK, resp)
}
