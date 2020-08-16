package router

import (
	"github.com/gin-gonic/gin"
	v1 "go_admin/Server/api/v1"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("login", v1.Login)
		BaseRouter.POST("register", v1.Register)
	}
	return BaseRouter
}
