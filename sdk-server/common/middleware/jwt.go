package middleware

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"dilu/common/utils"

	"github.com/baowk/dilu-core/core"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAppHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		accessToken, err := GetAccessToken(authorization)
		if err != nil {
			Fail(c, 401, err.Error())
			return
		}

		//slog.Debug("JwtAppHandler", "accessToken", accessToken)
		var appClaims utils.AppClaims
		//fmt.Println(accessToken)
		// 解析Token
		err = utils.ParseApp(accessToken, &appClaims, core.Cfg.JWT.SignKey, jwt.WithSubject(core.Cfg.JWT.Subject))
		if err != nil || appClaims.Uid == 0 {
			slog.Error("JwtAppHandler", "err", err, "uid ", appClaims.Uid, "URL", c.Request.URL.Path)
			Fail(c, 401, err.Error())
			return
		}

		utils.SetAppClaims(appClaims, c)

		c.Next()
	}
}

func JwtAdminHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		accessToken, err := GetAccessToken(authorization)
		if err != nil {
			Fail(c, 401, err.Error())
			return
		}
		//slog.Debug("JwtAdminHandler", "accessToken", accessToken)
		var admC utils.AdminClaims
		// 解析Token
		err = utils.ParseAdmin(accessToken, &admC, core.Cfg.JWT.SignKey, jwt.WithSubject(core.Cfg.JWT.Subject))
		if err != nil || admC.Uid == 0 {
			slog.Error("JwtAdminHandler", "err", err, "uid ", admC.Uid, "URL", c.Request.URL.Path)
			Fail(c, 401, err.Error())
			return
		}
		utils.SetAdminClaims(&admC, c)
		c.Next()
	}
}

func Fail(c *gin.Context, code int, msg string, data ...any) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

// GetAccessToken 获取jwt的Token
func GetAccessToken(authorization string) (accessToken string, err error) {
	if authorization == "" {
		return "", errors.New("authorization header is missing")
	}

	// 检查 Authorization 头的格式
	if !strings.HasPrefix(authorization, "Bearer ") {
		return "", errors.New("invalid Authorization header format")
	}

	// 提取 Token 的值
	accessToken = strings.TrimPrefix(authorization, "Bearer ")
	return
}
