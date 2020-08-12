package main

import (
	"go_admin/Server/initialiaze"
	_ "go_admin/Server/core"
)

/**
程序的主入口
*/
func main() {
	initialiaze.Mysql()
	initialiaze.DBTables()
}
