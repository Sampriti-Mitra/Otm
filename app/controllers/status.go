package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"otm/app/constants"
	"otm/app/dtos"
	"otm/app/providers/database"
)

// statusController : struct to keep the status controller
type statusController struct {
	baseController
}

var (
	// Status will hold the instance of statusController
	Status           statusController
	databaseInstance = database.GetClient
	// AppRunningStatus : application running status
	AppRunningStatus = true
)

// NewStatusController : function to initialize new status controller instance
func NewStatusController() {
	Status = statusController{}
}

// Status will provide the status of the service we been used in the application
func (controller statusController) Status(ctx *gin.Context) {
	dbErr := databaseInstance().Ping(ctx)
	if dbErr == nil && AppRunningStatus {
		resp := dtos.Base{}
		resp.Success = true
		ctx.JSON(http.StatusOK, resp)
	} else {
		errResult := dtos.Base{}
		errResult.Success = false
		errResult.Error = dbErr
		//if dbErr != nil {
		//		//	if errResult.Error != nil {
		//		//		errResult.Error = errResult.Error + ","
		//		//	}
		//		//	errResult.Error = errResult.Error + dbErr.ErrorCode()
		//		//}
		ctx.JSON(http.StatusServiceUnavailable, errResult)
	}

}

func (controller statusController) Ping(ctx *gin.Context) {

	result := make(map[string]interface{})

	result[constants.CommitID] = os.Getenv(constants.GitCommitHash)
	result[constants.ContainerID] = os.Getenv(constants.Hostname)

	ctx.JSON(http.StatusOK, result)
}
