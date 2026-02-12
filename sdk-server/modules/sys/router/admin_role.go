package router

import (
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerAdminRoleRouter)
}

// 默认需登录认证的路由
func registerAdminRoleRouter(v1 *gin.RouterGroup) {
	r := v1.Group("admin-role") //.Use(middleware.JwtAdminHandler()).Use(middleware.PermAdminHandler())
	{
		r.POST("/get", apis.ApiAdminRole.Get)
		r.POST("/create", apis.ApiAdminRole.Create)
		r.POST("/update", apis.ApiAdminRole.Update)
		r.POST("/page", apis.ApiAdminRole.QueryPage)
		r.POST("/del", apis.ApiAdminRole.Del)
		r.POST("list", apis.ApiAdminRole.List)
	}
}
