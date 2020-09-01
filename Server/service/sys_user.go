package service

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"go_admin/Server/global"
	"go_admin/Server/model"
	"go_admin/Server/model/request"
	"go_admin/Server/utils"
)

// 用户注册
func Register(u model.SysUser) (err error, userInter model.SysUser) {
	var user model.SysUser
	// 判断用户名是否注册 未注册 返回true
	notFound := global.GVA_DB.
		Where("username = ?", u.Username).
		First(&user).
		RecordNotFound()

	if !notFound {
		return errors.New("用户名已注册"), userInter
	} else {
		u.Password = utils.MD5([]byte(u.Password))
		// 生成用户唯一标示
		u.UUID, _ = uuid.NewV4()
		// 数据落库
		err = global.GVA_DB.Create(&u).Error
	}
	return err, u
}

// 通过用户名和密码返回用户信息
func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5([]byte(u.Password))
	err = global.GVA_DB.
		Where("username=? AND password=?", u.Username, u.Password).
		Preload("Authority").
		First(&user).Error
	return err, &user
}

// 修改密码
func ChangePassword(u *model.SysUser, newPassword string) (err error, userInter *model.SysUser) {

	var user model.SysUser
	u.Password = utils.MD5([]byte(u.Password))

	err = global.GVA_DB.
		Where("username = ? AND password = ? ", u.Username, u.Password).
		First(&user).
		Update("password", utils.MD5([]byte(newPassword))).
		Error

	return err, u
}

// 更新用户头像 by uuid
func UploadHeaderImg(uuid uuid.UUID, filePath string) (err error, userInter *model.SysUser) {
	var user model.SysUser
	err = global.GVA_DB.
		Where("uuid = ?", uuid).
		First(&user).
		Update("header_img", filePath).
		Error
	return err, &user
}

// 分页获取用户信息
func GetUserInfoList(info request.PageInfo) (err error, list interface{}, totle int) {
	// 每页个数
	limit := info.PageSize
	// 开始位置
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&model.SysUser{})
	var userList []model.SysUser
	// 总个数
	err = db.Count(&totle).Error
	// 列表
	err = db.
		Limit(limit).
		Offset(offset).
		Preload("Authority"). // SysUser 定义的外键关联关系 ,每个用户附属一个Authority 对应的角色信息
		Find(&userList).
		Error
	return err, userList, totle
}

// 分配角色给用户 by uuid
func SetUserAuthority(uuid uuid.UUID, authorityId string) (err error) {
	err = global.GVA_DB.
		Where("uuid = ?", uuid).
		First(&model.SysUser{}).
		Update("authority_id", authorityId).
		Error
	return err
}

//删除用户 by id
func DeleteUser(id float64) error {
	var user model.SysUser
	return global.GVA_DB.Where("id = ?", id).Delete(&user).Error
}
