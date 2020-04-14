package interfaces

import (
	"github.com/gin-gonic/gin"
	"otm/app/dtos"
)

type ICollab interface {
	Create(ctx *gin.Context, request dtos.CollabRequest) (dtos.CollabResponse, error)
	Get(ctx *gin.Context, collabId int) (dtos.CollabResponse, error)
	Update(ctx *gin.Context, request dtos.CollabRequest, userId int) (dtos.CollabResponse, error)
	AcceptCollab(ctx *gin.Context, collabId int, username string) (dtos.CollabResponse, error)
	RejectCollab(ctx *gin.Context, collabId int, username string) (dtos.CollabResponse, error)
	Delete(ctx *gin.Context, userId int) error
}
