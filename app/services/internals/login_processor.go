package internals

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"otm/app/dtos"
	"otm/app/interfaces"
	"otm/app/repository"
	"strconv"
)

type loginProcessor struct {
}

func (l loginProcessor) Create(ctx *gin.Context, request dtos.RegisterRequest) (dtos.RegisterResponse, error) {
	var response dtos.RegisterResponse
	repo := repository.GetLoginRepo()
	_, iError := repo.Transaction(
		ctx,
		func() (data interface{}, iError error) {

			// Create func for namespace
			iError = repo.Create(ctx, &request)
			if iError != nil {
				return response, iError
			}

			//Create Domain Model
			var aboutRequest dtos.AboutRequest
			aboutRequest.UserId = request.Id
			aboutRequest.About = ""
			response.AboutResponse, iError = GetAboutProcessor().Create(ctx, aboutRequest)
			if iError != nil {
				return
			}

			return data, iError
		},
	)
	if iError != nil {
		return response, iError
	}

	response.Model = request.Model
	response.Success = true
	response.Username = request.Username
	return response, nil
}

func (l loginProcessor) Get(ctx *gin.Context, requestId int) (dtos.RegisterResponse, error) {
	var response dtos.RegisterResponse
	var users dtos.Users
	repo := repository.GetLoginRepo()
	err := repo.Find(ctx, &users, map[string]interface{}{
		"created_by": requestId,
	})
	response.Username = users.Username
	response.Model = users.Model
	if err != nil {
		response.Success = false
	} else {
		response.Success = true
	}
	return response, err
}

func (l loginProcessor) Update(ctx *gin.Context, request dtos.RegisterRequest, registrationId string) (dtos.RegisterResponse, error) {
	var users dtos.Users
	var response dtos.RegisterResponse
	repo := repository.GetLoginRepo()
	attributes := map[string]interface{}{
		"username": request.Username,
		"name":     request.Name,
	}
	err := repo.Update(ctx, &users, attributes, map[string]interface{}{
		"id": registrationId,
	})
	if err != nil {
		return response, err
	}
	response.Model = users.Model
	response.Success = true
	response.Username = users.Username
	return response, nil
}

func (l loginProcessor) Delete(ctx *gin.Context, requestId string) dtos.DeletedResponse {
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

func (l loginProcessor) GetByUsername(ctx *gin.Context, username string) (dtos.Users, error) {
	var users dtos.Users
	repo := repository.GetLoginRepo()
	err := repo.Find(ctx, &users, map[string]interface{}{
		"username": username,
	})
	return users, err
}

func GetLoginProcessor() interfaces.ILogin {
	return loginProcessor{}
}
