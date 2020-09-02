package request

import "go_admin/Server/model"

type AuthorityIdInfo struct {
	AuthorityId string
}

type AddMenuAuthorityInfo struct {
	Menus       []model.SysBaseMenu
	AuthorityId string
}
