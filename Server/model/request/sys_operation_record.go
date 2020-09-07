package request

import "go_admin/Server/model"

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	PageInfo
}
