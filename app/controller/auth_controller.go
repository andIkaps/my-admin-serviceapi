package controller

import (
	"fmt"
	"myadmin/app/request"
	"myadmin/app/service"
	"myadmin/helper"
	"myadmin/security"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service service.UserService
	auth    service.AuthService
}

func NewAuthController(s service.UserService, a service.AuthService) AuthController {
	return AuthController{
		service: s,
		auth:    a,
	}
}

// @Summary		Register User
// @Description	Handle register user and return the received data.
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		input	body		request.RegisterRequest	true	"Input data in JSON format"
// @Success		200		{object}	helper.Response
// @Failure		400		{object}	helper.Response
// @Failure		404		{object}	helper.Response
// @Failure		500		{object}	helper.Response
// @Router		/auth/register [post]
func (controller AuthController) Register(ctx *gin.Context) {
	var registerReq request.RegisterRequest

	if err := ctx.ShouldBindJSON(&registerReq); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create user", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	errValidate := helper.ValidateFromStruct(registerReq)

	if errValidate != nil {
		resp := helper.ErrorJSON(ctx, "Register Failed", http.StatusBadRequest, errValidate)

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	_, err := security.WeaknessPassword(registerReq.Password)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Password Weakness Detected", http.StatusBadRequest, err.Error())

		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		// ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	register, err := controller.service.Insert(registerReq)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create user", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Registered User", http.StatusOK, register)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary		Login User
// @Description	Handle login user and return the received data.
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		input	body		request.LoginRequest	true	"Input data in JSON format"
// @Success		200		{object}	helper.Response
// @Failure		400		{object}	helper.Response
// @Failure		404		{object}	helper.Response
// @Failure		500		{object}	helper.Response
// @Router		/auth/login [post]
func (controller AuthController) Login(ctx *gin.Context) {
	var loginReq request.LoginRequest

	// validate bind json
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		resp := helper.ErrorJSON(ctx, "Login Failed", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	errValidate := helper.ValidateFromStruct(loginReq)

	if errValidate != nil {
		resp := helper.ErrorJSON(ctx, "Login Failed", http.StatusBadRequest, errValidate)

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	// validate user struct
	// if err := validate.Struct(loginReq); err != nil {
	// 	errors := helper.FormatValidationErrors(err.(validator.ValidationErrors))
	// 	resp := helper.ErrorJSON(ctx, "Login Failed", http.StatusBadRequest, errors)

	// 	ctx.JSON(http.StatusBadRequest, resp)

	// 	return
	// }

	findUser, _ := controller.service.FindByUsername(loginReq.Username)

	if findUser.ID == "" {
		resp := helper.ErrorJSON(ctx, "The Username or Password is Incorrect", http.StatusUnprocessableEntity, nil)

		ctx.JSON(http.StatusUnprocessableEntity, resp)

		return
	}

	loggedIn, err := controller.auth.Login(loginReq)

	fmt.Println(err)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "The Username or Password is Incorrect", http.StatusBadRequest, nil)

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}
	token, err := security.GenerateToken(loggedIn.ID)

	// to := []string{
	// 	"azriel.fatur1@gmail.com",
	// }

	// utils.Email(to, []byte("Coba Kirim Email"))

	if err != nil {
		resp := helper.ErrorJSON(ctx, "There was an error generating the API token. Please try again", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Login Successfully", http.StatusOK, token)

	ctx.JSON(http.StatusOK, resp)
}
