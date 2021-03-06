package v1

import (
	"encoding/json"
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
	//claims, _ := c.Get("claims")
	//customClaims := claims.(*request.CustomClaims)
	//service.GetMenuTree(customClaims.AuthorityId)
	data := `{"menus":[{"ID":0,"CreatedAt":"2019-09-19T22:05:18+08:00","UpdatedAt":"2020-05-30T15:43:06+08:00","DeletedAt":null,"parentId":"0","path":"dashboard","name":"dashboard","hidden":false,"component":"view/dashboard/index.vue","sort":1,"meta":{"keepAlive":false,"defaultMenu":false,"title":"仪表盘","icon":"setting"},"authoritys":null,"menuId":"1","children":null},{"ID":0,"CreatedAt":"2019-09-19T22:06:38+08:00","UpdatedAt":"2020-04-24T10:16:43+08:00","DeletedAt":null,"parentId":"0","path":"admin","name":"superAdmin","hidden":false,"component":"view/superAdmin/index.vue","sort":3,"meta":{"keepAlive":false,"defaultMenu":false,"title":"超级管理员","icon":"user-solid"},"authoritys":null,"menuId":"3","children":[{"ID":0,"CreatedAt":"2019-09-19T22:11:53+08:00","UpdatedAt":"2020-05-30T15:43:25+08:00","DeletedAt":null,"parentId":"3","path":"authority","name":"authority","hidden":false,"component":"view/superAdmin/authority/authority.vue","sort":1,"meta":{"keepAlive":false,"defaultMenu":false,"title":"角色管理","icon":"s-custom"},"authoritys":null,"menuId":"4","children":null},{"ID":0,"CreatedAt":"2020-06-24T19:51:33+08:00","UpdatedAt":"2020-06-28T20:35:04+08:00","DeletedAt":null,"parentId":"3","path":"dictionaryDetail/:id","name":"dictionaryDetail","hidden":true,"component":"view/superAdmin/dictionary/sysDictionaryDetail.vue","sort":1,"meta":{"keepAlive":false,"defaultMenu":false,"title":"字典详情","icon":"s-order"},"authoritys":null,"menuId":"51","children":null},{"ID":0,"CreatedAt":"2019-09-19T22:13:18+08:00","UpdatedAt":"2020-04-30T17:45:27+08:00","DeletedAt":null,"parentId":"3","path":"menu","name":"menu","hidden":false,"component":"view/superAdmin/menu/menu.vue","sort":2,"meta":{"keepAlive":true,"defaultMenu":false,"title":"菜单管理","icon":"s-order"},"authoritys":null,"menuId":"5","children":null},{"ID":0,"CreatedAt":"2019-09-19T22:13:36+08:00","UpdatedAt":"2020-04-24T10:16:43+08:00","DeletedAt":null,"parentId":"3","path":"api","name":"api","hidden":false,"component":"view/superAdmin/api/api.vue","sort":3,"meta":{"keepAlive":true,"defaultMenu":false,"title":"api管理","icon":"s-platform"},"authoritys":null,"menuId":"6","children":null},{"ID":0,"CreatedAt":"2019-10-09T15:12:29+08:00","UpdatedAt":"2020-04-24T10:16:43+08:00","DeletedAt":null,"parentId":"3","path":"user","name":"user","hidden":false,"component":"view/superAdmin/user/user.vue","sort":4,"meta":{"keepAlive":false,"defaultMenu":false,"title":"用户管理","icon":"coordinate"},"authoritys":null,"menuId":"17","children":null},{"ID":0,"CreatedAt":"2020-06-24T19:49:54+08:00","UpdatedAt":"2020-06-28T20:34:47+08:00","DeletedAt":null,"parentId":"3","path":"dictionary","name":"dictionary","hidden":false,"component":"view/superAdmin/dictionary/sysDictionary.vue","sort":5,"meta":{"keepAlive":false,"defaultMenu":false,"title":"字典管理","icon":"notebook-2"},"authoritys":null,"menuId":"50","children":null},{"ID":0,"CreatedAt":"2020-06-29T13:31:17+08:00","UpdatedAt":"2020-07-07T16:05:34+08:00","DeletedAt":null,"parentId":"3","path":"operation","name":"operation","hidden":false,"component":"view/superAdmin/operation/sysOperationRecord.vue","sort":6,"meta":{"keepAlive":false,"defaultMenu":false,"title":"操作历史","icon":"time"},"authoritys":null,"menuId":"52","children":null}]},{"ID":0,"CreatedAt":"2019-10-15T22:27:22+08:00","UpdatedAt":"2020-05-10T21:31:36+08:00","DeletedAt":null,"parentId":"0","path":"person","name":"person","hidden":true,"component":"view/person/person.vue","sort":4,"meta":{"keepAlive":false,"defaultMenu":false,"title":"个人信息","icon":"message-solid"},"authoritys":null,"menuId":"18","children":null},{"ID":0,"CreatedAt":"2020-03-29T21:31:03+08:00","UpdatedAt":"2020-04-24T10:16:43+08:00","DeletedAt":null,"parentId":"0","path":"systemTools","name":"systemTools","hidden":false,"component":"view/systemTools/index.vue","sort":5,"meta":{"keepAlive":false,"defaultMenu":false,"title":"系统工具","icon":"s-cooperation"},"authoritys":null,"menuId":"38","children":[{"ID":0,"CreatedAt":"2020-03-29T21:35:10+08:00","UpdatedAt":"2020-05-03T21:38:49+08:00","DeletedAt":null,"parentId":"38","path":"autoCode","name":"autoCode","hidden":false,"component":"view/systemTools/autoCode/index.vue","sort":1,"meta":{"keepAlive":true,"defaultMenu":false,"title":"代码生成器","icon":"cpu"},"authoritys":null,"menuId":"40","children":null},{"ID":0,"CreatedAt":"2020-03-29T21:36:26+08:00","UpdatedAt":"2020-05-03T21:38:43+08:00","DeletedAt":null,"parentId":"38","path":"formCreate","name":"formCreate","hidden":false,"component":"view/systemTools/formCreate/index.vue","sort":2,"meta":{"keepAlive":true,"defaultMenu":false,"title":"表单生成器","icon":"magic-stick"},"authoritys":null,"menuId":"41","children":null}]},{"ID":0,"CreatedAt":"2019-10-20T11:14:42+08:00","UpdatedAt":"2020-04-24T10:16:43+08:00","DeletedAt":null,"parentId":"0","path":"example","name":"example","hidden":false,"component":"view/example/index.vue","sort":6,"meta":{"keepAlive":false,"defaultMenu":false,"title":"示例文件","icon":"s-management"},"authoritys":null,"menuId":"19","children":[{"ID":0,"CreatedAt":"2019-10-20T11:18:11+08:00","UpdatedAt":"2020-04-24T10:16:42+08:00","DeletedAt":null,"parentId":"19","path":"table","name":"table","hidden":false,"component":"view/example/table/table.vue","sort":1,"meta":{"keepAlive":false,"defaultMenu":false,"title":"表格示例","icon":"s-order"},"authoritys":null,"menuId":"20","children":null},{"ID":0,"CreatedAt":"2019-10-20T11:19:52+08:00","UpdatedAt":"2020-04-24T10:16:43+08:00","DeletedAt":null,"parentId":"19","path":"form","name":"form","hidden":false,"component":"view/example/form/form.vue","sort":2,"meta":{"keepAlive":false,"defaultMenu":false,"title":"表单示例","icon":"document"},"authoritys":null,"menuId":"21","children":null},{"ID":0,"CreatedAt":"2019-10-20T11:22:19+08:00","UpdatedAt":"2020-04-24T10:16:43+08:00","DeletedAt":null,"parentId":"19","path":"rte","name":"rte","hidden":false,"component":"view/example/rte/rte.vue","sort":3,"meta":{"keepAlive":false,"defaultMenu":false,"title":"富文本编辑器","icon":"reading"},"authoritys":null,"menuId":"22","children":null},{"ID":0,"CreatedAt":"2019-10-20T11:23:39+08:00","UpdatedAt":"2020-04-24T10:16:43+08:00","DeletedAt":null,"parentId":"19","path":"excel","name":"excel","hidden":false,"component":"view/example/excel/excel.vue","sort":4,"meta":{"keepAlive":false,"defaultMenu":false,"title":"excel导入导出","icon":"s-marketing"},"authoritys":null,"menuId":"23","children":null},{"ID":0,"CreatedAt":"2019-10-20T11:27:02+08:00","UpdatedAt":"2020-04-24T10:16:43+08:00","DeletedAt":null,"parentId":"19","path":"upload","name":"upload","hidden":false,"component":"view/example/upload/upload.vue","sort":5,"meta":{"keepAlive":false,"defaultMenu":false,"title":"上传下载","icon":"upload"},"authoritys":null,"menuId":"26","children":null},{"ID":0,"CreatedAt":"2020-02-17T16:20:47+08:00","UpdatedAt":"2020-04-24T10:16:43+08:00","DeletedAt":null,"parentId":"19","path":"breakpoint","name":"breakpoint","hidden":false,"component":"view/example/breakpoint/breakpoint.vue","sort":6,"meta":{"keepAlive":false,"defaultMenu":false,"title":"断点续传","icon":"upload"},"authoritys":null,"menuId":"33","children":null},{"ID":0,"CreatedAt":"2020-02-24T19:48:37+08:00","UpdatedAt":"2020-04-24T10:16:43+08:00","DeletedAt":null,"parentId":"19","path":"customer","name":"customer","hidden":false,"component":"view/example/customer/customer.vue","sort":7,"meta":{"keepAlive":false,"defaultMenu":false,"title":"客户列表（资源示例）","icon":"s-custom"},"authoritys":null,"menuId":"34","children":null}]},{"ID":0,"CreatedAt":"2019-09-19T22:06:17+08:00","UpdatedAt":"2020-05-10T21:31:50+08:00","DeletedAt":null,"parentId":"0","path":"about","name":"about","hidden":false,"component":"view/about/index.vue","sort":7,"meta":{"keepAlive":false,"defaultMenu":false,"title":"关于我们","icon":"info"},"authoritys":null,"menuId":"2","children":null}]}`
	var mapResult map[string]interface{}

	err := json.Unmarshal([]byte(data), &mapResult)
	fmt.Println(err)
	response.OkWithData(mapResult, c)
}

func GetMenuList(c *gin.Context) {

}

func GetBaseMenuTree(c *gin.Context) {
	err, menus := service.GetBaseMenuTree()
	if err != nil {
		response.FailWithMsg("获取失败"+err.Error(), c)
	} else {
		response.OkWithData(response2.SysBaseMenusResponse{Menus: menus}, c)
	}
}

func AddMenuAuthority(c *gin.Context) {

}

// 获取指定角色的菜单
func GetMenuAuthority(c *gin.Context) {
	var authorityIdInfo request.AddMenuAuthorityInfo
	_ = c.ShouldBindJSON(&authorityIdInfo)
	menuVerify := utils.Rules{
		"AuthorityId": {utils.NotEmpty()},
	}
	err := utils.Verify(authorityIdInfo, menuVerify)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	err, menus := service.GetMenuAuthority(authorityIdInfo.AuthorityId)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
	} else {
		response.OkWithData(menus, c)
	}

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
