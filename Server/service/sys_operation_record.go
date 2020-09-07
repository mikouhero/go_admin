package service

import (
	"go_admin/Server/global"
	"go_admin/Server/model"
	"go_admin/Server/model/request"
)

func CreateSysOperationRecord(record model.SysOperationRecord) (err error) {
	return global.GVA_DB.Create(&record).Error
}

func DeleteSysOperationRecordByIds(ids request.IdsReq) (err error) {
	return global.GVA_DB.Delete(&[]model.SysOperationRecord{}, " id in (?)", ids).Error
}

func DeleteSysOperationRecord(record model.SysOperationRecord) (err error) {
	return global.GVA_DB.Delete(&record).Error
}

func GetSysOperationRecord(id uint) (err error, record model.SysOperationRecord) {
	err = global.GVA_DB.Where("id = ?", id).First(&record).Error
	return
}
func GetSysOperationRecordInfoList(info request.SysOperationRecordSearch) (err error, list interface{}, total int64) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.SysOperationRecord{})
	var records []model.SysOperationRecord
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&records).Error
	return err, records, total
}
