package system

import (
	"awesomeProject/model/common/response"
	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (u *UserApi) GetUserInfo(c *gin.Context) {
	value := c.Query("value")
	if value == "" {
		response.FailWithMessage("参数异常，请检查!", c)
		return
	}
	user, err := userService.GetUserInfo(value)
	if err == nil {
		response.FailWithMessage("用户不存在!", c)
		return
	} else {
		response.OkWithDetailed(user, "查询成功!", c)
		return
	}

}
