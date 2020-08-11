package core

import (
	"fmt"
	"github.com/spf13/viper"
	"go_admin/Server/global"
)

const defaultConfigFile = "config.yaml"

func init() {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 检测变化
	v.WatchConfig()

	global.GVA_VP = v

}
