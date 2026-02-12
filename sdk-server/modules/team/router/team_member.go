package router

import (
	"dilu/modules/team/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerPermSysMemberRouter)
	routerCheckRole = append(routerCheckRole, registerSysMemberRouter)
}

// 默认需登录认证的路由
func registerPermSysMemberRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/team/member") //.Use(middleware.JwtAppHandler()).Use(middleware.PermTeamHandler())
	{
		r.POST("/get", apis.ApiTeamMember.Get)
		r.POST("/create", apis.ApiTeamMember.Create)
		r.POST("/update", apis.ApiTeamMember.Update)
		r.POST("/page", apis.ApiTeamMember.QueryPage)
		r.POST("/del", apis.ApiTeamMember.Del)
	}
}

func registerSysMemberRouter(v1 *gin.RouterGroup) {
	r2 := v1.Group("/team/member") //.Use(middleware.JwtAppHandler())
	{
		r2.POST("myTeams", apis.ApiTeamMember.MyTeams)
		r2.POST("myInfo", apis.ApiTeamMember.MyInfo)
		r2.POST("changeMyInfo", apis.ApiTeamMember.ChangeMyInfo)
		r2.POST("exist", apis.ApiTeamMember.Exits)
	}
}
