package service

import (
	"errors"
	"github.com/casbin/casbin/util"
	"go_admin/Server/global"
	"go_admin/Server/model"
	"go_admin/Server/model/request"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"strings"
)

func UpdateCasbin(authorityId string, casbinInfos []request.CasbinInfo) error {
	for _, v := range casbinInfos {
		cm := model.CasbinModel{
			Ptype:       "p",
			AuthorityId: authorityId,
			Path:        v.Path,
			Method:      v.Method,
		}
		addflag := AddCasbin(cm)
		if addflag == false {
			return errors.New("存在相同API 添加失败")
		}
	}
	return nil
}

func AddCasbin(cm model.CasbinModel) bool {
		//e:= Casbin()
	return false
}


func Casbin() *casbin.Enforcer {
	admin := global.GVA_CONFIG.Mysql
	a, _ := gormadapter.NewAdapter(global.GVA_CONFIG.System.DbType, admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname, true)
	e, _ := casbin.NewEnforcer(global.GVA_CONFIG.Casbin.ModelPath, a)
	e.AddFunction("ParamsMatch", ParamsMatchFunc)
	_ = e.LoadPolicy()
	return e
}

func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return ParamsMatch(name1, name2), nil
}

func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}