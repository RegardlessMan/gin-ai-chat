package main

import (
	"awesomeProject/core"
	"awesomeProject/global"
	"awesomeProject/initialize"
)

func main() {
	global.VP = core.Viper()
	// other initialize
	initialize.OtherInit()

	// zap log initialize
	global.Log = core.Zap()
	// gorm initialize
	global.DB = initialize.Gorm()
	// start a Web Server
	core.RunWindowsServer()
}
