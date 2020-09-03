package request

import "go_admin/Server/model"

type SearchApiParams struct {
	model.SysApi
	PageInfo
	OrderKey string `json:"orderKey"`
	Desc     bool   `json:"desc"`
}
