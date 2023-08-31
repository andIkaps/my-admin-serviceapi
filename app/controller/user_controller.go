package controller

import (
	"myadmin/app/request"
	"myadmin/app/service"
	"myadmin/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
	role    service.UserRoleService
}

func NewUserController(s service.UserService, r service.UserRoleService) UserController {
	return UserController{
		service: s,
		role:    r,
	}
}

// @Summary		Get list Users
// @Description	REST API for Users
// @Tags		Users
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Success		200	{object}	helper.Response
// @Failure		400	{object}	helper.Response
// @Failure		404	{object}	helper.Response
// @Failure		500	{object}	helper.Response
// @Router		/users [get]
func (controller UserController) Index(ctx *gin.Context) {
	users, err := controller.service.List()

	if err != nil {
		resp := helper.ErrorJSON(ctx, "No Data Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Data Found", http.StatusOK, users)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Get one Users
// @Description	REST API for Users
// @Tags		Users
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		id	path		string	false	"Users ID"
// @Success		200	{object}	helper.Response
// @Failure		400	{object}	helper.Response
// @Failure		404	{object}	helper.Response
// @Failure		500	{object}	helper.Response
// @Router		/users/{id} [get]
func (controller UserController) Show(ctx *gin.Context) {
	ID := ctx.Param("id") // Get Param ID

	user, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "User not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "User Found", http.StatusOK, user)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Update Users
// @Description	REST API for Users
// @Tags		Users
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		id		query		string						false	"Users ID"
// @Param		input	body		request.UserUpdateRequest	true	"Input data in JSON format"
// @Success		200		{object}	helper.Response
// @Failure		400		{object}	helper.Response
// @Failure		404		{object}	helper.Response
// @Failure		500		{object}	helper.Response
// @Router		/users/{id} [put]
func (controller UserController) Update(ctx *gin.Context) {
	ID := ctx.Param("id") // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "User not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	var updateReq request.UserUpdateRequest

	if err := ctx.ShouldBindJSON(&updateReq); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update user", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	errValidate := helper.ValidateFromStruct(updateReq)

	if errValidate != nil {
		resp := helper.ErrorJSON(ctx, "Register Failed", http.StatusBadRequest, errValidate)

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	user, err := controller.service.Update(updateReq, ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update user", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Updated User", http.StatusOK, user)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Delete Users
// @Description	REST API for Users
// @Tags		Users
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		id	path		string	false	"Users ID"
// @Success		200	{object}	helper.Response
// @Failure		400	{object}	helper.Response
// @Failure		404	{object}	helper.Response
// @Failure		500	{object}	helper.Response
// @Router		/users/{id} [delete]
func (controller UserController) Delete(ctx *gin.Context) {
	ID := ctx.Param("id") // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "User not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	user, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to delete user", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Deleted User", http.StatusOK, user)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Attach user has role
// @Description	REST API User Role
// @Tags		Users
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		input	body		request.UserHasRole	true	"Input data in JSON format"
// @Success		200		{object}	helper.Response
// @Failure		400		{object}	helper.Response
// @Failure		404		{object}	helper.Response
// @Failure		500		{object}	helper.Response
// @Router		/users/roles/assign [post]
func (controller UserController) AttachRole(ctx *gin.Context) {
	var req request.UserHasRole

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Attach role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	attach, err := controller.role.AttachRole(req)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Attach role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Attached Role", http.StatusOK, attach)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Detach user has role
// @Description	REST API User Role
// @Tags		Users
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		id		query		string				false	"Users ID"
// @Param		input	body		request.UserHasRole	true	"Input data in JSON format"
// @Success		200		{object}	helper.Response
// @Failure		400		{object}	helper.Response
// @Failure		404		{object}	helper.Response
// @Failure		500		{object}	helper.Response
// @Router		/users/roles/unassign [delete]
func (controller UserController) DetachRole(ctx *gin.Context) {
	var req request.UserHasRole

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Detach role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	detach, err := controller.role.DetachRole(req)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Detach role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Detached Role", http.StatusOK, detach)

	ctx.JSON(http.StatusOK, resp)
}
