package core

import (
	"awesomeProject/global"
	"awesomeProject/initialize"
	"fmt"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	s := initServer(address, Router)

	global.Log.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 ai-chat-go
	当前版本:v1.0.0
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, address)
	global.Log.Error(s.ListenAndServe().Error())
}
