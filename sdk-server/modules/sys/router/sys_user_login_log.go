package router

import (
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerSysUserLoginLogRouter)
}

// 默认需登录认证的路由
func registerSysUserLoginLogRouter(v1 *gin.RouterGroup) {
	r := v1.Group("user-login-log") //.Use(middleware.JwtAdminHandler()).Use(middleware.PermAdminHandler())
	{
		r.GET("/get", apis.ApiSysUserLoginLog.Get)
		r.POST("/page", apis.ApiSysUserLoginLog.QueryPage)
		// r.POST("/del", apis.ApiSysUserLoginLog.Del)
		// r.POST("/create", apis.ApiSysUserLoginLog.Create)
		// r.POST("/update", apis.ApiSysUserLoginLog.Update)
	}
}
