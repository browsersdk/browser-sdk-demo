package router

import (
	"dilu/modules/team/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerPubNoticeRouter)
}

// 默认需登录认证的路由
func registerPubNoticeRouter(v1 *gin.RouterGroup) {
	r := v1.Group("team/pub-notice") //.Use(middleware.JwtAppHandler()).Use(middleware.PermTeamHandler())
	{
		r.POST("/get", apis.ApiPubNotice.Get)
		r.POST("/create", apis.ApiPubNotice.Create)
		r.POST("/update", apis.ApiPubNotice.Update)
		r.POST("/page", apis.ApiPubNotice.QueryPage)
		r.POST("/del", apis.ApiPubNotice.Del)
	}
}
