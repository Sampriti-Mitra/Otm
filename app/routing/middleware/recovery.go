package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"otm/app/dtos"
	//"otm/app/rzperrors"
	//"otm/app/utils"
	"net/http"
)

// Recovery : Middlewares to catch uncaught rzperror
// This will catch all un-caught panics (if any)
// In case of panic default error response will be set
func Recovery(ctx *gin.Context) {
	defer func(ctx *gin.Context) {
		if rec := recover(); rec != nil {
			//err := utils.GetError(rec)

			//errorCode := rzperrors.ErrorInternalServerError
			//payloadCreationError := rzperrors.NewRzpError(ctx, errorCode, err)
			//logger.RzpError(ctx, payloadCreationError)

			response := dtos.Base{}
			response.Success = false
			//response.Error = payloadCreationError.ErrorCode()
			response.Error = errors.New("error")
			ctx.JSON(http.StatusInternalServerError, response)
		}
	}(ctx)

	ctx.Next()
}
