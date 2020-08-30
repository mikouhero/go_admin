package request

import uuid "github.com/satori/go.uuid"

type RegisterStruct struct {
	Username    string `json:"userName"`
	Password    string `json:"passWord"`
	Nickname    string `json:"nickName" gorm:"default:'test'"`
	HeaderImg   string `json:"headerImg" gorm:"default:''"`
	AuthorityId string `json:"authorityId" gorm:"default:0"`
}

type RegisterAndLoginStruct struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

type ChangePasswordStruct struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

type SetUserAuth struct {
	UUID        uuid.UUID `json:"uuid"`
	AuthorityId string    `json:"authorityId"`
}
