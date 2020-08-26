package service

import (
	"errors"
	"go_admin/Server/global"
	"go_admin/Server/model"
)

func DeleteBaseMenu(id float64) (err error) {

	count := 0
	global.GVA_DB.Model(&model.SysBaseMenu{}).Where("parent_id = ?", id).Count(&count)
	if count == 0 {
		var menu model.SysBaseMenu
		db := global.GVA_DB.Preload("SysAuthoritys").Where("id = ?", id).First(&menu).Delete(&menu)

		if len(menu.SysAuthoritys) > 0 {
			err = db.Association("SysAuthoritys").Delete(menu.SysAuthoritys).Error
		}
		return err

	} else {
		return errors.New("该菜单存在子菜单不可删除")
	}

}

func UpdateBaseMenu(menu model.SysBaseMenu) (err error) {
	var oldMenu model.SysBaseMenu
	updateMap := make(map[string]interface{})
	updateMap["keep_alive"] = menu.KeepAlive
	updateMap["default_menu"] = menu.DefaultMenu
	updateMap["default_menu"] = menu.DefaultMenu
	updateMap["parent_id"] = menu.ParentId
	updateMap["path"] = menu.Path
	updateMap["name"] = menu.Name
	updateMap["hidden"] = menu.Hidden
	updateMap["component"] = menu.Component
	updateMap["title"] = menu.Title
	updateMap["icon"] = menu.Icon
	updateMap["sort"] = menu.Sort

	db := global.GVA_DB.Where("id = ?", menu.ID).Find(&oldMenu)

	if oldMenu.Name != menu.Name {
		notFound := global.GVA_DB.Where("id <> ? and name = ?", menu.ID, menu.Name).First(&model.SysBaseMenu{}).RecordNotFound()
		if !notFound {
			return errors.New("存在相同的name")
		}
	}
	err = db.Updates(updateMap).Error
	return err
}

func GetBaseMenuById(id float64) (err error, menu model.SysBaseMenu) {
	err = global.GVA_DB.Where("id = ?", id).First(&menu).Error
	return
}
