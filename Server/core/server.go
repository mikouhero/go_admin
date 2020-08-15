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

	fmt.Println("程序主进程")

	Router := initialiaze.Routers()

	//自给定文件系统根目录的文件。
	//Router.Static("")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	fmt.Println(Router,address)
	s := initServer(address, Router)
	s.ListenAndServe()
}
