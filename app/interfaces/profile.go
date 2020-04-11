package interfaces

import (
	"github.com/gin-gonic/gin"
	"otm/app/dtos"
)

type IProfile interface {
	Create(ctx *gin.Context, request dtos.UploadRequest) (dtos.UploadResponse, error)
	List(ctx *gin.Context, requestId int) ([]dtos.UploadResponse, error)
	ListFeed(ctx *gin.Context, requestId int) ([]dtos.UploadResponse, error)
	ListTrending(ctx *gin.Context) ([]dtos.UploadResponse, error)
	Get(ctx *gin.Context, videoId int) (dtos.UploadResponse, error)
	Update(ctx *gin.Context, request dtos.UploadRequest, registrationId int) (dtos.UploadResponse, error)
	Delete(ctx *gin.Context, requestId string) dtos.DeletedResponse
	GetByUsername(ctx *gin.Context, username string) (dtos.UploadResponse, error)
	Applaud(ctx *gin.Context, requestedBy string, videoId int) (dtos.Applaud, error)
}
