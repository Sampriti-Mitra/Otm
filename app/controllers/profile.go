package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"otm/app/dtos"
	"otm/app/services/internals"
	"strconv"
)

type profileController struct {
	baseController
}

var (
	ProfileController profileController
	ProfileInstance   = internals.GetProfileProcessor
)

func NewProfileController() {
	ProfileController = profileController{}
}

//TODO:: password encrypt
func (controller profileController) CreatePost(ctx *gin.Context) {
	uploadRequest := dtos.UploadRequest{}

	iError := controller.deserialize(ctx, &uploadRequest)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	createNSResp, iError := ProfileInstance().Create(ctx, uploadRequest)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	ctx.JSON(http.StatusOK, createNSResp)
	return
}

func (controller profileController) ListPost(ctx *gin.Context) {
	userId := ctx.Param("profile_id")

	loginInstance := internals.GetLoginProcessor
	loginresponse, iError := loginInstance().GetByUsername(ctx, userId)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	postResponse, iError := ProfileInstance().List(ctx, loginresponse.Id)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}
	aboutResponse, iError := AboutInstance().Get(ctx, loginresponse.Id)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}
	var profileResponse dtos.Profile
	profileResponse.Posts = postResponse
	profileResponse.AboutResponse = aboutResponse

	ctx.JSON(http.StatusOK, profileResponse)
	return
}

func (controller profileController) GetPost(ctx *gin.Context) {
	Request := dtos.Upload{}
	VideoId, iError := strconv.Atoi(ctx.Param("video_id"))
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}
	iError = controller.deserialize(ctx, &Request)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}
	loginInstance := internals.GetLoginProcessor
	loginresponse, iError := loginInstance().GetByUsername(ctx, Request.RequestedBy)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	fmt.Println(loginresponse)

	response, iError := ProfileInstance().Get(ctx, VideoId)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	ctx.JSON(http.StatusOK, response)
	return
}

func (controller profileController) UpdatePost(ctx *gin.Context) {
	registerRequest := dtos.UploadRequest{}
	VideoId, iError := strconv.Atoi(ctx.Param("video_id"))
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	iError = controller.deserialize(ctx, &registerRequest)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	createNSResp, iError := ProfileInstance().Update(ctx, registerRequest, VideoId)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	ctx.JSON(http.StatusOK, createNSResp)
	return
}

func (controller profileController) DeletePost(ctx *gin.Context) {
	videoId := ctx.Param("video_id")

	iError := ProfileInstance().Delete(ctx, videoId)
	if iError.Error != nil {
		controller.setErrorResponse(ctx, iError.Error)
		return
	}

	var response dtos.DeletedResponse
	response.Id = videoId
	response.Success = true

	ctx.JSON(http.StatusOK, response)
	return
}

func (controller profileController) ApplausePost(ctx *gin.Context) {
	Request := dtos.Upload{} // who is applauding the post
	VideoId, iError := strconv.Atoi(ctx.Param("video_id"))
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}
	iError = controller.deserialize(ctx, &Request)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	response, iError := ProfileInstance().Applaud(ctx, Request.RequestedBy, VideoId)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	ctx.JSON(http.StatusOK, response)
	return
}
