package router

import (
	"dilu/modules/sys/apis"
	//"dilu/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerSysSmsRouter)
}

// 默认需登录认证的路由
func registerSysSmsRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys-sms")//.Use(middleware.JwtAdminHandler()).Use(middleware.PermAdminHandler())
	{
		r.POST("/get", apis.ApiSysSms.Get)
		r.POST("/create", apis.ApiSysSms.Create)
		r.POST("/update", apis.ApiSysSms.Update)
		r.POST("/page", apis.ApiSysSms.QueryPage)
		r.POST("/del", apis.ApiSysSms.Del)
	}
}