package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type FsRouter struct{}

func (* FsRouter) InitFsRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	fsRouter := Router.Group("fs")
	fsApi := v1.ApiGroupApp.SystemApiGroup.FsApi
	{
		fsRouter.POST("login", fsApi.FsLogin)
	}
	return fsRouter
}