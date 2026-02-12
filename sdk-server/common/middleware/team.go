package middleware

import (
	"dilu/common/consts"
	"dilu/common/utils"
	"dilu/modules/sys/service"

	"github.com/gin-gonic/gin"
)

func TeamHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		appKey := c.GetHeader(consts.TeamHeader)

		if appKey == "" {
			Fail(c, 401, "appKey is empty")
			return
		}

		teamId, err := service.SerSysTeam.ValidAppKey(appKey)
		if err != nil {
			Fail(c, 401, err.Error())
			return
		}

		utils.SetAppTeamId(teamId, c)

		c.Next()
	}
}
