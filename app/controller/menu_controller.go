package controller

import (
	"myadmin/app/request"
	"myadmin/app/service"
	"myadmin/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	service service.MenuService
}

func NewMenuController(s service.MenuService) MenuController {
	return MenuController{
		service: s,
	}
}

// @Summary		Get list menu (requires authentication)
// @Description	REST API Menu
// @Tags		Menus
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Success		200	{object}	helper.Response
// @Failure		400	{object}	helper.Response
// @Failure		404	{object}	helper.Response
// @Failure		500	{object}	helper.Response
// @Router			/menus [get]
func (controller MenuController) Index(ctx *gin.Context) {
	menus, err := controller.service.List()

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Menu not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Menu Found", http.StatusOK, menus)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Get one menu
// @Description	REST API menu
// @Tags		Menus
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		id	path		string	false	"Menu ID"
// @Success		200	{object}	helper.Response
// @Failure		400	{object}	helper.Response
// @Failure		404	{object}	helper.Response
// @Failure		500	{object}	helper.Response
// @Router		/menus/{id} [get]
func (controller MenuController) Show(ctx *gin.Context) {
	ID := ctx.Param("id") // Get Param ID

	menu, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Menu not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Menu Found", http.StatusOK, menu)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Insert menu
// @Description	REST API menu
// @Tags		Menus
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		input	body		request.MenuRequest	true	"Input data in JSON format"
// @Success		200		{object}	helper.Response
// @Failure		400		{object}	helper.Response
// @Failure		404		{object}	helper.Response
// @Failure		500		{object}	helper.Response
// @Router		/menus [post]
func (controller MenuController) Store(ctx *gin.Context) {
	var menuReq request.MenuRequest

	if err := ctx.ShouldBindJSON(&menuReq); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create menu", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	errValidate := helper.ValidateFromStruct(menuReq)

	if errValidate != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create menu", http.StatusBadRequest, errValidate)

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	menu, err := controller.service.Insert(menuReq)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create menu", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Created Menu", http.StatusOK, menu)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Update menu
// @Description	REST API menu
// @Tags		Menus
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		id		path		string				false	"Menu ID"
// @Param		input	body		request.MenuRequest	true	"Input data in JSON format"
// @Success		200		{object}	helper.Response
// @Failure		400		{object}	helper.Response
// @Failure		404		{object}	helper.Response
// @Failure		500		{object}	helper.Response
// @Router			/menus/{id} [put]
func (controller MenuController) Update(ctx *gin.Context) {
	ID := ctx.Param("id") // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Menu not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var updateReq request.MenuRequest

	if err := ctx.ShouldBindJSON(&updateReq); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update menu", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	menu, err := controller.service.Update(updateReq, ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update menu", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.SuccessJSON(ctx, "Successfully to Update Menu", http.StatusOK, menu)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Delete menu
// @Description	REST API menu
// @Tags        Menus
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		id	path		string	false	"Menu ID"
// @Success		200	{object}	helper.Response
// @Failure		400	{object}	helper.Response
// @Failure		404	{object}	helper.Response
// @Failure		500	{object}	helper.Response
// @Router			/menus/{id} [delete]
func (controller MenuController) Delete(ctx *gin.Context) {
	ID := ctx.Param("id") // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Menu not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	menu, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to delete menu", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Delete Menu", http.StatusOK, menu)

	ctx.JSON(http.StatusOK, resp)
}
