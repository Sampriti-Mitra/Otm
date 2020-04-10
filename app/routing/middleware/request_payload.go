package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"otm/app/constants"
)

func SetRequestPayload() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//TODO: This is added temporarily
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		traceData := getRequestDetails(ctx)

		traceData[constants.Mode] = ctx.GetString(constants.Mode)
		traceData[constants.RequestID] = generateUUID()
		traceData[constants.TaskID] = getTaskId(ctx)
		ctx.Set(constants.Request, traceData)
	}
}

func getRequestDetails(ctx *gin.Context) map[string]interface{} {
	data := map[string]interface{}{}

	if ctx.Request == nil {
		return data
	}

	data["uri"] = ctx.Request.RequestURI
	data["host"] = ctx.Request.Host
	data["referer"] = ctx.Request.Referer()

	return data
}

func getTaskId(ctx *gin.Context) string {
	taskId := ctx.Request.Header.Get(constants.X_RAZORPAY_TASK_ID)

	if len(taskId) == 0 {
		taskId = generateUUID()
	}
	return taskId
}

func generateUUID() string {
	uui, _ := uuid.NewV4()
	return uuid.NewV5(uui, "name").String()
}
