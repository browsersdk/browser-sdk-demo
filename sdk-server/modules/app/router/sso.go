package router

import (
	"dilu/modules/app/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, sysNoCheckSsoRouter)
	routerCheckRole = append(routerCheckRole, sysCheckSsoRouter)
}

// 无需认证的路由示例
func sysNoCheckSsoRouter(v1 *gin.RouterGroup) {
	r := v1.Group("app")
	{
		r.POST("/login", apis.ApiSso.Login)
		r.POST("register", apis.ApiSso.Register)
		r.POST("sendCode", apis.ApiSso.SendCode)
		r.POST("getOneClickToken", apis.ApiSso.GetOneClickToken)
		r.POST("oneClick", apis.ApiSso.OneClickLogin)
		r.POST("getFusionToken", apis.ApiSso.GetFusionToken)
		r.POST("verifyFusion", apis.ApiSso.VerifyFusionLogin)
		r.POST("refreshToken", apis.ApiSso.RefreshToken)
		//r.GET("pbtest", apis.ApiSso.PSTest)
	}

}

func sysCheckSsoRouter(v1 *gin.RouterGroup) {
	r := v1.Group("app/auth")
	{
		r.POST("/logout", apis.ApiSso.Logout)
		r.POST("/auth/bind", apis.ApiSso.Bind)
		//r.POST("/user/changeUserinfo", apis.ApiSso.ChangeUserinfo)
	}
}
