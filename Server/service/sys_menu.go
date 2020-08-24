package service

import (
	"errors"
	"go_admin/Server/global"
	"go_admin/Server/model"
)

func GetMenuTreeMap() {

}

func GetMenuTree() {

}

func getChildrenList() {

}

func GetInfoList() {

}

func GetBaseChildrenList() {

}

// 添加菜单
func AddBaseMenu(menu model.SysBaseMenu) error {
	err := global.GVA_DB.Where("name = ? ", menu.Name).Find(&model.SysBaseMenu{}).Error
	if err != nil {
		menu.Hidden = "1"
		err = global.GVA_DB.Create(&menu).Error
	} else {
		return errors.New("存在相同的菜单名称(name)，请修改name")
	}
	return err
}

func GetBaseMenuTreeMap() {

}

func GetBaseMenuTree() {

}

func AddMenuAuthority() {

}

func GetMenuAuthority() {

}
