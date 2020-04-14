package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type IBase interface {
	RunQuery(ctx *gin.Context, query string, arguments []interface{}) error
	Find(ctx *gin.Context, model interface{}, condition map[string]interface{}) error
	FindAll(ctx *gin.Context, model interface{}, condition map[string]interface{}) error
	FindTrending(ctx *gin.Context, model interface{}, condition map[string]interface{}) error
	Update(ctx *gin.Context, model interface{}, attributes map[string]interface{}, condition map[string]interface{}) error
	Create(ctx *gin.Context, model interface{}) error
	Delete(ctx *gin.Context, model interface{}, condition map[string]interface{}) error
	Transaction(ctx *gin.Context, handlers ...func() (interface{}, error)) (data interface{}, ierr error)
	GetDbInstance(ctx *gin.Context) *gorm.DB
}

type ILoginRepo interface {
	IBase
}

type IProfileRepo interface {
	IBase
}

type IFollowerRepo interface {
	IBase
}

type ICollabRepo interface {
	IBase
}
