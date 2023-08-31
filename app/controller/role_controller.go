package controller

import (
	"myadmin/app/request"
	"myadmin/app/service"
	"myadmin/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	service    service.RoleService
	menu       service.RoleMenuService
	permission service.RolePermissionService
}

func NewRoleController(s service.RoleService, m service.RoleMenuService, p service.RolePermissionService) RoleController {
	return RoleController{
		service:    s,
		menu:       m,
		permission: p,
	}
}

// @Summary		Get list Roles
// @Description	REST API for Roles
// @Tags		Roles
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Success		200	{object}	helper.Response
// @Failure		400	{object}	helper.Response
// @Failure		404	{object}	helper.Response
// @Failure		500	{object}	helper.Response
// @Router		/roles [get]
func (controller RoleController) Index(ctx *gin.Context) {
	roles, err := controller.service.List()

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Role not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Role Found", http.StatusOK, roles)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Get one Roles
// @Description	REST API for Roles
// @Tags		Roles
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		id	path		string	false	"Roles ID"
// @Success		200	{object}	helper.Response
// @Failure		400	{object}	helper.Response
// @Failure		404	{object}	helper.Response
// @Failure		500	{object}	helper.Response
// @Router		/roles/{id} [get]
func (controller RoleController) Show(ctx *gin.Context) {
	ID := ctx.Param("id") // Get Param ID

	role, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Role not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.ErrorJSON(ctx, "Role Found", http.StatusOK, role)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Insert Roles
// @Description	REST API for Roles
// @Tags		Roles
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		input	body		request.RoleRequest	true	"Input data in JSON format"
// @Success		200		{object}	helper.Response
// @Failure		400		{object}	helper.Response
// @Failure		404		{object}	helper.Response
// @Failure		500		{object}	helper.Response
// @Router		/roles [post]
func (controller RoleController) Store(ctx *gin.Context) {
	var roleReq request.RoleRequest

	if err := ctx.ShouldBindJSON(&roleReq); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	errValidate := helper.ValidateFromStruct(roleReq)

	if errValidate != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create role", http.StatusBadRequest, errValidate)

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	role, err := controller.service.Insert(roleReq)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Created Role", http.StatusOK, role)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Update Roles
// @Description	REST API for Roles
// @Tags		Roles
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		id		query		string				false	"Roles ID"
// @Param		input	body		request.RoleRequest	true	"Input data in JSON format"
// @Success		200		{object}	helper.Response
// @Failure		400		{object}	helper.Response
// @Failure		404		{object}	helper.Response
// @Failure		500		{object}	helper.Response
// @Router		/roles/{id} [put]
func (controller RoleController) Update(ctx *gin.Context) {
	ID := ctx.Param("id") // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Role not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var updateReq request.RoleRequest

	if err := ctx.ShouldBindJSON(&updateReq); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	errValidate := helper.ValidateFromStruct(updateReq)

	if errValidate != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update role", http.StatusBadRequest, errValidate)

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	role, err := controller.service.Update(updateReq, ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.SuccessJSON(ctx, "Successfully to Update Role", http.StatusOK, role)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Delete Roles
// @Description	REST API for Roles
// @Tags		Roles
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		id	path		string	false	"Roles ID"
// @Success		200	{object}	helper.Response
// @Failure		400	{object}	helper.Response
// @Failure		404	{object}	helper.Response
// @Failure		500	{object}	helper.Response
// @Router		/roles/{id} [delete]
func (controller RoleController) Delete(ctx *gin.Context) {
	ID := ctx.Param("id") // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Role not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	role, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to delete role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Delete Role", http.StatusOK, role)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Attach role has menu
// @Description	REST API Role Menu
// @Tags		Roles
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		input	body		request.RoleHasMenu	true	"Input data in JSON format"
// @Success		200		{object}	helper.Response
// @Failure		400		{object}	helper.Response
// @Failure		404		{object}	helper.Response
// @Failure		500		{object}	helper.Response
// @Router		/roles/menus [post]
func (controller RoleController) AttachMenu(ctx *gin.Context) {
	var req request.RoleHasMenu

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Attach menu", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	attach, err := controller.menu.AttachMenu(req)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Attach menu", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Attached Menu", http.StatusOK, attach)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Detach role has menus
// @Description	REST API Role Menu
// @Tags		Roles
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		input	body		request.RoleHasMenu	true	"Input data in JSON format"
// @Success		200		{object}	helper.Response
// @Failure		400		{object}	helper.Response
// @Failure		404		{object}	helper.Response
// @Failure		500		{object}	helper.Response
// @Router		/roles/menu [delete]
func (controller RoleController) DetachMenu(ctx *gin.Context) {
	var req request.RoleHasMenu
	ID := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Detach menu", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	detach, err := controller.menu.DetachMenu(req, ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Detach menu", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Detached Role", http.StatusOK, detach)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Attach role has permissions
// @Description	REST API Role Permission
// @Tags		Roles
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		input	body		request.RoleHasPermission	true	"Input data in JSON format"
// @Success		200		{object}	helper.Response
// @Failure		400		{object}	helper.Response
// @Failure		404		{object}	helper.Response
// @Failure		500		{object}	helper.Response
// @Router		/roles/permissions [post]
func (controller RoleController) AttachPermission(ctx *gin.Context) {
	var req request.RoleHasPermission

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Attach permission", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	attach, err := controller.permission.AttachPermission(req)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Attach permission", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Attached Permission", http.StatusOK, attach)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Detach role has permissions
// @Description	REST API Role Permission
// @Tags		Roles
// @Tags		Roles
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		input	body		request.RoleHasPermission	true	"Input data in JSON format"
// @Success		200		{object}	helper.Response
// @Failure		400		{object}	helper.Response
// @Failure		404		{object}	helper.Response
// @Failure		500		{object}	helper.Response
// @Router		/roles/permissions [delete]
func (controller RoleController) DetachPermission(ctx *gin.Context) {
	var req request.RoleHasPermission

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Detach permission", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	detach, err := controller.permission.DetachPermission(req)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Detach permission", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Detached Permission", http.StatusOK, detach)

	ctx.JSON(http.StatusOK, resp)
}
