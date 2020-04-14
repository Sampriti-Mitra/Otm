package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"otm/app/dtos"
	"otm/app/services/internals"
	"strconv"
)

type collabController struct {
	baseController
}

var (
	CollabController collabController
	CollabInstance   = internals.GetCollabProcessor
)

func NewCollabController() {
	CollabController = collabController{}
}

//TODO:: password encrypt
func (controller collabController) CreateCollab(ctx *gin.Context) {
	aboutRequest := dtos.CollabRequest{}

	iError := controller.deserialize(ctx, &aboutRequest)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	createNSResp, iError := CollabInstance().Create(ctx, aboutRequest)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	ctx.JSON(http.StatusOK, createNSResp)
	return
}

func (controller collabController) GetCollab(ctx *gin.Context) {
	collabId, iError := strconv.Atoi(ctx.Param("collab_id"))
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}
	collabInstance := internals.GetCollabProcessor
	collabResponse, iError := collabInstance().Get(ctx, collabId)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	ctx.JSON(http.StatusOK, collabResponse)
	return
}

func (controller collabController) AcceptCollab(ctx *gin.Context) {
	username := ctx.Param("username")
	collabId, iError := strconv.Atoi(ctx.Param("collab_id"))
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}
	createNSResp, iError := CollabInstance().AcceptCollab(ctx, collabId, username)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	ctx.JSON(http.StatusOK, createNSResp)
	return
}

func (controller collabController) RejectCollab(ctx *gin.Context) {
	username := ctx.Param("username")
	collabId, iError := strconv.Atoi(ctx.Param("collab_id"))
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}
	createNSResp, iError := CollabInstance().RejectCollab(ctx, collabId, username)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	ctx.JSON(http.StatusOK, createNSResp)
	return
}

func (controller collabController) UpdateCollab(ctx *gin.Context) {
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
