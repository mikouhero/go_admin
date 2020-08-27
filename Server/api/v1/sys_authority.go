package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/Server/global/response"
	"go_admin/Server/model"
	"go_admin/Server/model/request"
	response2 "go_admin/Server/model/response"
	"go_admin/Server/service"
	"go_admin/Server/utils"
)

// 创建角色
func CreateAuthority(c *gin.Context) {
	var auth model.SysAuthority
	_ = c.ShouldBindJSON(&auth)

	rules := utils.Rules{
		"AuthorityId":   {utils.NotEmpty()},
		"AuthorityName": {utils.NotEmpty()},
		"ParentId":      {utils.NotEmpty()},
	}

	err := utils.Verify(auth, rules)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	err, authority := service.CreateAuthority(auth)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	} else {
		response.OkWithData(response2.SysAuthorityResponse{Authority: authority}, c)
	}

}

// 删除角色
func DeleteAuthority(c *gin.Context) {
	var a model.SysAuthority
	_ = c.ShouldBindJSON(&a)
	AuthorityIdVerifyErr := utils.Verify(a, utils.CustomizeMap["AuthorityIdVerify"])
	if AuthorityIdVerifyErr != nil {
		response.FailWithMsg(AuthorityIdVerifyErr.Error(), c)
		return
	}
	err := service.DeleteAuthority(&a)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	} else {
		response.Ok(c)
	}
}

// 更新角色
func UpdateAuthority(c *gin.Context) {
	var auth model.SysAuthority
	_ = c.ShouldBindJSON(&auth)
	rules := utils.Rules{
		"AuthorityId":   {utils.NotEmpty()},
		"AuthorityName": {utils.NotEmpty()},
		"ParentId":      {utils.NotEmpty()},
	}
	err := utils.Verify(auth, rules)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	err, authority := service.UpdateAuthority(auth)

	if err != nil {
		response.FailWithMsg(fmt.Sprintf("更新失败 ,%v", err), c)
		return
	} else {
		response.OkWithData(response2.SysAuthorityResponse{Authority: authority}, c)
	}
}

func CopyAuthority(c *gin.Context) {

}

// 分页获取角色列表
func GetAuthorityList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	err, list, totle := service.GetAuthorityInfoList(pageInfo)

	if err != nil {
		response.FailWithMsg(err.Error(), c)
	} else {

		response.OkWithData(response2.PageResult{
			List:     list,
			Totle:    totle,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}

// 设置资源权限
func SetDataAuthority(c *gin.Context) {

}
