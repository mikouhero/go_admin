package core

import "fmt"

type server interface {
	ListenAndServer() error
}

func RunWindowsServer() {

	fmt.Println("程序主进程")
}
