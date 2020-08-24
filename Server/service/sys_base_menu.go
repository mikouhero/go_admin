package service

import (
	"errors"
	"fmt"
	"go_admin/Server/global"
	"go_admin/Server/model"
)

func DeleteBaseMenu(id float64) (err error) {

	count := 0
	global.GVA_DB.Model(&model.SysBaseMenu{}).Where("parent_id = ?", id).Count(&count)
	if count == 0 {
		var menu model.SysBaseMenu
		db:= global.GVA_DB.Preload("SysAuthoritys").Where("id = ?", id).First(&menu).Delete(&menu)

		fmt.Println(menu.SysAuthoritys)
		if len(menu.SysAuthoritys) > 0 {
			err = db.Association("SysAuthoritys").Delete(menu.SysAuthoritys).Error
			fmt.Println(err)
		}

	} else {
		return errors.New("该菜单存在子菜单不可删除")
	}
	return err
}

func UpdateBaseMenu() {

}

func GetBaseMenuById() {

}
