package apis

import (
	"context"
	"dilu/common/ckey"
	"dilu/common/codes"
	"dilu/common/config"
	"dilu/common/third/feishus"
	"dilu/common/utils"
	"dilu/modules/sys/enums"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"dilu/common/regexps"

	"github.com/baowk/dilu-core/common/utils/ips"
	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

var ApiSso = SSO{}

type SSO struct {
	base.BaseApi
}

// var once sync.Once

// // pubsubtest
// // SendCode pubsubtest
// // @Summary pubsubtest
// // @Description pubsubtest
// // @Tags app-sso
// // @Product application/json
// // @Param api-sk header string false "api-sk"
// // @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// // @Router /api/app/pbtest [get]
// func (e *SSO) PSTest(c *gin.Context) {
// 	rdb, err := core.CacheRedis()
// 	if err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	once.Do(func() {
// 		sub()
// 	})

// 	msgs := []string{"aaa", "bbb", "ddd"}
// 	for _, v := range msgs {
// 		rdb.Publish(context.Background(), chanName, v)
// 		time.Sleep(time.Second)
// 	}

// 	e.Ok(c)
// }

// const chanName = "cname"

// func sub() error {
// 	rdb, err := core.CacheRedis()
// 	if err != nil {
// 		return err
// 	}

// 	go func() {
// 		ctx := context.Background()
// 		sub := rdb.Subscribe(ctx, chanName)
// 		for {
// 			msg, err := sub.ReceiveMessage(ctx)
// 			if err != nil {
// 				continue
// 			}

// 			fmt.Println(msg.Channel, msg.Payload)
// 		}

// 	}()
// 	return nil
// }

// 发送验证码
// SendCode 发送邮箱或者短信验证码
// @Summary 发送邮箱或者短信验证码
// @Description 发送邮箱或者短信验证码
// @Tags app-sso
// @Accept application/json
// @Product application/json
// @Param api-sk header string false "api-sk"
// @Param data body dto.CodeSendReq true "data"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/app/sendCode [post]
func (e *SSO) SendCode(c *gin.Context) {
	req := dto.CodeSendReq{}
	if err := c.ShouldBind(&req); err != nil {
		return
	}

	// if req.Code == "" || req.UUID == "" {
	// 	e.Err(c, codes.ErrInvalidParameter(e.GetReqId(c), "code or uuid is all nil"))
	// 	return
	// }

	// if !service.Verify(req.UUID, req.Code, true) {
	// 	e.Err(c, errs.Err(codes.CaptchaVerifyErr, e.GetReqId(c), nil))
	// 	return
	// }

	//是否手机
	if regexps.CheckMobile(req.Username) {
		if req.CheckExist {
			var count int64
			service.SerSysUser.CountByPhone(req.Username, &count)
			if count > 0 {
				e.Code(c, codes.PhoneExistErr)
				return
			}
		}
		var tempId string
		service.SerSms.Send(req.Username, tempId)

	} else if regexps.CheckEmail(req.Username) { //是否邮箱
		if req.CheckExist {
			var count int64
			service.SerSysUser.CountByEmail(req.Username, &count)
			if count > 0 {
				e.Code(c, codes.EmailExistErr)
				return
			}
		}
		service.SerEmail.Send(req.Username)
	} else {
		e.Code(c, codes.ErrMobileOrEmail)
		return
	}
	e.Ok(c)
}

// GetOneClickToken 手机一键登录获取Token
// @Summary OneClickLogin 手机一键登录获取Token
// @Tags app-sso
// @Accept application/json
// @Product application/json
// @Param api-sk header string false "api-sk"
// @Param data body dto.GetOneClickTokenReq true "body"
// @Success 200 {object} base.Resp{data=dto.GetOneClickTokenResp} "{"code": 200, "data": [...]}"
// @Router /api/app/getOneClickToken [post]
func (e *SSO) GetOneClickToken(c *gin.Context) {
	var req dto.GetOneClickTokenReq
	err := c.ShouldBind(&req)
	if err != nil {
		slog.Error("bind", "err", err)
		e.Error(c, err)
		return
	}
	token, err := service.SerSysUser.GetOneClickToken(&req)
	if err != nil {
		e.Error(c, err)
		return
	}
	resp := dto.GetOneClickTokenResp{
		VerifyResult: token.VerifyResult,
		VerifyId:     token.VerifyId,
	}
	e.Ok(c, resp)
}

// OneClickLogin 用户手机一键登录
// @Summary OneClickLogin 用户手机一键登录
// @Tags app-sso
// @Accept application/json
// @Product application/json
// @Param api-sk header string false "api-sk"
// @Param data body dto.OneClickLoginReq true "body"
// @Success 200 {object} base.Resp{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// @Router /api/app/oneClick [post]
func (e *SSO) OneClickLogin(c *gin.Context) {
	var req dto.OneClickLoginReq
	err := c.ShouldBind(&req)
	if err != nil {
		slog.Error("bind", "err", err)
		e.Error(c, err)
		return
	}

	if logOk, err := service.SerSysUser.OneClickLogin(&req, ips.GetIP(c), utils.GetAppTeamId(c), utils.GetAppId(c)); err != nil {
		slog.Error(err.Error())
		e.Error(c, err)
		return
	} else {
		e.Ok(c, logOk)
	}
}

// GetFusionToken 获取融合认证token
// @Summary GetFusionToken 获取融合认证token
// @Tags app-sso
// @Accept application/json
// @Product application/json
// @Param api-sk header string false "api-sk"
// @Param data body dto.GetFusionTokenReq true "body"
// @Success 200 {object} base.Resp{data=string} "{"code": 200, "data": [...]}"
// @Router /api/app/getFusionToken [post]
func (e *SSO) GetFusionToken(c *gin.Context) {
	var req dto.GetFusionTokenReq
	err := c.ShouldBind(&req)
	if err != nil {
		slog.Error("bind", "err", err)
		e.Error(c, err)
		return
	}
	token, err := service.SerSysUser.GetFusionToken(&req)
	if err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, token)
}

// VerifyFusionLogin 验证融合认证token
// @Summary VerifyFusionLogin 验证融合认证token
// @Tags app-sso
// @Accept application/json
// @Product application/json
// @Param api-sk header string false "api-sk"
// @Param data body dto.OneClickLoginReq true "body"
// @Success 200 {object} base.Resp{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// @Router /api/app/verifyFusion [post]
func (e *SSO) VerifyFusionLogin(c *gin.Context) {
	var req dto.OneClickLoginReq
	err := c.ShouldBind(&req)
	if err != nil {
		slog.Error("bind", "err", err)
		e.Error(c, err)
		return
	}

	if logOk, err := service.SerSysUser.VerifyFusion(&req, ips.GetIP(c), utils.GetAppTeamId(c), utils.GetAppId(c)); err != nil {
		slog.Error(err.Error())
		e.Error(c, err)
		return
	} else {
		e.Ok(c, logOk)
	}
}

// 用户注册
// Register 用户注册
// @Summary 用户注册
// @Description 用户注册
// @Tags app-sso
// @Accept application/json
// @Product application/json
// @Param api-sk header string false "api-sk"
// @Param data body dto.RegisterReq true "data"
// @Success 200 {object} base.Resp{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// @Router /api/app/register [post]
func (e *SSO) Register(c *gin.Context) {
	req := dto.RegisterReq{}
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if req.Password != req.RePassword {
		e.Code(c, codes.ErrRePassword)
		return
	}
	//密码规则
	if req.Password != "" && !regexps.CheckPwd(req.Password) {
		e.Code(c, codes.ErrPasswordFMT)
		return
	}

	loginType := enums.LT_PHONE
	//是否手机
	if regexps.CheckMobile(req.Username) {
		if !service.SerSms.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
	} else if regexps.CheckEmail(req.Username) { //是否邮箱
		if !service.SerEmail.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
		loginType = enums.LT_EMAIL
	} else {
		e.Code(c, codes.ErrMobileOrEmail)
		return
	}

	ip := ips.GetIP(c)
	if logOk, err := service.SerSysUser.Register(loginType, &req, ip, utils.GetAppTeamId(c), utils.GetAppId(c)); err != nil {
		core.Log.Error("sso", "register", err)
		e.Error(c, err)
		return
	} else {
		e.Ok(c, logOk)
	}
}

// // 验证码校验
// // VerifyCode 验证码校验
// // @Summary 验证码校验
// // @Description 验证码校验
// // @Tags app-sso
// // @Accept application/json
// // @Product application/json
// // @Param api-sk header string false "api-sk"
// // @Param data body dto.VerifyCodeReq true "data"
// // @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// // @Router /api/app/verify/code [post]
// func (e *SSO) VerifyCode(c *gin.Context) {
// 	req := dto.VerifyCodeReq{}
// 	if err := c.ShouldBind(&req); err != nil {
// 		e.Error(c, err)
// 		return
// 	}

// 	//是否手机
// 	if regexps.CheckMobile(req.Username) {
// 		if !service.SerSms.Verify(req.Username, req.Code) {
// 			e.Code(c, codes.ErrVerifyCode)
// 			return
// 		}
// 	} else if regexps.CheckEmail(req.Username) { //是否邮箱
// 		if !service.SerEmail.Verify(req.Username, req.Code) {
// 			e.Code(c, codes.ErrVerifyCode)
// 			return
// 		}
// 	} else {
// 		e.Code(c, codes.ErrMobileOrEmail)
// 		return
// 	}
// 	e.Ok(c)
// }

// 用户登录
// Login 用户登录
// @Summary 用户登录
// @Description 用户登录
// @Tags app-sso
// @Accept application/json
// @Product application/json
// @Param data body dto.LoginReq true "data"
// @Success 200 {object} base.Resp{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// @Router /api/app/login [post]
func (e *SSO) Login(c *gin.Context) {
	req := dto.LoginReq{}
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	ip := ips.GetIP(c)
	appId := utils.GetAppId(c)
	teamId := utils.GetAppTeamId(c)
	if req.Password == "" {
		//是否手机
		if regexps.CheckMobile(req.Username) {
			if !service.SerSms.Verify(req.Username, req.Code) {
				e.Code(c, codes.ErrVerifyCode)
				return
			}
		} else if regexps.CheckEmail(req.Username) { //是否邮箱
			if !service.SerEmail.Verify(req.Username, req.Code) {
				e.Code(c, codes.ErrVerifyCode)
				return
			}
			// } else {
			// 	e.Code(c, codes.ErrMobileOrEmail)
			// 	return
		}
	} else {
		if logOk, err := service.SerSysUser.LoginPwd(&req, ip, teamId, appId); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				e.Code(c, codes.UserNotExist)
				return
			} else {
				e.Err(c, err)
				return
			}
		} else {
			e.Ok(c, logOk)
			return
		}
	}
	if logOk, err := service.SerSysUser.LoginCode(&req, ip, teamId, appId); err != nil {
		core.Log.Error("sso", "err", err)
		e.Err(c, err)
		return
	} else {
		e.Ok(c, logOk)
	}
}

// @Summary		刷新token
// @Description	刷新token
// @Tags	    app-sso
// @Accept		application/json
// @Product		application/json
// @Param api-sk header string false "api-sk"
// @Param		data	body		dto.RefreshTokenReq						true	"data"
// @Success		200		{object}	base.Resp{data=dto.LoginOK}	"{"code": 200, "data": [...]}"
// @Router		/api/app/refreshToken [post]
func (e *SSO) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenReq
	err := c.ShouldBind(&req)
	if err != nil {
		slog.Error("bind", "err", err)
		e.Error(c, err)
		return
	}

	if logOk, err := service.SerSysUser.RefreshToken(&req, ips.GetIP(c), utils.GetAppTeamId(c), utils.GetAppId(c)); err != nil {
		slog.Error(err.Error())
		e.Error(c, err)
		return
	} else {
		e.Ok(c, logOk)
	}
}

// 绑定手机号或者邮箱
// Bind 绑定手机号或者邮箱
// @Summary 绑定手机号或者邮箱
// @Description 绑定手机号或者邮箱
// @Tags app-sso
// @Accept application/json
// @Product application/json
// @Param api-sk header string false "api-sk"
// @Param data body dto.BindReq true "data"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/app/auth/bind [post]
// @Security Bearer
func (e *SSO) Bind(c *gin.Context) {
	var req dto.BindReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	uid := utils.GetAppUid(c)
	if uid == 0 {
		e.Code(c, codes.InvalidToken_401)
		return
	}
	//是否手机
	if regexps.CheckMobile(req.Username) {
		if !service.SerSms.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
	} else if regexps.CheckEmail(req.Username) { //是否邮箱
		if !service.SerEmail.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
	} else {
		e.Code(c, codes.ErrMobileOrEmail)
		return
	}

	if err := service.SerSysUser.Bind(uid, &req); err != nil {
		e.Code(c, codes.ErrBind)
		return
	}

	e.Ok(c)
}

// 退出登录
// Logout 退出登录
// @Summary 退出登录
// @Description 退出登录（调用后清空本地token）
// @Tags app-sso
// @Param api-sk header string false "api-sk"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/app/auth/logout [post]
// @Security Bearer
func (e *SSO) Logout(c *gin.Context) {

	e.Ok(c)
}

func (e *SSO) LoginWithFeishu(c *gin.Context) {
	state := uuid.NewString()
	// 生成 PKCE 需要的 code verifier
	verifier := oauth2.GenerateVerifier()
	// 将 code verifier 存入 session 中
	core.Cache.Set(ckey.LoginFeishuVerifyKey(state), verifier, 5*time.Minute)
	oauth2C := feishus.OAuth2Config(config.Ext.Feishu.AppId, config.Ext.Feishu.AppSecret)
	url := oauth2C.AuthCodeURL(state, oauth2.S256ChallengeOption(verifier))
	// 用户点击登录时，重定向到授权页面
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (e *SSO) FeishuLoginCallback(c *gin.Context) {
	ctx := context.Background()
	// 从 session 中获取 state

	state := c.Query("state")
	if state == "" {
		log.Printf("error: %s", c.Query("error"))
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	code := c.Query("code")
	// 如果 code 为空，说明用户拒绝了授权
	if code == "" {
		log.Printf("error: %s", c.Query("error"))
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	codeVerifier, err := core.Cache.Get(ckey.LoginFeishuVerifyKey(state))
	if err != nil {
		log.Printf("error: %s", c.Query("error"))
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	oauth2C := feishus.OAuth2Config(config.Ext.Feishu.AppId, config.Ext.Feishu.AppSecret)
	// 使用获取到的 code 获取 token
	token, err := oauth2C.Exchange(ctx, code, oauth2.VerifierOption(codeVerifier))
	if err != nil {
		log.Printf("oauthConfig.Exchange() failed with '%s'\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	cli := feishus.NewClient(config.Ext.Feishu.AppId, config.Ext.Feishu.AppSecret)
	fs, err := cli.GetUserInfo(token.AccessToken)
	if err != nil {

	}

	fmt.Println(fs)

}

// // 微信登录
// // LoginByWechat 微信登录
// // @Summary 微信登录
// // @Description 微信登录
// // @Tags app-sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.LoginWechatReq true "data"
// // @Success 200 {object} base.Resp{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// // @Router /api/app/loginWechat [post]
// func (e *SSO) LoginByWechat(c *gin.Context) {
// 	req := dto.LoginWechatReq{}
// 	s := service.User{}
// 	err := e.MakeContext(c).MakeOrm().
// 		Bind(&req).MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		core.Log.Error(err)
// 		e.Code(500, err, err.Error())
// 		return
// 	}

// 	ip := common.GetClientIP(c)

// 	if logOk, err := s.LoginWechat(&req, ip); err != nil {
// 		core.Log.Error(err)
// 		e.Code(500, err, fmt.Sprintf("登录失败，\r\n失败信息 %s", err.Error()))
// 		return
// 	} else {
// 		e.Ok(c,logOk)
// 	}
// }

// // 获取钉钉登录配置信息
// // GeDingCfg 获取钉钉登录配置信息
// // @Summary 获取钉钉登录配置信息
// // @Description 获取钉钉登录配置信息
// // @Tags app-sso
// // @Success 200 {object} base.Resp{data=dto.DingCfgResp} "{"code": 200, "data": [...]}"
// // @Router /api/app/getDingCfg [post]
// func (e *SSO) GetDingCfg(c *gin.Context) {
// 	err := e.MakeContext(c).MakeOrm().
// 		Errors
// 	if err != nil {
// 		core.Log.Error(err)
// 		e.Code(500, err, err.Error())
// 		return
// 	}
// 	domain := "http://" + c.Request.Host
// 	var cfg dto.DingCfgResp
// 	service.GetDingConfig(domain, &cfg)
// 	e.Ok(c,cfg)
// }

// // 钉钉
// // LoginByDing 钉钉登录
// // @Summary 钉钉
// // @Description 钉钉登录
// // @Tags app-sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.LoginDingReq true "data"
// // @Success 200 {object} base.Resp{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// // @Router /api/app/loginDing [post]
// func (e *SSO) LoginByDing(c *gin.Context) {
// 	req := dto.LoginDingReq{}
// 	s := service.User{}
// 	err := e.MakeContext(c).MakeOrm().
// 		Bind(&req).MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		core.Log.Error(err)
// 		e.Code(500, err, err.Error())
// 		return
// 	}

// 	ip := common.GetClientIP(c)

// 	if logOk, err := s.LoginDing(&req, ip); err != nil {
// 		core.Log.Error(err)
// 		e.Code(500, err, fmt.Sprintf("登录失败，\r\n失败信息 %s", err.Error()))
// 		return
// 	} else {
// 		e.Ok(c,logOk)
// 	}

// }

// // 绑定钉钉
// // BindDing 绑定钉钉
// // @Summary 绑定钉钉
// // @Description 绑定钉钉
// // @Tags app-sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.LoginDingReq true "data"
// // @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// // @Router /api/app/auth/bindDing [post]
// // @Security Bearer
// func (e *SSO) BindDing(c *gin.Context) {
// 	req := dto.LoginDingReq{}
// 	s := service.User{}
// 	err := e.MakeContext(c).
// 		Bind(&req).
// 		MakeOrm().
// 		MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		core.Log.Error(err)
// 		e.Code(500, err, err.Error())
// 		return
// 	}
// 	if e.GetUserId() == "" {
// 		e.Code(c, codes.ErrUnLogin)
// 		return
// 	}

// 	if err := s.BindDing(&req, e.GetUserId()); err != nil {
// 		e.Code(c, codes.ErrBind)
// 		return
// 	}
// 	e.Ok(c)
// }

// // 绑定微信
// // BindWechat 绑定微信
// // @Summary 绑定微信
// // @Description 绑定微信
// // @Tags app-sso
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.LoginWechatReq true "data"
// // @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// // @Router /api/app/auth/bindWechat [post]
// // @Security Bearer
// func (e *SSO) BindWechat(c *gin.Context) {
// 	req := dto.LoginWechatReq{}
// 	s := service.User{}
// 	err := e.MakeContext(c).
// 		Bind(&req).
// 		MakeOrm().
// 		MakeService(&s.Service).
// 		Errors
// 	if err != nil {
// 		core.Log.Error(err)
// 		e.Code(500, err, err.Error())
// 		return
// 	}
// 	if e.GetUserId() == "" {
// 		e.Code(c, codes.ErrUnLogin)
// 		return
// 	}

// 	if err := s.BindWechat(&req, e.GetUserId()); err != nil {
// 		e.Code(c, codes.ErrBind)
// 		return
// 	}
// 	e.Ok(c)
// }
