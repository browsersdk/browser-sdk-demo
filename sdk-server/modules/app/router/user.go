package router

import (
	//"dilu/common/middleware"

	"dilu/modules/app/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerUserRouter)
}

// 默认需登录认证的路由
func registerUserRouter(v1 *gin.RouterGroup) {
	r := v1.Group("app/user")
	{
		r.GET("info", apis.ApiAppUser.GetUserInfo)
		r.POST("/upload", apis.ApiAppUser.Upload)
	}
}
