package initialiaze

import (
	"github.com/gin-gonic/gin"
	"go_admin/Server/middleware"
	"go_admin/Server/router"
)

// 初始化总路由
func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.Use(middleware.Cors())
	ApiGroup := Router.Group("")
	router.InitUserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)
	router.InitMenuRouter(ApiGroup)
	router.InitAuthority(ApiGroup)
	router.InitApiRouter(ApiGroup)
	router.InitSysOperationRecord(ApiGroup)

	return Router
}
