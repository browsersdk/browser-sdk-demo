package router

import (
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerPermSysMemberRouter)
	routerCheckRole = append(routerCheckRole, registerSysMemberRouter)
}

// 默认需登录认证的路由
func registerPermSysMemberRouter(v1 *gin.RouterGroup) {
	r := v1.Group("member") //.Use(middleware.JwtAdminHandler()).Use(middleware.PermAdminHandler())
	{
		r.POST("/get", apis.ApiSysMember.Get)
		r.POST("/create", apis.ApiSysMember.Create)
		r.POST("/update", apis.ApiSysMember.Update)
		r.POST("/page", apis.ApiSysMember.QueryPage)
		r.POST("/del", apis.ApiSysMember.Del)
	}
}

func registerSysMemberRouter(v1 *gin.RouterGroup) {
	r2 := v1.Group("member") //.Use(middleware.JwtAdminHandler())
	{
		//r2.POST("myTeams", apis.ApiSysMember.MyTeams)
		//r2.POST("myInfo", apis.ApiSysMember.MyInfo)
		//r2.POST("changeMyInfo", apis.ApiSysMember.ChangeMyInfo)
		r2.POST("members", apis.ApiSysMember.GetMembers)
		r2.POST("exist", apis.ApiSysMember.Exits)
	}
}
