package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type SysOperationRecord struct {
	gorm.Model
	Ip           string        `json:"ip" form:"ip" gorm:"column:ip;comment:'请求ip'"`
	Method       string        `json:"method" form:"mehod" gorm:"column:method"`
	Path         string        `json:"path" form:"path" gorm:"column:path;comment:''"`
	Status       int           `json:"status" form:"status" gorm:"column:status;comment:''"`
	Latency      time.Duration `json:"latency" form:"latency" gorm:"column:latency;comment:''"`
	Agent        string        `json:"agent" form:"agent" gorm:"column:agent;comment:''"`
	ErrorMessage string        `json:"error_message" form:"error_message" gorm:"column:error_message;comment:''"`
	Body         string        `json:"body" form:"body" gorm:"column:body;comment:'请求Body'"`
	Resp         string        `json:"resp" form:"resp" gorm:"column:resp;comment:'响应Body'"`
	UserId       string        `json:"user_id" form:"user_id" gorm:"column:user_id;comment:''"`
	User         SysUser       `json:"user"`
}
