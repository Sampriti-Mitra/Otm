package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"otm/app/constants"
	"otm/app/dtos"
	"otm/app/services/internals"
	"strconv"
)

type loginController struct {
	baseController
}

var (
	LoginController loginController
	LoginInstance   = internals.GetLoginProcessor
)

func NewLoginController() {
	LoginController = loginController{}
}

//TODO:: password encrypt
func (controller loginController) CreateRegister(ctx *gin.Context) {
	registerRequest := dtos.RegisterRequest{}

	iError := controller.deserialize(ctx, &registerRequest)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	createNSResp, iError := LoginInstance().Create(ctx, registerRequest)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	ctx.JSON(http.StatusOK, createNSResp)
	return
}

func (controller loginController) GetRegister(ctx *gin.Context) {
	RegistrationId, err := strconv.Atoi(ctx.Param(constants.RegistrationId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	response, iError := LoginInstance().Get(ctx, RegistrationId)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	ctx.JSON(http.StatusOK, response)
	return
}

func (controller loginController) UpdateRegister(ctx *gin.Context) {
	registerRequest := dtos.RegisterRequest{}
	RegistrationId := ctx.Param(constants.RegistrationId)

	iError := controller.deserialize(ctx, &registerRequest)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	createNSResp, iError := LoginInstance().Update(ctx, registerRequest, RegistrationId)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	ctx.JSON(http.StatusOK, createNSResp)
	return
}

func (controller loginController) DeleteRegister(ctx *gin.Context) {
	RegistrationId := ctx.Param(constants.RegistrationId)

	iError := LoginInstance().Delete(ctx, RegistrationId)
	if iError.Error != nil {
		controller.setErrorResponse(ctx, iError.Error)
		return
	}

	var response dtos.DeletedResponse
	response.Id = RegistrationId
	response.Success = true

	ctx.JSON(http.StatusOK, response)
	return
}

func (controller loginController) GetLogin(ctx *gin.Context) {
	registerRequest := dtos.LoginRequest{}

	iError := controller.deserialize(ctx, &registerRequest)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	resp, iError := LoginInstance().GetByUsername(ctx, registerRequest.Username)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}
	if registerRequest.Password == resp.Password {
		ctx.JSON(http.StatusOK, resp)
		return
	}
	iError = errors.New("Password invalid")
	controller.setErrorResponse(ctx, iError)
	return
}
