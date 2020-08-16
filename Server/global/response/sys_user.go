package response

import "go_admin/Server/model"

type SysUserResponse struct {
	User model.SysUser `json:"user"`
}