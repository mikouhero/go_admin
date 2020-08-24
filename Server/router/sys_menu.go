package router

import (
	"github.com/gin-gonic/gin"
	v1 "go_admin/Server/api/v1"
)

func InitMenuRouter(Router *gin.RouterGroup) {
	MenuRouter := Router.Group("menu").
		//Use(middleware.JWTAuth()).
		Use()
	{
		MenuRouter.POST("addBaseMenu", v1.AddBaseMenu)
		MenuRouter.POST("updateBaseMenu", v1.UpdateBaseMenu)
		MenuRouter.POST("deleteBaseMenu", v1.DeleteBaseMenu)
		MenuRouter.POST("getBaseMenuById", v1.GetBaseMenuById)
	}

}
