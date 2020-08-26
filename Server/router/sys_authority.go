package router

import (
	"github.com/gin-gonic/gin"
	v1 "go_admin/Server/api/v1"
)

func InitAuthority(Router *gin.RouterGroup) {
	AuthorityRouter := Router.Group("authority").
		//Use(middleware.JWTAuth()).
		Use()
	{
		AuthorityRouter.POST("createAuthority", v1.CreateAuthority)   // 创建角色
		AuthorityRouter.POST("deleteAuthority", v1.DeleteAuthority)   // 删除角色
		AuthorityRouter.POST("updateAuthority", v1.UpdateAuthority)    // 更新角色
		AuthorityRouter.POST("copyAuthority", v1.CopyAuthority)       // 更新角色
		AuthorityRouter.POST("getAuthorityList", v1.GetAuthorityList) // 获取角色列表
		AuthorityRouter.POST("setDataAuthority", v1.SetDataAuthority) // 设置角色资源权限
	}
}
