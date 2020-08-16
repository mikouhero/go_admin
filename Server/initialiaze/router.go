package initialiaze

import (
	"github.com/gin-gonic/gin"
	"go_admin/Server/router"
)

// 初始化总路由
func Routers() *gin.Engine {
	var Router = gin.Default()

	ApiGroup := Router.Group("")
	router.InitUserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)

	return Router
}
