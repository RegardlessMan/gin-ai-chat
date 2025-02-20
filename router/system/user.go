package system

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (r *UserRouter) InitUserApi(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")

	{
		userRouter.GET("getUserList", userApi.GetUserInfo)
	}
}
