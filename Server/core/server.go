package core

import (
	"fmt"
	"go_admin/Server/global"
	"go_admin/Server/initialiaze"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {

	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialiaze.Redis()
	}
	fmt.Println("程序主进程")

	Router := initialiaze.Routers()

	//自给定文件系统根目录的文件。
	//Router.Static("")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	s.ListenAndServe()
}
