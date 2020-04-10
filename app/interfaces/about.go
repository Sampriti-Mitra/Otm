package interfaces

import (
	"github.com/gin-gonic/gin"
	"otm/app/dtos"
)

type IAbout interface {
	Create(ctx *gin.Context, request dtos.AboutRequest) (dtos.AboutResponse, error)
	Get(ctx *gin.Context, requestId int) (dtos.AboutResponse, error)
	Update(ctx *gin.Context, request dtos.AboutRequest, registrationId int) (dtos.AboutResponse, error)
	Delete(ctx *gin.Context, userId int) error
}
