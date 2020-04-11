package interfaces

import (
	"github.com/gin-gonic/gin"
	"otm/app/dtos"
)

type IFollower interface {
	Create(ctx *gin.Context, request dtos.FollowRequest) (dtos.FollowResponse, error)
	List(ctx *gin.Context, requestTo string) ([]dtos.FollowResponse, error)
	UpdateAccept(ctx *gin.Context, profileId string, request_by string) (dtos.FollowResponse, error)
	UpdateReject(ctx *gin.Context, profileId string, request_by string) (dtos.FollowResponse, error)
	Delete(ctx *gin.Context, requestId string) dtos.DeletedResponse
	Get(ctx *gin.Context, userId string, requestBy string) (dtos.FollowResponse, error)
}
