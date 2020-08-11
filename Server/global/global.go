package global

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go_admin/Server/config"
)

//定义全局变量
var (
	GVA_DB     *gorm.DB
	GVA_VP     *viper.Viper
	GVA_CONFGI config.Server
)
