package router

import (
	"dilu/modules/team/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerSysRoleRouter)
}

// 默认需登录认证的路由
func registerSysRoleRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/team/role") //.Use(middleware.JwtAppHandler()).Use(middleware.PermTeamHandler())
	{
		r.POST("/get", apis.ApiTeamRole.Get)
		r.POST("/create", apis.ApiTeamRole.Create)
		r.POST("/update", apis.ApiTeamRole.Update)
		r.POST("/page", apis.ApiTeamRole.QueryPage)
		r.POST("/del", apis.ApiTeamRole.Del)
		r.POST("/list", apis.ApiTeamRole.List)
	}
}
