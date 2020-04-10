package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type IDbProvider interface {
	Instance() *gorm.DB
	Ping(ctx *gin.Context) error
}
