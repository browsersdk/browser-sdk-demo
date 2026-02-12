package router

import (
	"dilu/modules/sys/apis"
	//"dilu/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerSysEmailRouter)
}

// 默认需登录认证的路由
func registerSysEmailRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys-email")//.Use(middleware.JwtAdminHandler()).Use(middleware.PermAdminHandler())
	{
		r.POST("/get", apis.ApiSysEmail.Get)
		r.POST("/create", apis.ApiSysEmail.Create)
		r.POST("/update", apis.ApiSysEmail.Update)
		r.POST("/page", apis.ApiSysEmail.QueryPage)
		r.POST("/del", apis.ApiSysEmail.Del)
	}
}