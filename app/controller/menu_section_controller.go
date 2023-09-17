package controller

import (
	"myadmin/app/formatter"
	"myadmin/app/request"
	"myadmin/app/service"
	"myadmin/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MenuSectionController struct {
	service service.MenuSectionService
}

func NewMenuSectionController(s service.MenuSectionService) MenuSectionController {
	return MenuSectionController{
		service: s,
	}
}

// @Summary		Get list menu section (requires authentication)
// @Description	REST API Menu Section
// @Tags		Menu Sections
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Success		200	{object}	helper.Response
// @Failure		400	{object}	helper.Response
// @Failure		404	{object}	helper.Response
// @Failure		500	{object}	helper.Response
// @Router		/menu-section [get]
func (controller MenuSectionController) Index(ctx *gin.Context) {
	menuSections, err := controller.service.List()

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Menu Section not found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)
		return
	}

	resp := helper.SuccessJSON(ctx, "Menu section found", http.StatusOK, formatter.FormatMenuSections(menuSections))
	// resp := helper.SuccessJSON(ctx, "Menu section found", http.StatusOK, menuSections)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Get one menu section
// @Description	REST API menu section
// @Tags		Menu Sections
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		id	path		string	false	"Menu ID"
// @Success		200	{object}	helper.Response
// @Failure		400	{object}	helper.Response
// @Failure		404	{object}	helper.Response
// @Failure		500	{object}	helper.Response
// @Router		/menu-section/{id} [get]
func (controller MenuSectionController) Show(ctx *gin.Context) {
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

// @Summary		Insert menu section
// @Description	REST API menu section
// @Tags		Menu Sections
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		input	body		request.MenuRequest	true	"Input data in JSON format"
// @Success		200		{object}	helper.Response
// @Failure		400		{object}	helper.Response
// @Failure		404		{object}	helper.Response
// @Failure		500		{object}	helper.Response
// @Router		/menu-section [post]
func (controller MenuSectionController) Store(ctx *gin.Context) {
	var menuReq request.MenuSectionRequest

	if err := ctx.ShouldBindJSON(&menuReq); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create menu section", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	errValidate := helper.ValidateFromStruct(menuReq)

	if errValidate != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create menu section", http.StatusBadRequest, errValidate)

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	menu, err := controller.service.Insert(menuReq)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create menu section", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Created Menu Section", http.StatusOK, menu)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Update menu section
// @Description	REST API menu section
// @Tags		Menu Sections
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		id		path		string				false	"Menu ID"
// @Param		input	body		request.MenuRequest	true	"Input data in JSON format"
// @Success		200		{object}	helper.Response
// @Failure		400		{object}	helper.Response
// @Failure		404		{object}	helper.Response
// @Failure		500		{object}	helper.Response
// @Router		/menu-section/{id} [put]
func (controller MenuSectionController) Update(ctx *gin.Context) {
	ID := ctx.Param("id") // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Menu Section not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var updateReq request.MenuSectionRequest

	if err := ctx.ShouldBindJSON(&updateReq); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update menu section", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	menu, err := controller.service.Update(updateReq, ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update menu section", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.SuccessJSON(ctx, "Successfully to Update Menu Section", http.StatusOK, menu)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Delete menu section
// @Description	REST API menu section
// @Tags        Menu Sections
// @Accept		json
// @Produce		json
// @Security	BearerAuth
// @Param		id	path		string	false	"Menu ID"
// @Success		200	{object}	helper.Response
// @Failure		400	{object}	helper.Response
// @Failure		404	{object}	helper.Response
// @Failure		500	{object}	helper.Response
// @Router		/menu-section/{id} [delete]
func (controller MenuSectionController) Delete(ctx *gin.Context) {
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
