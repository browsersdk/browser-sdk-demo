package router

import (
	"dilu/common/consts"

	"github.com/baowk/dilu-core/core"
	"github.com/gin-gonic/gin"

	"dilu/common/middleware"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole   = make([]func(*gin.RouterGroup), 0)
	routerApiKeyRole  = make([]func(*gin.RouterGroup), 0)
)

// InitRouter 路由初始化
func InitRouter() {
	r := core.GetGinEngine()
	noCheckRoleRouter(r)
	checkRoleRouter(r)
	apiKeyRouter(r)
}

// noCheckRoleRouter 无需认证的路由
func noCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v := r.Group(consts.ApiRoot)

	for _, f := range routerNoCheckRole {
		f(v)
	}
}

func apiKeyRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v := r.Group(consts.ApiRoot)

	for _, f := range routerApiKeyRole {
		f(v)
	}
}

func checkRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v := r.Group(consts.ApiRoot, middleware.JwtAppHandler())

	for _, f := range routerCheckRole {
		f(v)
	}
}
