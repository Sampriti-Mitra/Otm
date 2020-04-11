package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"otm/app/dtos"
	"sort"
)

type feedController struct {
	baseController
}

var (
	FeedController feedController
)

func NewFeedController() {
	FeedController = feedController{}
}

func (controller feedController) Feed(ctx *gin.Context) {
	profileId := ctx.Param("profile_id")
	FeedResponse := []dtos.UploadResponse{}
	response, iError := FollowerInstance().ListFeed(ctx, profileId)
	if iError != nil {
		controller.setErrorResponse(ctx, iError)
		return
	}
	for _, resp := range response {
		userResp, iError := LoginInstance().GetByUsername(ctx, resp.RequestTo)
		if iError != nil {
			controller.setErrorResponse(ctx, iError)
			return
		}
		feedResp, iError := ProfileInstance().ListFeed(ctx, userResp.Id)
		if iError != nil {
			controller.setErrorResponse(ctx, iError)
			return
		}
		FeedResponse = append(FeedResponse, feedResp...)
	}
	sort.Slice(FeedResponse, func(i, j int) bool {
		return FeedResponse[i].CreatedAt.After(FeedResponse[j].CreatedAt)
	})
	ctx.JSON(http.StatusOK, FeedResponse)
	return
}
