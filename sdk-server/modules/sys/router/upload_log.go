package router

import (
	"dilu/modules/sys/apis"
	//"dilu/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerUploadLogRouter)
}

// 默认需登录认证的路由
func registerUploadLogRouter(v1 *gin.RouterGroup) {
	r := v1.Group("upload-log") //.Use(middleware.JwtAdminHandler()).Use(middleware.PermAdminHandler())
	{
		r.POST("/get", apis.ApiUploadLog.Get)
		r.POST("/create", apis.ApiUploadLog.Create)
		r.POST("/update", apis.ApiUploadLog.Update)
		r.POST("/page", apis.ApiUploadLog.QueryPage)
		r.POST("/del", apis.ApiUploadLog.Del)
	}
}
