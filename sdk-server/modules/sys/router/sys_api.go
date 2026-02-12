package router

import (
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerSysApiRouter)
}

// 默认需登录认证的路由
func registerSysApiRouter(v1 *gin.RouterGroup) {
	r := v1.Group("api") //.Use(middleware.JwtAdminHandler()).Use(middleware.PermAdminHandler())
	{
		r.POST("/get", apis.SysApiA.Get)
		r.POST("/create", apis.SysApiA.Create)
		r.POST("/update", apis.SysApiA.Update)
		r.POST("/page", apis.SysApiA.QueryPage)
		r.POST("/del", apis.SysApiA.Del)
	}
	r2 := v1.Group("api")
	{
		r2.POST("/all", apis.SysApiA.Query)
	}
}
