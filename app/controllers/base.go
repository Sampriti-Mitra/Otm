package controllers

import (
	"github.com/gin-gonic/gin"
	"otm/app/dtos"
)

type baseController struct {
}

func (b baseController) setErrorResponse(ctx *gin.Context, err error) {
	var errorResponse dtos.ErrorResponse

	errorResponse.Success = false
	message := err.Error()
	errorResponse.Error = dtos.Error{err, message}
	ctx.JSON(500, errorResponse)
}

func (b baseController) deserialize(ctx *gin.Context, target interface{}) error {
	var err error

	err = ctx.ShouldBindJSON(target)
	return err
}
