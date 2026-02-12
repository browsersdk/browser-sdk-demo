package router

import (
	//"dilu/common/middleware"

	"dilu/modules/app/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerBrowserRouter)
}

// 默认需登录认证的路由
func registerBrowserRouter(v1 *gin.RouterGroup) {
	r := v1.Group("app/browser")
	{
		r.POST("page", apis.ApiBrowser.QueryPage)
		r.POST("get", apis.ApiBrowser.Get)
		r.POST("create", apis.ApiBrowser.Create)
		r.POST("update", apis.ApiBrowser.Update)
		r.POST("del", apis.ApiBrowser.Del)
	}
}
