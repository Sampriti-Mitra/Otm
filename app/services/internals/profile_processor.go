package internals

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"otm/app/dtos"
	"otm/app/interfaces"
	"otm/app/repository"
	"otm/app/utils"
)

type profileProcessor struct {
}

func (l profileProcessor) Create(ctx *gin.Context, request dtos.UploadRequest) (dtos.UploadResponse, error) {
	var response dtos.UploadResponse
	repo := repository.GetProfileRepo()
	err := repo.Create(ctx, &request)
	if err != nil {
		return response, err
	}
	l.Mapper(request, &response)
	response.Success = true
	return response, nil
}

func (l profileProcessor) List(ctx *gin.Context, requestId int) ([]dtos.UploadResponse, error) {
	var response []dtos.UploadResponse
	repo := repository.GetProfileRepo()
	err := repo.FindAll(ctx, &response, map[string]interface{}{
		"created_by": fmt.Sprintf("%d", requestId),
	})
	return response, err
}

func (l profileProcessor) Get(ctx *gin.Context, videoId int) (dtos.UploadResponse, error) {
	var response dtos.UploadResponse
	repo := repository.GetProfileRepo()
	err := repo.Find(ctx, &response, map[string]interface{}{
		"id":         videoId,
	})
	return response, err
}

func (l profileProcessor) Update(ctx *gin.Context, request dtos.UploadRequest, videoId int) (dtos.UploadResponse, error) {
	var response dtos.UploadResponse
	repo := repository.GetProfileRepo()
	attributes := map[string]interface{}{
		"title": request.Title,
		"tags":  request.Tags,
	}
	err := repo.Update(ctx, &response, attributes, map[string]interface{}{
		"id": videoId,
	})
	if err != nil {
		return response, err
	}
	return response, nil
}

func (l profileProcessor) Delete(ctx *gin.Context, requestId string) dtos.DeletedResponse {
	var errorResp dtos.DeletedResponse
	var user dtos.Upload
	repo := repository.GetProfileRepo()
	err := repo.Delete(ctx, &user, map[string]interface{}{
		"id": requestId,
	})
	if err != nil {
		errorResp.Success = false
	} else {
		errorResp.Success = true
	}
	errorResp.Id = fmt.Sprintf("%d", user.Id)
	errorResp.Error = err
	return errorResp
}

func (l profileProcessor) GetByUsername(ctx *gin.Context, username string) (dtos.UploadResponse, error) {
	var users dtos.UploadResponse
	repo := repository.GetLoginRepo()
	err := repo.Find(ctx, &users, map[string]interface{}{
		"username": username,
	})
	return users, err
}

func (l profileProcessor) Applaud(ctx *gin.Context, requestedBy string, videoId int) (dtos.Applaud, error) {
	var response dtos.Applaud
	var applaudedBy []string
	var applauses int
	var like string
	var modelResp dtos.UploadResponse
	repo := repository.GetProfileRepo()

	respCurrent,err:=l.Get(ctx,videoId)
	if err != nil {
		return response, err
	}

	err=json.Unmarshal(respCurrent.ApplaudedBy,&applaudedBy)
	if err != nil &&respCurrent.ApplaudedBy!=nil {
		return response, err
	}

	if respCurrent.ApplaudedBy==nil || !utils.Contains(applaudedBy,requestedBy){
		applaudedBy=append(applaudedBy,requestedBy)
		applauses=respCurrent.Applause+1
		like="applauded"
	}else {
		applaudedBy=utils.Remove(applaudedBy,requestedBy)
		applauses=respCurrent.Applause-1
		like="applaud removed"
	}

	applaudedJson,err:=json.Marshal(applaudedBy)

	attributes := map[string]interface{}{
		"applauded_by": applaudedJson,
		"applause":  applauses,
	}
	err = repo.Update(ctx, &modelResp, attributes, map[string]interface{}{
		"id": videoId,
	})
	if err!=nil{
		response.Success=false
		response.Message="Applauding failed"
		response.Model=modelResp.Model
	}else{
		response.Success=true
		response.Message=like
		response.Model=modelResp.Model
	}

	return response, err
}

func GetProfileProcessor() interfaces.IProfile {
	return profileProcessor{}
}

func (l profileProcessor) Mapper(request dtos.UploadRequest, response *dtos.UploadResponse) {
	response.Model = request.Model
	response.CreatedBy = request.CreatedBy
	response.Tags = request.Tags
	response.Title = request.Title
	response.Videolink = request.Videolink
}
