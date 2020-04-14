package internals

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"otm/app/dtos"
	"otm/app/interfaces"
	"otm/app/repository"
)

type collabProcessor struct {
}

func (l collabProcessor) Create(ctx *gin.Context, request dtos.CollabRequest) (dtos.CollabResponse, error) {
	var response dtos.CollabResponse
	var resp dtos.Collab
	l.MemberStatusDefault(&request)
	l.Mapper(request, &resp)
	repo := repository.GetCollabRepo()
	err := repo.Create(ctx, &resp)
	if err != nil {
		return response, err
	}
	l.ResponseMapper(resp, &response)
	return response, nil
}

func (l collabProcessor) Get(ctx *gin.Context, collab_id int) (dtos.CollabResponse, error) {
	var collab dtos.Collab
	var response dtos.CollabResponse
	repo := repository.GetCollabRepo()
	err := repo.Find(ctx, &collab, map[string]interface{}{
		"id": collab_id,
	})
	l.ResponseMapper(collab, &response)
	return response, err
}

func (l collabProcessor) AcceptCollab(ctx *gin.Context, collabId int, username string) (dtos.CollabResponse, error) {
	var response dtos.CollabResponse
	var resp dtos.Collab
	CollabReq, iError := l.Get(ctx, collabId)
	checkPending, ok := CollabReq.MembersStatus[username]
	if !ok {
		return response, errors.New("Collab not supported for this username")
	} else if checkPending == "accept" || checkPending == "reject" {
		return response, errors.New(fmt.Sprintf("Collab already %sed", checkPending))
	} else {
		CollabReq.MembersStatus[username] = "accept"
	}
	repo := repository.GetCollabRepo()
	memberstatusJson, err := json.Marshal(CollabReq.MembersStatus)
	if err != nil {
		return response, err
	}
	attributes := map[string]interface{}{
		"members_status": memberstatusJson,
	}
	err = repo.Update(ctx, &resp, attributes, map[string]interface{}{
		"id": collabId,
	})
	l.ResponseMapper(resp, &response)
	if err != nil {
		return response, err
	}
	return response, iError
}

func (l collabProcessor) RejectCollab(ctx *gin.Context, collabId int, username string) (dtos.CollabResponse, error) {
	var response dtos.CollabResponse
	var resp dtos.Collab
	CollabReq, iError := l.Get(ctx, collabId)
	checkPending, ok := CollabReq.MembersStatus[username]
	if !ok {
		return response, errors.New("Collab not supported for this username")
	} else if checkPending == "reject" {
		return response, errors.New(fmt.Sprintf("Collab already %sed", checkPending))
	} else {
		CollabReq.MembersStatus[username] = "reject"
	}
	repo := repository.GetCollabRepo()
	memberstatusJson, err := json.Marshal(CollabReq.MembersStatus)
	if err != nil {
		return response, err
	}
	attributes := map[string]interface{}{
		"members_status": memberstatusJson,
	}
	err = repo.Update(ctx, &resp, attributes, map[string]interface{}{
		"id": collabId,
	})
	l.ResponseMapper(resp, &response)
	if err != nil {
		return response, err
	}
	return response, iError
}

func (l collabProcessor) Update(ctx *gin.Context, request dtos.CollabRequest, userId int) (dtos.CollabResponse, error) {
	//var about dtos.About
	var response dtos.CollabResponse
	//repo := repository.GetAboutRepo()
	//attributes := map[string]interface{}{
	//	"about": request.About,
	//}
	//err := repo.Update(ctx, &about, attributes, map[string]interface{}{
	//	"user_id": userId,
	//})
	//l.ResponseMapper(about, &response)
	//if err != nil {
	//	return response, err
	//}
	//response.Success = true
	return response, nil
}

func (l collabProcessor) Delete(ctx *gin.Context, userId int) error {
	var response dtos.About
	repo := repository.GetAboutRepo()
	err := repo.Delete(ctx, &response, map[string]interface{}{
		"user_id": userId,
	})
	if err != nil {
		return err
	}
	return nil
}

func GetCollabProcessor() interfaces.ICollab {
	return collabProcessor{}
}

func (l collabProcessor) Mapper(request dtos.CollabRequest, response *dtos.Collab) {
	var err error
	response.CreatedBy = request.CreatedBy
	response.Members, err = json.Marshal(request.Members)
	response.ProjectTitle = request.ProjectTitle
	response.Videolink = request.Videolink
	response.MembersStatus, err = json.Marshal(request.MembersStatus)
	fmt.Println(err)
}

func (l collabProcessor) ResponseMapper(request dtos.Collab, response *dtos.CollabResponse) {
	response.Videolink = request.Videolink
	response.ProjectTitle = request.ProjectTitle
	response.CreatedBy = request.CreatedBy
	err := json.Unmarshal(request.Members, &response.Members)
	fmt.Println(err)
	err = json.Unmarshal(request.MembersStatus, &response.MembersStatus)
	fmt.Println(err)
}

func (l collabProcessor) MemberStatusDefault(request *dtos.CollabRequest) {
	request.MembersStatus = make(map[string]string)
	for _, member := range request.Members {
		request.MembersStatus[member] = "pending"
	}
}
