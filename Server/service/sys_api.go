package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"go_admin/Server/global"
	"go_admin/Server/model"
	"go_admin/Server/model/request"
)

func CreateApi(api model.SysApi) error {
	var num int
	global.GVA_DB.
		Where("path = ? AND method = ? ", api.Path, api.Method).
		First(&model.SysApi{}).
		Count(&num)
	if num > 0 {
		return errors.New("存在相同的api")
	}
	return global.GVA_DB.Create(&api).Error
}

func DeleteApi(api model.SysApi) error {
	err := global.GVA_DB.Delete(&api).Error
	//ClearCasbin(1, api.Path, api.Method)
	return err
}

func GetApiById(id float64) (err error, api model.SysApi) {
	err = global.GVA_DB.Where("id = ?", id).First(&api).Error
	return
}

func UpdateApi(api model.SysApi) error {
	var oldApi model.SysApi
	err := global.GVA_DB.Where("id = ?", api.ID).First(&oldApi).Error

	if oldApi.Path != api.Path || oldApi.Method != api.Method {
		if !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&model.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		//err = UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		//if err != nil {
		//	return err
		//} else {
		err = global.GVA_DB.Save(&api).Error
		//}
	}

	return err
}

func GetAllApis() (err error, apis []model.SysApi) {
	err = global.GVA_DB.Find(&apis).Error
	return
}

func GetApiInfoList(api model.SysApi, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.SysApi{})
	var apiList []model.SysApi

	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}

	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, apiList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return err, apiList, total
}
