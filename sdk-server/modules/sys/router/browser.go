package router

import (
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerBrowserRouter)
}

// 默认需登录认证的路由
func registerBrowserRouter(v1 *gin.RouterGroup) {
	r := v1.Group("browser") //.Use(middleware.JwtAdminHandler()).Use(middleware.PermAdminHandler())
	{
		r.POST("/get", apis.ApiBrowser.Get)
		r.POST("/create", apis.ApiBrowser.Create)
		r.POST("/update", apis.ApiBrowser.Update)
		r.POST("/page", apis.ApiBrowser.QueryPage)
		r.POST("/del", apis.ApiBrowser.Del)
	}
}