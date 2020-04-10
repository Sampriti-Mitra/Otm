package controllers

import (
	"errors"
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
	var resp []dtos.FollowResponse
	for _, val := range response {
		resp = append(resp, val)
	}
	ctx.JSON(http.StatusOK, resp)
	return
}

func (controller followerController) AcceptOrRejectFollowers(ctx *gin.Context) {
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

func (controller followerController) GetLogin(ctx *gin.Context) {
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
