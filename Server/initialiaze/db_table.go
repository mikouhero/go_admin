package initialiaze

import (
	"go_admin/Server/global"
	"go_admin/Server/model"
)

// 自动同步数据表结构
func DBTables() {
	db := global.GVA_DB
	db.AutoMigrate(
		model.SysUser{},
		model.SysAuthority{},
	)

}
