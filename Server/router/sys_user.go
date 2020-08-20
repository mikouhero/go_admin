package router

import (
	"github.com/gin-gonic/gin"
	v1 "go_admin/Server/api/v1"
	"go_admin/Server/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").
		Use(middleware.JWTAuth()).
		Use()
	{
		UserRouter.POST("changePassword", v1.ChangePassword)     // 修改密码
		UserRouter.POST("uploadHeaderImg",v1.UploadHeaderImg)    // 上传头像
		UserRouter.POST("getUserList",v1.GetUserList)  //  分页获取用户列表
	}
}
