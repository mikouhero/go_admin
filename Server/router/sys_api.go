package router

import (
	"github.com/gin-gonic/gin"
	v1 "go_admin/Server/api/v1"
	"go_admin/Server/middleware"
)

func InitApiRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("api").
		Use(middleware.JWTAuth()).
		Use()
	{
		ApiRouter.POST("createApi", v1.CreateApi)
		ApiRouter.POST("deleteApi", v1.DeleteApi)
		ApiRouter.POST("getApiList", v1.GetApiList)
		ApiRouter.POST("getApiById", v1.GetApiById)
		ApiRouter.POST("updateApi", v1.UpdateApi)
		ApiRouter.POST("getAllApis", v1.GetAllApis)

	}

}
