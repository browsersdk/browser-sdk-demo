package router

import (
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerAdminRouter)
}

// 默认需登录认证的路由
func registerAdminRouter(v1 *gin.RouterGroup) {
	r := v1.Group("admin") //.Use(middleware.JwtAdminHandler()).Use(middleware.PermAdminHandler())
	{
		r.POST("/get", apis.ApiAdmin.Get)
		r.POST("/create", apis.ApiAdmin.Create)
		r.POST("/update", apis.ApiAdmin.Update)
		r.POST("/page", apis.ApiAdmin.QueryPage)
		r.POST("/del", apis.ApiAdmin.Del)
	}
}
