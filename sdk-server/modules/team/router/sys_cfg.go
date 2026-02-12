package router

import (
	"dilu/modules/team/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerPermSysCfgRouter)
	routerCheckRole = append(routerCheckRole, registerSysCfgRouter)
}

// 默认需登录认证的路由
func registerPermSysCfgRouter(v1 *gin.RouterGroup) {
	r := v1.Group("team/cfg") //.Use(middleware.JwtAppHandler()).Use(middleware.PermTeamHandler())
	{
		r.POST("/create", apis.ApiTeamCfg.Create)
		r.POST("/update", apis.ApiTeamCfg.Update)
		r.POST("/del", apis.ApiTeamCfg.Del)
	}

}

func registerSysCfgRouter(v1 *gin.RouterGroup) {
	r2 := v1.Group("team/cfg") //.Use(middleware.JwtAppHandler())
	{
		r2.POST("/get", apis.ApiTeamCfg.Get)
		r2.POST("/page", apis.ApiTeamCfg.QueryPage)
	}
}
