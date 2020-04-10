package interfaces

import (
	"github.com/gin-gonic/gin"
	"otm/app/dtos"
)

type ILogin interface {
	Create(ctx *gin.Context, request dtos.RegisterRequest) (dtos.RegisterResponse, error)
	Get(ctx *gin.Context, requestId int) (dtos.RegisterResponse, error)
	Update(ctx *gin.Context, request dtos.RegisterRequest, registrationId string) (dtos.RegisterResponse, error)
	Delete(ctx *gin.Context, requestId string) dtos.DeletedResponse
	GetByUsername(ctx *gin.Context, username string) (dtos.Users, error)
}
