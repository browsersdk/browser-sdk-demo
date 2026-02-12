package router

import (
	"dilu/common/middleware"
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckPermRole = append(routerCheckPermRole, registerPermSysMenuRouter)
	routerCheckRole = append(routerCheckRole, registerSysMenuRouter)
}

// 默认需登录认证的路由
func registerPermSysMenuRouter(v1 *gin.RouterGroup) {
	r := v1.Group("menu") //.Use(middleware.JwtAdminHandler()).Use(middleware.PermAdminHandler())
	{
		r.POST("/create", apis.SysMenuA.Create)
		r.POST("/update", apis.SysMenuA.Update)
		r.POST("/all", apis.SysMenuA.GetMenus)
		r.POST("/del", apis.SysMenuA.Del)
		r.POST("/get", apis.SysMenuA.Get)
	}
}

func registerSysMenuRouter(v1 *gin.RouterGroup) {
	r2 := v1.Group("").Use(middleware.JwtAdminHandler())
	{
		r2.POST("menu/grant", apis.SysMenuA.GetGrantMenus)
		r2.POST("menu/apis", apis.SysMenuA.GetApis)
		r2.POST("menu/set_apis", apis.SysMenuA.SetApis)
		r2.POST("user/menus", apis.SysMenuA.GetUserMenus)
		//r2.POST("canAccess", apis.SysMenuA.CanAccess)
	}

}
