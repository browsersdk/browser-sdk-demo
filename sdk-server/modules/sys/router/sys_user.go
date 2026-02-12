package router

import (
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerSysUserRouter)
}

// 默认需登录认证的路由
func registerSysUserRouter(v1 *gin.RouterGroup) {
	r := v1.Group("user") //.Use(middleware.JwtAdminHandler()).Use(middleware.PermAdminHandler())
	{
		r.POST("/get", apis.ApiSysUser.Get)
		r.POST("/create", apis.ApiSysUser.Create)
		r.POST("/update", apis.ApiSysUser.Update)
		r.POST("/page", apis.ApiSysUser.QueryPage)
		r.POST("/del", apis.ApiSysUser.Del)
		//r.POST("/list", apis.ApiSysUser.List)
	}
}
