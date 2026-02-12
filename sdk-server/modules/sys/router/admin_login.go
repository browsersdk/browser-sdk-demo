package router

import (
	"dilu/common/middleware"
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, sysNoCheckAdminRouter)

	routerCheckRole = append(routerCheckRole, sysCheckAdminRouter)
}

// 无需认证的路由示例
func sysNoCheckAdminRouter(v1 *gin.RouterGroup) {
	r := v1.Group("admin")
	{
		r.POST("login", apis.ApiAdminLogin.Login)
		// r.POST("register", apis.ApiAdminLogin.Register)
		// r.POST("sendCode", apis.ApiAdminLogin.SendCode)
		r.POST("forgetPwd", apis.ApiAdminLogin.ForgetPwd)
		r.POST("refreshToken", apis.ApiAdminLogin.RefreshToken)
	}

}

func sysCheckAdminRouter(v1 *gin.RouterGroup) {

	rj := v1.Group("admin").Use(middleware.JwtAdminHandler())
	{
		rj.POST("changePwd", apis.ApiAdminLogin.ChangePwd)
		rj.GET("myUserinfo", apis.ApiAdminLogin.MyUserInfo)
	}
}
