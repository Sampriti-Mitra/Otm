package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"otm/app/constants"
	"otm/app/dtos"
	"otm/app/services/internals"
)

type followerController struct {
	baseController
}

var (
	FollowerController followerController
	FollowerInstance   = internals.GetFollowerProcessor
)

func NewFollowerController() {
	FollowerController = followerController{}
}

//TODO:: password encrypt
func (controller followerController) Follow(ctx *gin.Context) {
	followRequest := dtos.FollowRequest{}
	profileId := ctx.Param("profile_id")

	iError := controller.deserialize(ctx, &followRequest)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}
	followRequest.RequestTo = profileId

	createNSResp, iError := FollowerInstance().Create(ctx, followRequest)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	ctx.JSON(http.StatusOK, createNSResp)
	return
}

func (controller followerController) ListFollowers(ctx *gin.Context) {
	profileId := ctx.Param("profile_id")
	response, iError := FollowerInstance().List(ctx, profileId)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}
	ctx.JSON(http.StatusOK, response)
	return
}

func (controller followerController) ListFollowerRequest(ctx *gin.Context) {
	profileId := ctx.Param("profile_id")
	response, iError := FollowerInstance().ListFollowerRequest(ctx, profileId)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}
	ctx.JSON(http.StatusOK, response)
	return
}

func (controller followerController) ListFollowing(ctx *gin.Context) {
	profileId := ctx.Param("profile_id")
	response, iError := FollowerInstance().ListFollowing(ctx, profileId)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}
	ctx.JSON(http.StatusOK, response)
	return
}

func (controller followerController) AcceptFollowers(ctx *gin.Context) {
	profileId := ctx.Param("profile_id")
	request_by := ctx.Param("request_by")

	createNSResp, iError := FollowerInstance().UpdateAccept(ctx, profileId, request_by)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	ctx.JSON(http.StatusOK, createNSResp)
	return
}

func (controller followerController) RejectFollowers(ctx *gin.Context) {
	profileId := ctx.Param("profile_id")
	request_by := ctx.Param("request_by")

	createNSResp, iError := FollowerInstance().UpdateReject(ctx, profileId, request_by)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}

	ctx.JSON(http.StatusOK, createNSResp)
	return
}

func (controller followerController) DeleteRegister(ctx *gin.Context) {
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

func (controller followerController) GetFollower(ctx *gin.Context) {
	profileId := ctx.Param("profile_id")
	request_by := ctx.Param("request_by")

	resp, iError := FollowerInstance().Get(ctx, profileId, request_by)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}
	ctx.JSON(http.StatusOK, resp)
	return
}
