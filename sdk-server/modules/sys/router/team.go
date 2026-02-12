package router

import (
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysTeamRouter)
}

// 默认需登录认证的路由
func registerSysTeamRouter(v1 *gin.RouterGroup) {
	r1 := v1.Group("team") //.Use(middleware.JwtAdminHandler()).Use(middleware.PermAdminHandler())
	{
		r1.POST("/change", apis.ApiSysTeam.ChangeName)
		r1.POST("/get", apis.ApiSysTeam.Get)
		r1.POST("/update", apis.ApiSysTeam.Update)
		r1.POST("/create", apis.ApiSysTeam.Create)
		r1.POST("/page", apis.ApiSysTeam.QueryPage)
		r1.POST("/del", apis.ApiSysTeam.Del)
	}
}
