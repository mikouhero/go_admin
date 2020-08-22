package initialiaze

import (
	"go_admin/Server/global"
	"go_admin/Server/model"
)

// 自动同步数据表结构
func DBTables() {
	db := global.GVA_DB
	db.AutoMigrate(
		model.SysAuthority{},
		model.SysApi{},
		model.SysBaseMenu{},
		model.JwtBlacklist{},
		model.SysWorkflow{},
		model.SysWorkflowStepInfo{},
		model.SysDictionary{},
		model.SysDictionaryDetail{},
		model.ExaFileUploadAndDownload{},
		model.ExaFile{},
		model.ExaFileChunk{},
		model.ExaCustomer{},
		model.SysOperationRecord{},
	)

}
