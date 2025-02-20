package system

import api "awesomeProject/api/v1"

type RouterGroup struct {
	UserRouter
}

var (
	userApi = api.ApiGroupApp.SystemApiGroup.UserApi
)
