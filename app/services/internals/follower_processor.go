package internals

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"otm/app/dtos"
	"otm/app/interfaces"
	"otm/app/repository"
	"strconv"
)

type followerProcessor struct {
}

func (l followerProcessor) Create(ctx *gin.Context, request dtos.FollowRequest) (dtos.FollowResponse, error) {
	var response dtos.FollowResponse
	repo := repository.GetFollowerRepo()
	iError := repo.Create(ctx, &request)
	if iError != nil {
		return response, iError
	}

	response.Model = request.Model
	response.RequestBy = request.RequestBy
	response.RequestTo = request.RequestTo
	return response, nil
}

func (l followerProcessor) List(ctx *gin.Context, userId string) ([]dtos.FollowResponse, error) {
	var response []dtos.FollowResponse
	repo := repository.GetFollowerRepo()
	err := repo.FindAll(ctx, &response, map[string]interface{}{
		"request_to": userId,
	})
	return response, err
}

func (l followerProcessor) Update(ctx *gin.Context, request dtos.FollowRequest, registrationId string) (dtos.FollowResponse, error) {
	var users dtos.Users
	var response dtos.FollowResponse
	repo := repository.GetLoginRepo()
	attributes := map[string]interface{}{
		//"username": request.Username,
		//"name":     request.Name,
	}
	err := repo.Update(ctx, &users, attributes, map[string]interface{}{
		"id": registrationId,
	})
	if err != nil {
		return response, err
	}
	response.Model = users.Model
	//response.Success = true
	//response.Username = users.Username
	return response, nil
}

func (l followerProcessor) Delete(ctx *gin.Context, requestId string) dtos.DeletedResponse {
	var errorResp dtos.DeletedResponse
	var users dtos.Users
	repo := repository.GetLoginRepo()
	_, err := repo.Transaction(
		ctx,
		func() (data interface{}, iError error) {

			// Create func for namespace
			errResp := repo.Delete(ctx, &users, map[string]interface{}{
				"id": requestId,
			})
			if errResp != nil {
				return nil, errResp
			}

			//Create Domain Model
			userId, iError := strconv.Atoi(requestId)
			if iError != nil {
				return
			}
			iError = GetAboutProcessor().Delete(ctx, userId)
			if iError != nil {
				return
			}

			return data, iError
		},
	)
	if err != nil {
		errorResp.Success = false
	} else {
		errorResp.Success = true
	}
	errorResp.Id = fmt.Sprintf("%d", users.Id)
	errorResp.Error = err
	return errorResp
}

func GetFollowerProcessor() interfaces.IFollower {
	return followerProcessor{}
}
