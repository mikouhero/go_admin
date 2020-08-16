package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "go_admin/Server/api/v1"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use()
	{
		UserRouter.GET("msg", func(context *gin.Context) {
			fmt.Println("user msg")
		})
		UserRouter.POST("register", v1.Register)
	}
}
