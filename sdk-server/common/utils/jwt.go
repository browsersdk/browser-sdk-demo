package utils

import (
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 生成admin token
func GenerateAdminToken(uid, roleId int, issuer, subject, secretKey string, expiresAt time.Time) (string, error) {
	secretKey = secretKey + TOKEN_ADMIN
	return generate(newAdminClaims(uid, roleId, expiresAt, issuer, subject), secretKey)
}

// 解析admin token
func ParseAdmin(accessToken string, claims jwt.Claims, secretKey string, options ...jwt.ParserOption) error {
	secretKey = secretKey + TOKEN_ADMIN
	return parse(accessToken, claims, secretKey, options...)
}

// 生成admin refresh token
func GenerateAdminRefreshToken(uid, roleId int, issuer, subject, secretKey string, expiresAt time.Time) (string, error) {
	secretKey = secretKey + TOKEN_ADMIN_REFRESH
	return generate(newAdminClaims(uid, roleId, expiresAt, issuer, subject), secretKey)
}

// 解析admin refresh token
func ParseAdminRefresh(accessToken string, claims jwt.Claims, secretKey string, options ...jwt.ParserOption) error {
	secretKey = secretKey + TOKEN_ADMIN_REFRESH
	return parse(accessToken, claims, secretKey, options...)
}

// 生成app token
func GenerateAppToken(uid int, issuer, subject, secretKey string, expiresAt time.Time) (string, error) {
	secretKey = secretKey + TOKEN_APP

	return generate(newAppClaims(uid, expiresAt, issuer, subject), secretKey)
}

// 生成app refresh token
func GenerateAppRefreshToken(uid int, issuer, subject, secretKey string, expiresAt time.Time) (string, error) {
	secretKey = secretKey + TOKEN_APP_REFRESH
	return generate(newAppClaims(uid, expiresAt, issuer, subject), secretKey)
}

// 解析app token
func ParseApp(accessToken string, claims jwt.Claims, secretKey string, options ...jwt.ParserOption) error {
	secretKey = secretKey + TOKEN_APP
	return parse(accessToken, claims, secretKey, options...)
}

// 解析app refresh token
func ParseAppRefresh(accessToken string, claims jwt.Claims, secretKey string, options ...jwt.ParserOption) error {
	secretKey = secretKey + TOKEN_APP_REFRESH
	return parse(accessToken, claims, secretKey, options...)
}

func GetAppUid(c *gin.Context) int {
	return GetInt(c, AppUid)
}

// 获取当前登录用户adminId
func GetAdminId(c *gin.Context) int {
	return GetInt(c, AdmUid)
}

func GetUint64(c *gin.Context, key string) uint64 {
	uid := c.GetUint64(key)
	if uid > 0 {
		return uid
	}
	suid := c.GetHeader(key)
	if suid != "" {
		uid, _ = strconv.ParseUint(suid, 10, 64)
		return uid
	}
	return 0
}

func GetInt(c *gin.Context, key string) int {
	uid := c.GetInt(key)
	if uid > 0 {
		return uid
	}
	suid := c.GetHeader(key)
	if suid != "" {
		uid, _ = strconv.Atoi(suid)
		return uid
	}
	return 0
}

func GetRoleId(c *gin.Context) int {
	rid := c.GetInt(AdmRid)
	if rid == 0 {
		suid := c.GetHeader(AdmRid)
		if suid != "" {
			rid, _ = strconv.Atoi(suid)
		}
	}
	return rid
}

func SetAdminClaims(admC *AdminClaims, c *gin.Context) {
	c.Set(AdmUid, admC.Uid)
}

func SetAppClaims(appC AppClaims, c *gin.Context) {
	c.Set(AppUid, appC.Uid)
}

// Parse 解析token
func parse(accessToken string, claims jwt.Claims, secretKey string, options ...jwt.ParserOption) error {
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secretKey), err
	}, options...)
	if err != nil {
		return err
	}

	// 对token对象中的Claim进行类型断言
	if token.Valid { // 校验token
		return nil
	}

	return errors.New("invalid token")
}

// Generate 生成JWT Token
func generate(claims jwt.Claims, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成签名字符串
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

type AppClaims struct {
	Uid                  int `json:"uid,omitempty"`
	jwt.RegisteredClaims     // 内嵌标准的声明
}

// AdminClaims 自定义格式内容
type AdminClaims struct {
	Uid                  int `json:"uid,omitempty"`
	Rid                  int `json:"rid,omitempty"`
	jwt.RegisteredClaims     // 内嵌标准的声明
}

// NewAdminCustomClaims 初始化AdminCustomClaims
func newAppClaims(uid int, expiresAt time.Time, issuer, subject string) AppClaims {
	//now := time.Now()
	return AppClaims{
		Uid: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt), // 定义过期时间
			Issuer:    issuer,                        // 签发人
			//IssuedAt:  jwt.NewNumericDate(now),       // 签发时间
			Subject: subject, // 签发主体
			//NotBefore: jwt.NewNumericDate(now),       // 生效时间
		},
	}
}

func newAdminClaims(userId int, roleId int, expiresAt time.Time, issuer, subject string) AdminClaims {
	//now := time.Now()
	return AdminClaims{
		Uid: userId,
		Rid: roleId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt), // 定义过期时间
			Issuer:    issuer,                        // 签发人
			//IssuedAt:  jwt.NewNumericDate(now),       // 签发时间
			Subject: subject, // 签发主体
			//NotBefore: jwt.NewNumericDate(now),       // 生效时间
		},
	}
}

const (
	TOKEN_APP           = "app"
	TOKEN_APP_REFRESH   = "app_refresh"
	TOKEN_ADMIN         = "admin"
	TOKEN_ADMIN_REFRESH = "admin_refresh"

	AdmUid = "adm_uid"
	AdmRid = "adm_rid"
	AppUid = "app_uid"
)
