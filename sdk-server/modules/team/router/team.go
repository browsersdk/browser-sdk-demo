package router

import (
	"dilu/modules/team/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerPermSysTeamRouter)
	routerCheckRole = append(routerCheckRole, registerSysTeamRouter)
}

// 默认需登录认证的路由
func registerSysTeamRouter(v1 *gin.RouterGroup) {
	r2 := v1.Group("/team") //.Use(middleware.JwtAppHandler())
	{
		r2.POST("/get", apis.ApiTeam.Get)
		r2.POST("user/menus", apis.ApiTeam.GetUserMenus)
	}
}

func registerPermSysTeamRouter(v1 *gin.RouterGroup) {
	r1 := v1.Group("/team") //.Use(middleware.JwtAppHandler()).Use(middleware.PermTeamHandler())
	{
		r1.POST("/update", apis.ApiTeam.Update)
		r1.POST("menu/grant", apis.ApiTeam.GetGrantMenus)
	}
}
