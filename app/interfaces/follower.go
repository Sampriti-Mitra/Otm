package interfaces

import (
	"github.com/gin-gonic/gin"
	"otm/app/dtos"
)

type IFollower interface {
	Create(ctx *gin.Context, request dtos.FollowRequest) (dtos.FollowResponse, error)
	List(ctx *gin.Context, requestTo string) ([]dtos.FollowResponse, error)
	Update(ctx *gin.Context, request dtos.FollowRequest, registrationId string) (dtos.FollowResponse, error)
	Delete(ctx *gin.Context, requestId string) dtos.DeletedResponse
}
