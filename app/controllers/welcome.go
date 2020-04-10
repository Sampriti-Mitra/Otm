package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type appController struct {
	baseController
}

var App appController

func NewAppController() {
	App = appController{}
}

func (app appController) Welcome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Message": "Welcome to OneTouchMusic"})
}
