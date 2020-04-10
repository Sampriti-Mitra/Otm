package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"otm/app/dtos"
	"otm/app/services/internals"
	"strconv"
)

type aboutController struct {
	baseController
}

var (
	AboutController aboutController
	AboutInstance   = internals.GetAboutProcessor
)

func NewAboutController() {
	AboutController = aboutController{}
}

//TODO:: password encrypt
func (controller aboutController) CreateAbout(ctx *gin.Context) {
	aboutRequest := dtos.AboutRequest{}

	iError := controller.deserialize(ctx, &aboutRequest)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	createNSResp, iError := AboutInstance().Create(ctx, aboutRequest)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	ctx.JSON(http.StatusOK, createNSResp)
	return
}

func (controller aboutController) GetAbout(ctx *gin.Context) {
	userId, iError := strconv.Atoi(ctx.Param("user_id"))
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}
	aboutInstance := internals.GetAboutProcessor
	aboutResponse, iError := aboutInstance().Get(ctx, userId)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	ctx.JSON(http.StatusOK, aboutResponse)
	return
}

func (controller aboutController) UpdateAbout(ctx *gin.Context) {
	aboutRequest := dtos.AboutRequest{}
	userId, iError := strconv.Atoi(ctx.Param("user_id"))
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	iError = controller.deserialize(ctx, &aboutRequest)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	createNSResp, iError := AboutInstance().Update(ctx, aboutRequest, userId)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	ctx.JSON(http.StatusOK, createNSResp)
	return
}
