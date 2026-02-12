package router

import (
	"dilu/modules/team/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerPermUserNoticeRouter)
	routerCheckRole = append(routerCheckRole, registerUserNoticeRouter)
}

// 默认需登录认证的路由
func registerPermUserNoticeRouter(v1 *gin.RouterGroup) {
	r := v1.Group("team/user-notice") //.Use(middleware.JwtAppHandler()).Use(middleware.PermTeamHandler())
	{
		r.POST("/get", apis.ApiUserNotice.Get)
		r.POST("/create", apis.ApiUserNotice.Create)
		r.POST("/update", apis.ApiUserNotice.Update)
		r.POST("/page", apis.ApiUserNotice.QueryPage)
		r.POST("/del", apis.ApiUserNotice.Del)
	}

}

func registerUserNoticeRouter(v1 *gin.RouterGroup) {
	r2 := v1.Group("team/user-notice") //.Use(middleware.JwtAppHandler())
	{
		r2.POST("my", apis.ApiUserNotice.GetUserNotice)
		r2.POST("read", apis.ApiUserNotice.Read)
	}
}
