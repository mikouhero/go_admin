package main

import (
	"go_admin/Server/core"
	"go_admin/Server/global"
	"go_admin/Server/initialiaze"
)

/**
程序的主入口
*/
func main() {
	// 连接数据库
	initialiaze.Mysql()
	// 数据表结构
	initialiaze.DBTables()
	//程序结束 关闭mysql连接
	defer global.GVA_DB.Close()

	// 执行程序主体
	core.RunWindowsServer()
}
