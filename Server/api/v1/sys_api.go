package v1

import (
	"github.com/gin-gonic/gin"
	"go_admin/Server/global/response"
	"go_admin/Server/model"
	"go_admin/Server/model/request"
	response2 "go_admin/Server/model/response"
	"go_admin/Server/service"
	"go_admin/Server/utils"
)

func CreateApi(c *gin.Context) {
	var api model.SysApi
	_ = c.ShouldBindJSON(&api)

	ApiVerity := utils.Rules{
		"Path":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"ApiGroup":    {utils.NotEmpty()},
		"Method":      {utils.NotEmpty()},
	}

	err := utils.Verify(api, ApiVerity)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	err = service.CreateApi(api)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
	} else {
		response.Ok(c)
	}

}

func DeleteApi(c *gin.Context) {

	var a model.SysApi
	_ = c.ShouldBindJSON(&a)
	ApiVerity := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	err := utils.Verify(a.Model, ApiVerity)

	if err != nil {
		response.FailWithMsg(err.Error(), c)
	}

	err = service.DeleteApi(a)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
	} else {
		response.Ok(c)
	}

}

func GetApiList(c *gin.Context) {
	var sp request.SearchApiParams
	_ = c.ShouldBindJSON(&sp)
	PageVerifyErr := utils.Verify(sp.PageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMsg(PageVerifyErr.Error(), c)
		return
	}
	PageVerifyErr, list, total := service.GetApiInfoList(sp.SysApi, sp.PageInfo, sp.OrderKey, sp.Desc)

	if PageVerifyErr != nil {
		response.FailWithMsg(PageVerifyErr.Error(), c)
	} else {
		response.OkWithData(response2.PageResult{
			List:     list,
			Totle:    total,
			Page:     sp.Page,
			PageSize: sp.PageSize,
		}, c)
	}

}
func GetApiById(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	err := utils.Verify(idInfo, utils.CustomizeMap["IdVerify"])
	if err != nil {
		response.FailWithMsg(err.Error(), c)
	}
	err, api := service.GetApiById(idInfo.Id)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
	} else {
		response.OkWithData(response2.SysAPIResponse{Api: api}, c)
	}
}

func UpdateApi(c *gin.Context) {
	var api model.SysApi
	_ = c.ShouldBindJSON(&api)

	ApiVerity := utils.Rules{
		"ID":          {utils.NotEmpty()},
		"Path":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"ApiGroup":    {utils.NotEmpty()},
		"Method":      {utils.NotEmpty()},
	}

	err := utils.Verify(api, ApiVerity)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	err = service.UpdateApi(api)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
	} else {
		response.Ok(c)
	}
}

func GetAllApis(c *gin.Context) {

}
