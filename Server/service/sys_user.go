package service

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"go_admin/Server/global"
	"go_admin/Server/model"
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
		//唯一标示
		u.UUID, _ = uuid.NewV4()
		err = global.GVA_DB.Create(&u).Error
	}
	return err, u
}
