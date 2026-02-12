package router

import (
	"dilu/modules/sys/apis"
	//"dilu/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerCounterRouter)
}

// 默认需登录认证的路由
func registerCounterRouter(v1 *gin.RouterGroup) {
	r := v1.Group("counter")//.Use(middleware.JwtAdminHandler()).Use(middleware.PermAdminHandler())
	{
		r.POST("/get", apis.ApiCounter.Get)
		r.POST("/create", apis.ApiCounter.Create)
		r.POST("/update", apis.ApiCounter.Update)
		r.POST("/page", apis.ApiCounter.QueryPage)
		r.POST("/del", apis.ApiCounter.Del)
	}
}