package router

import (
	"dilu/modules/team/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, sysNoCheckSsoRouter)
	routerCheckRole = append(routerCheckRole, sysCheckSsoRouter)
}

// 无需认证的路由示例
func sysNoCheckSsoRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys")
	{
		r.POST("/login", apis.ApiSso.Login)
		r.POST("register", apis.ApiSso.Register)
		r.POST("sendCode", apis.ApiSso.SendCode)
		r.POST("forgetPwd", apis.ApiSso.ForgetPwd)
		r.POST("refreshToken", apis.ApiSso.RefreshToken)
		r.POST("oneClick", apis.ApiSso.OneClickLogin)
	}
}

// 无需认证的路由示例
func sysCheckSsoRouter(v1 *gin.RouterGroup) {
	rj := v1.Group("/sys") //.Use(middleware.JwtAppHandler())
	{
		rj.POST("changePwd", apis.ApiSso.ChangePwd)
		rj.GET("myUserinfo", apis.ApiSso.MyUserInfo)

	}

	r2 := v1.Group("team") //.Use(middleware.JwtAppHandler())
	{
		r2.POST("user/page", apis.ApiSysUser.QueryPage)
	}
}
