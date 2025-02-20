package system

import (
	"awesomeProject/global"
	model "awesomeProject/model/system"
)

type UserService struct{}

func (UserService) GetUserInfo(value string) (user model.User, err error) {
	err = global.DB.Where("id = ?", value).First(&user).Error
	return
}
