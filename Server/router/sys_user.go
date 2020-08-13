package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use()
	{
		UserRouter.GET("msg", func(context *gin.Context) {
			fmt.Println("user msg")
		})
	}
}
