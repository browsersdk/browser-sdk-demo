package router

import (
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerAdminDeptRouter)
}

func registerAdminDeptRouter(v1 *gin.RouterGroup) {
	r := v1.Group("admin-dept") //.Use(middleware.JwtAdminHandler()).Use(middleware.PermAdminHandler())
	{
		r.POST("/get", apis.ApiAdminDept.Get)
		r.POST("/create", apis.ApiAdminDept.Create)
		r.POST("/update", apis.ApiAdminDept.Update)
		r.POST("/page", apis.ApiAdminDept.QueryPage)
		r.POST("all", apis.ApiAdminDept.List)
		r.POST("/del", apis.ApiAdminDept.Del)
	}
}
