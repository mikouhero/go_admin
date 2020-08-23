package service

import (
	"errors"
	"go_admin/Server/global"
	"go_admin/Server/model"
)

// 创建角色
func CreateAuthority(auth model.SysAuthority) (err error, authority model.SysAuthority) {
	var authorityBox model.SysAuthority
	notHas := global.GVA_DB.Where("authority_id = ?", authority.AuthorityId).Find(&authorityBox).RecordNotFound()
	if !notHas {
		return errors.New("存在相同的角色ID"), auth
	}
	err = global.GVA_DB.Create(&auth).Error
	return err, auth
}

//func CopyAuthority(copyInfo response.SysAuthorityCopyResponse) (err error, authority model.SysAuthority) {
//	var authorityBox model.SysAuthority
//	notHas := global.GVA_DB.Where("authority_id = ?", copyInfo.Authority.AuthorityId).Find(&authorityBox).RecordNotFound()
//	if !notHas {
//		return errors.New("存在相同的角色ID"), authority
//	}
//
//	copyInfo.Authority.Children = []model.SysAuthority{}
//}

func UpdateAuthority(auth model.SysAuthority) (err error, authority model.SysAuthority) {
	err = global.GVA_DB.Where("authority_id =? ", auth.AuthorityId).First(&model.SysAuthority{}).Update(&auth).Error
	return err, authority
}

func DeleteAuthority(auth *model.SysAuthority) (err error) {
	err = global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).Find(&model.SysUser{}).Error
	if err != nil {
		return errors.New("角色正在使用 禁止删除")
	}
	err = global.GVA_DB.Where("parend_id = ?", auth.AuthorityId).Find(&model.SysAuthority{}).Error
	if err != nil {
		return errors.New("此角色存在子角色 不允许删除")
	}
	db := global.GVA_DB.Preload("SysBaseMenus").Where("authority_id =?", auth.AuthorityId).First(auth).Unscoped().Delete(auth)
	if len(auth.SysBaseMenus) > 0 {
		err = db.Association("SysBaseMenus").Delete(auth.SysBaseMenus).Error
	} else {
		err = db.Error
	}
	//ClearCasbin(0, auth.AuthorityId)
	return err
}

func GetAuthorityInfoList() {

}

func GetAuthorityInfo() {

}

func SetDataAuthority() {

}

func SetMenuAuthority() {

}

func FindChildrenAuthority() {

}
