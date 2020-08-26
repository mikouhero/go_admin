package service

import (
	"errors"
	"go_admin/Server/global"
	"go_admin/Server/model"
	"go_admin/Server/model/request"
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
	err = global.GVA_DB.Where("parent_id = ?", auth.AuthorityId).Find(&model.SysAuthority{}).Error

	if err == nil {
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

func GetAuthorityInfoList(info request.PageInfo) (err error, list interface{}, totle int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB
	var authority []model.SysAuthority
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = 0").Find(&authority).Error
	if len(authority) > 0 {
		for k := range authority {
			err = FindChildrenAuthority(&authority[k])
		}
	}

	return err, authority, totle
}

// 获取所有角色信息
func GetAuthorityInfo(auth model.SysAuthority) (err error, sa model.SysAuthority) {
	err = global.GVA_DB.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return err, sa

}

// 设置资源角色权限
func SetDataAuthority(auth model.SysAuthority) error {
	var s model.SysAuthority
	global.GVA_DB.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.GVA_DB.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId).Error
	return err
}

// 菜单与角色绑定
func SetMenuAuthority(auth *model.SysAuthority) (err error) {
	var s model.SysAuthority
	global.GVA_DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err = global.GVA_DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus).Error
	return err
}

// 查询子角色
func FindChildrenAuthority(authority *model.SysAuthority) (err error) {
	err = global.GVA_DB.Preload("DataAuthotityId").Where("parend_id= ?", authority.AuthorityId).Find(&authority.Children).Error

	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = FindChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}
