package router

import (
	"dilu/modules/team/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerSysDeptRouter)
}

// 默认需登录认证的路由
func registerSysDeptRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/team/dept") //.Use(middleware.JwtAppHandler()).Use(middleware.PermTeamHandler())
	{
		r.POST("/get", apis.ApiTeamDept.Get)
		r.POST("/create", apis.ApiTeamDept.Create)
		r.POST("/update", apis.ApiTeamDept.Update)
		//r.POST("/page", apis.ApiTeamDept.QueryPage)
		r.POST("/del", apis.ApiTeamDept.Del)
		r.POST("/all", apis.ApiTeamDept.List)
	}
}
