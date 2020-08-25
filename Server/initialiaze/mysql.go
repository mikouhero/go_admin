package initialiaze

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_admin/Server/global"
	"os"
)

// 初始化数据库并产生数据库全局变量
func Mysql() {

	admin := global.GVA_CONFIG.Mysql
	if db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config); err != nil {
		fmt.Println("mysql faild", err)
		os.Exit(0)
	} else {
		db.LogMode(true)
		global.GVA_DB = db
		global.GVA_DB.DB().SetMaxIdleConns(admin.MaxIdleConns)
		global.GVA_DB.DB().SetMaxOpenConns(admin.MaxOpenConns)
		fmt.Println("mysql connected success")
	}

}
