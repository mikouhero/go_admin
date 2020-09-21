package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/Server/global/response"
	"go_admin/Server/model"
	"go_admin/Server/model/request"
	response2 "go_admin/Server/model/response"
	"go_admin/Server/service"
)

func CreateSysOperationRecord(c *gin.Context) {
	var record model.SysOperationRecord
	_ = c.ShouldBindJSON(&record)
	err := service.CreateSysOperationRecord(record)
	if err != nil {
		response.FailWithMsg("创建失败"+err.Error(), c)
	} else {
		response.OkWithData("ok", c)
	}
}

func DeleteSysOperationRecord(c *gin.Context) {
	var record model.SysOperationRecord
	_ = c.ShouldBindJSON(&record)
	err := service.DeleteSysOperationRecord(record)
	if err != nil {
		response.FailWithMsg("删除失败"+err.Error(), c)
	} else {
		response.Ok(c)
	}

}

func DeleteSysOperationRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteSysOperationRecordByIds(IDS)
	if err != nil {
		response.FailWithMsg("删除失败"+err.Error(), c)
	} else {
		response.Ok(c)
	}
}

func FindSysOperationRecord(c *gin.Context) {
	var record model.SysOperationRecord
	_ =c.ShouldBindJSON(&record)
	err, operationRecord := service.GetSysOperationRecord(record.ID)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"resysOperationRecord": operationRecord}, c)
	}
}

func GetSysOperationRecordList(c *gin.Context) {
	var pageInfo request.SysOperationRecordSearch
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := service.GetSysOperationRecordInfoList(pageInfo)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(response2.PageResult{
			List:     list,
			Totle:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}

}
