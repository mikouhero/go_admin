package service

import (
	"go_admin/Server/global"
	"go_admin/Server/model"
)

func IsBlackList(jwt string, jwtList model.JwtBlacklist) bool {

	isNotFound := global.GVA_DB.Where("jwt = ?", jwt).First(&jwtList).RecordNotFound()

	return !isNotFound
}
