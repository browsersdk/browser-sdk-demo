package router

import (
	"dilu/modules/team/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerTaskRouter)
}

// 默认需登录认证的路由
func registerTaskRouter(v1 *gin.RouterGroup) {
	r := v1.Group("team/task") //.Use(middleware.JwtAppHandler()).Use(middleware.PermTeamHandler())
	{
		r.POST("/get", apis.ApiTask.Get)
		r.POST("/create", apis.ApiTask.Create)
		r.POST("/update", apis.ApiTask.Update)
		r.POST("/page", apis.ApiTask.QueryPage)
		r.POST("/del", apis.ApiTask.Del)
	}

	// r2 := v1.Group("team/task").Use(middleware.JwtAppHandler())
	// {
	// 	r2.POST("my", apis.ApiTask.UserTasks)
	// }
}
