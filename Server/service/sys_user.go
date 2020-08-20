package service

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"go_admin/Server/global"
	"go_admin/Server/model"
	"go_admin/Server/model/request"
	"go_admin/Server/utils"
)

func Register(u model.SysUser) (err error, userInter model.SysUser) {
	var user model.SysUser
	// 判断用户名是否注册
	notFound := global.GVA_DB.Where("username = ?", u.Username).First(&user).RecordNotFound()

	if !notFound {
		return errors.New("用户名已注册"), userInter
	} else {
		//密码加密
		//u.Password =""
		//唯一标示,
		u.Password = utils.MD5([]byte(u.Password))
		u.UUID, _ = uuid.NewV4()
		err = global.GVA_DB.Create(&u).Error
	}
	return err, u
}

func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5([]byte(u.Password))
	err = global.GVA_DB.Where("username=? AND password=?", u.Username, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}

func ChangePassword(u *model.SysUser, newPassword string) (err error, userInter *model.SysUser) {

	var user model.SysUser
	u.Password = utils.MD5([]byte(u.Password))

	err = global.GVA_DB.Where("username = ? AND password = ? ", u.Username, u.Password).First(&user).Update("password", utils.MD5([]byte(newPassword))).Error

	return err, u
}

func UploadHeaderImg(uuid uuid.UUID, filePath string) (err error, userInter *model.SysUser) {
	var user model.SysUser
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&user).Update("header_img", filePath).Error
	return err, &user
}

func GetUserInfoList(info request.PageInfo) (err error, list interface{}, totle int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.SysUser{})
	var userList []model.SysUser
	err = db.Count(&totle).Error
	err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
	

	return err, userList, totle
}
