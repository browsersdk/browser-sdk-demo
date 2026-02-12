package router

import (
	"dilu/modules/notice/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerPeamTaskRouter)
	routerCheckRole = append(routerCheckRole, registerTaskRouter)
}

// 默认需登录认证的路由
func registerPeamTaskRouter(v1 *gin.RouterGroup) {
	r := v1.Group("task") //.Use(middleware.JwtAdminHandler()).Use(middleware.PermAdminHandler())
	{
		r.POST("/get", apis.ApiTask.Get)
		r.POST("/create", apis.ApiTask.Create)
		r.POST("/update", apis.ApiTask.Update)
		r.POST("/page", apis.ApiTask.QueryPage)
		r.POST("/del", apis.ApiTask.Del)
	}
}

func registerTaskRouter(v1 *gin.RouterGroup) {

	r2 := v1.Group("task") //.Use(middleware.JwtAdminHandler())
	{
		r2.POST("my", apis.ApiTask.UserTasks)
	}
}
