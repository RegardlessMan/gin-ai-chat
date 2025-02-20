package system

import "awesomeProject/service"

type ApiGroup struct {
	UserApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
