package router

import (
	"github.com/gin-gonic/gin"
	v1 "go_admin/Server/api/v1"
	"go_admin/Server/middleware"
)

func InitMenuRouter(Router *gin.RouterGroup) {
	MenuRouter := Router.Group("menu").
		Use(middleware.JWTAuth()).
		Use()
	{
		MenuRouter.POST("addBaseMenu", v1.AddBaseMenu)
		MenuRouter.POST("updateBaseMenu", v1.UpdateBaseMenu)
		MenuRouter.POST("deleteBaseMenu", v1.DeleteBaseMenu)
		MenuRouter.POST("getBaseMenuById", v1.GetBaseMenuById)
		MenuRouter.POST("getMenu", v1.GetMenu)
		MenuRouter.POST("getBaseMenuTree", v1.GetBaseMenuTree)   // 获取用户动态路由
		MenuRouter.POST("addMenuAuthority", v1.AddMenuAuthority) //	增加menu和角色关联关系
		MenuRouter.POST("getMenuAuthority", v1.GetMenuAuthority) // 获取指定角色menu

	}

}
