package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type SysUser struct {
	gorm.Model
	UUID        uuid.UUID    `json:"uuid" gorm:"conment:'用户uuid'"`
	Username    string       `json:"userName" gorm:"comment:'用戶名'"`
	Password    string       `json:"passWord" gorm:"comment:'密碼'`
	NickName    string       `json:"nickName" gorm:"default:'系统用户';comment:'昵称'"`
	HeaderImg   string       `json:"headerImg" gorm:"default:'';comment:'用户头像'"`
	Authority   SysAuthority `json:"authoruty" gorm:"ForeignKey:AuthorityId;AssociationForeignKey:AuthorityId;comment:'用户角色'"`
	AuthorityId string       `json:"authorityId" gorm:"default:0;comment:'用户角色ID'"`
}
