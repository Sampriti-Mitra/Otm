package internals

import (
	"github.com/gin-gonic/gin"
	"otm/app/dtos"
	"otm/app/interfaces"
	"otm/app/repository"
)

type aboutProcessor struct {
}

func (l aboutProcessor) Create(ctx *gin.Context, request dtos.AboutRequest) (dtos.AboutResponse, error) {
	var response dtos.AboutResponse
	var resp dtos.About
	repo := repository.GetAboutRepo()
	l.Mapper(request, &resp)
	err := repo.Create(ctx, &resp)
	if err != nil {
		return response, err
	}
	l.ResponseMapper(resp, &response)
	response.Success = true
	return response, nil
}

func (l aboutProcessor) Get(ctx *gin.Context, requestId int) (dtos.AboutResponse, error) {
	var about dtos.About
	var response dtos.AboutResponse
	repo := repository.GetAboutRepo()
	err := repo.Find(ctx, &about, map[string]interface{}{
		"user_id": requestId,
	})
	l.ResponseMapper(about, &response)
	return response, err
}

func (l aboutProcessor) Update(ctx *gin.Context, request dtos.AboutRequest, userId int) (dtos.AboutResponse, error) {
	var about dtos.About
	var response dtos.AboutResponse
	repo := repository.GetAboutRepo()
	attributes := map[string]interface{}{
		"about": request.About,
	}
	err := repo.Update(ctx, &about, attributes, map[string]interface{}{
		"user_id": userId,
	})
	l.ResponseMapper(about, &response)
	if err != nil {
		return response, err
	}
	response.Success = true
	return response, nil
}

func (l aboutProcessor) Delete(ctx *gin.Context, userId int) error {
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

func GetAboutProcessor() interfaces.IAbout {
	return aboutProcessor{}
}

func (l aboutProcessor) Mapper(request dtos.AboutRequest, response *dtos.About) {
	response.UserId = request.UserId
	response.About = request.About
}

func (l aboutProcessor) ResponseMapper(request dtos.About, response *dtos.AboutResponse) {
	response.UserId = request.UserId
	response.About = request.About
	response.Model = request.Model
}
