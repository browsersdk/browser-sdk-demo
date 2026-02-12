package middleware

// import (
// 	"dilu/common/consts"
// 	"dilu/common/utils"

// 	"github.com/gin-gonic/gin"
// )

// func AppHandler() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		appKey := c.GetHeader(consts.AppAPIKeyHeader)

// 		if appKey == "" {
// 			Fail(c, 401, "appKey is empty")
// 			return
// 		}

// 		teamId, appId, err := service.SerApp.ValidAppKey(appKey)
// 		if err != nil {
// 			Fail(c, 401, err.Error())
// 			return
// 		}

// 		utils.SetAppId(appId, c)
// 		utils.SetAppTeamId(teamId, c)

// 		c.Next()
// 	}
// }
