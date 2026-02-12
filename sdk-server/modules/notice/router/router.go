package router

import (
	"dilu/common/consts"
	"dilu/common/middleware"

	"github.com/gin-gonic/gin"

	"github.com/baowk/dilu-core/core"
)

var (
	//routerNoCheckRole   = make([]func(*gin.RouterGroup), 0)
	routerCheckRole     = make([]func(v1 *gin.RouterGroup), 0)
	routerCheckPermRole = make([]func(v1 *gin.RouterGroup), 0)
)

// InitRouter 路由初始化
func InitRouter() {
	r := core.GetGinEngine()
	checkRoleRouter(r)
	checkRolePermRouter(r)
}

func checkRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v := r.Group(consts.ApiRoot+"/notice", middleware.JwtAdminHandler())

	for _, f := range routerCheckRole {
		f(v)
	}
}

func checkRolePermRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v := r.Group(consts.ApiRoot+"/notice", middleware.JwtAdminHandler(), middleware.PermAdminHandler())

	for _, f := range routerCheckPermRole {
		f(v)
	}
}
