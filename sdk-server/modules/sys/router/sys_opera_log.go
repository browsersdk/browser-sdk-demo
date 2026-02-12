package router

import (
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerSysOperaLogRouter)
}

// 默认需登录认证的路由
func registerSysOperaLogRouter(v1 *gin.RouterGroup) {
	r := v1.Group("opera-log") //.Use(middleware.JwtAdminHandler())
	{
		r.POST("/get", apis.ApiSysOperaLog.Get)
		r.POST("/create", apis.ApiSysOperaLog.Create)
		r.POST("/update", apis.ApiSysOperaLog.Update)
		r.POST("/page", apis.ApiSysOperaLog.QueryPage)
		r.POST("/del", apis.ApiSysOperaLog.Del)
	}
}
