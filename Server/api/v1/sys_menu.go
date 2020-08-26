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

func GetMenu(c *gin.Context) {

}

func GetMenuList(c *gin.Context) {

}

func GetBaseMenuTree(c *gin.Context) {

}

func AddMenuAuthority(c *gin.Context) {

}

func GetMenuAuthority(c *gin.Context) {

}

// 添加菜单
func AddBaseMenu(c *gin.Context) {
	var menu model.SysBaseMenu
	_ = c.ShouldBindJSON(&menu)
	rules := utils.Rules{
		"Path":      {utils.NotEmpty()},
		"ParentId":  {utils.NotEmpty()},
		"Name":      {utils.NotEmpty()},
		"Component": {utils.NotEmpty()},
		"Sort":      {utils.NotEmpty()},
	}
	fmt.Println(menu)
	err := utils.Verify(menu, rules)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	MetaVertify := utils.Rules{
		"Title": {utils.NotEmpty()},
	}

	err = utils.Verify(menu.Meta, MetaVertify)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	err = service.AddBaseMenu(menu)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("添加菜单失败 %v", err), c)
	} else {
		response.OkWithMsg("success", c)
	}

}

// 删除菜单
func DeleteBaseMenu(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	IdVerityErr := utils.Verify(idInfo, utils.CustomizeMap["IdVerity"])
	if IdVerityErr != nil {
		response.FailWithMsg(IdVerityErr.Error(), c)
		return
	}
	err := service.DeleteBaseMenu(idInfo.Id)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("删除失败 %v", err), c)
	} else {
		response.OkWithMsg("success", c)
	}
}

// 更新菜单
func UpdateBaseMenu(c *gin.Context) {
	var menu model.SysBaseMenu
	_ = c.ShouldBindJSON(&menu)

	rules := utils.Rules{
		"Path":      {utils.NotEmpty()},
		"ID":        {utils.NotEmpty()},
		"Name":      {utils.NotEmpty()},
		"Component": {utils.NotEmpty()},
		"Sort":      {utils.Ge("0"), "ge=0"},
	}
	fmt.Println(menu)
	verifyErr := utils.Verify(menu, rules)
	if verifyErr != nil {
		response.FailWithMsg(verifyErr.Error(), c)
		return
	}

	metaRule := utils.Rules{
		"Title": {utils.NotEmpty()},
	}
	verifyErr = utils.Verify(menu.Meta, metaRule)
	if verifyErr != nil {
		response.FailWithMsg(verifyErr.Error(), c)
		return
	}

	err := service.UpdateBaseMenu(menu)
	if err != nil {
		response.FailWithMsg(verifyErr.Error(), c)
	} else {
		response.OkWithMsg("success", c)
	}

}

// 通过ID获取菜单信息
func GetBaseMenuById(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	IdVerityErr := utils.Verify(idInfo, utils.CustomizeMap["IdVerity"])
	if IdVerityErr != nil {
		response.FailWithMsg(IdVerityErr.Error(), c)
		return
	}

	err, menu := service.GetBaseMenuById(idInfo.Id)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	} else {
		response.OkWithData(response2.SysBaseMenuResponse{Menu: menu}, c)
	}
}
