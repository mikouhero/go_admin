package request

type RegisterStruct struct {
	Username    string `json:"userName"`
	Password    string `json:"passWord"`
	Nickname    string `json:"nickName" gorm:"default:'test'"`
	HeaderImg   string `json:"headerImg" gorm:"default:''"`
	AuthorityId string `json:"authorityId" gorm:"default:0"`
}
