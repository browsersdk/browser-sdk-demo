package middleware

import (
	"dilu/common/utils"
	"dilu/modules/sys/service"

	"github.com/gin-gonic/gin"
)

// 用户权限判断
func PermTeamHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		teamId := utils.GetTeamId(c)
		uid := utils.GetAppUid(c)
		curUri := c.Request.URL.Path
		curMethod := c.Request.Method
		if err := service.SerSysMenu.TeamUserCanAccess(teamId, uid, curMethod, curUri); err != nil {
			Fail(c, 403, "无权限")
			return
		}

		c.Next()
	}
}

// 超管权限控制
func PermAdminHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		curUri := c.Request.URL.Path
		curMethod := c.Request.Method
		adminId := utils.GetAdminId(c)
		if err := service.SerSysMenu.AdminCanAccess(adminId, curMethod, curUri); err != nil {
			Fail(c, 403, "无权限")
			return
		}
		c.Next()
	}
}

// var apis []models.SysApi

// func GetApis() []models.SysApi {
// 	if len(apis) == 0 {
// 		service.SerSysApi.GetByType(3, &apis)
// 	}
// 	return apis
// }

// KeyMatch2 determines whether key1 matches the pattern of key2 (similar to RESTful path), key2 can contain a *.
// For example, "/foo/bar" matches "/foo/*", "/resource1" matches "/:resource"
// func KeyMatch2(key1 string, key2 string) bool {
// 	key2 = strings.Replace(key2, "/*", "/.*", -1)

// 	re := regexp.MustCompile(`:[^/]+`)
// 	key2 = re.ReplaceAllString(key2, "$1[^/]+$2")

// 	return RegexMatch(key1, "^"+key2+"$")
// }

// // RegexMatch determines whether key1 matches the pattern of key2 in regular expression.
// func RegexMatch(key1 string, key2 string) bool {
// 	res, err := regexp.MatchString(key2, key1)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return res
// }
