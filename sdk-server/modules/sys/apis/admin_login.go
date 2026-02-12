package apis

import (
	"context"
	"dilu/common/ckey"
	"dilu/common/codes"
	"dilu/common/config"
	"dilu/common/third/feishus"
	"dilu/common/utils"
	"dilu/modules/sys/models"
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
	"github.com/jinzhu/copier"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

var ApiAdminLogin = AdminLogin{}

type AdminLogin struct {
	base.BaseApi
}

// 管理员登录
// Login 管理员登录
// @Summary 管理员登录
// @Description 管理员登录
// @Tags sys-AdminLogin
// @Accept application/json
// @Product application/json
// @Param data body dto.LoginReq true "data"
// @Success 200 {object} base.Resp{data=dto.LoginOK} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin/login [post]
func (e *AdminLogin) Login(c *gin.Context) {
	req := dto.LoginReq{}
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	ip := ips.GetIP(c)
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
		if logOk, err := service.SerAdmin.LoginPwd(&req, ip); err != nil {
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
	if logOk, err := service.SerAdmin.LoginCode(&req, ip); err != nil {
		core.Log.Error("AdminLogin", "err", err)
		e.Err(c, err)
		return
	} else {
		e.Ok(c, logOk)
	}
}

// @Summary		刷新token
// @Description	刷新token
// @Tags			sys-AdminLogin
// @Accept			application/json
// @Product		application/json
// @Param			data	body		dto.RefreshTokenReq						true	"data"
// @Success		200		{object}	base.Resp{data=dto.LoginOK}	"{"code": 200, "data": [...]}"
// @Router			/api/sys/admin/refreshToken [post]
func (e *AdminLogin) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenReq
	err := c.ShouldBind(&req)
	if err != nil {
		slog.Error("bind", "err", err)
		e.Error(c, err)
		return
	}

	if logOk, err := service.SerAdmin.RefreshToken(&req, ips.GetIP(c)); err != nil {
		slog.Error(err.Error())
		e.Error(c, err)
		return
	} else {
		e.Ok(c, logOk)
	}
}

// 忘记密码
// ForgetPwd 忘记密码
// @Summary 忘记密码
// @Description 忘记密码
// @Tags sys-AdminLogin
// @Accept application/json
// @Product application/json
// @Param data body dto.ForgetPwdReq true "data"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin/forgetPwd [post]
func (e *AdminLogin) ForgetPwd(c *gin.Context) {
	req := dto.ForgetPwdReq{}
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	if !regexps.CheckPwd(req.Password) {
		e.Code(c, codes.ErrPasswordFMT)
		return
	}

	var mobile, email string
	//是否手机
	if regexps.CheckMobile(req.Username) {
		if !service.SerSms.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
		mobile = req.Username
	} else if regexps.CheckEmail(req.Username) { //是否邮箱

		if !service.SerEmail.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
		email = req.Username
	} else {
		e.Code(c, codes.ErrMobileOrEmail)
		return
	}
	if err := service.SerAdmin.ChangePwd(mobile, email, req.Password); err != nil {
		core.Log.Error("AdminLogin", "err", err)
		e.Error(c, err)
		return
	} else {
		e.Ok(c)
	}

}

// 获取管理员信息
// GetUserInfo 获取管理员信息
// @Summary 获取管理员信息
// @Description 获取管理员信息
// @Tags sys-AdminLogin
// @Param data body base.ReqId true "data"
// @Success 200 {object} base.Resp{data=dto.UserinfoResp} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin/getUserinfo [post]
func (e *AdminLogin) GetUserInfo(c *gin.Context) {
	req := base.ReqId{}
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var object models.Admin

	err := service.SerAdmin.Get(req.Id, &object)
	if err != nil {
		core.Log.Error("AdminLogin", "err", err)
		e.Error(c, err)
		return
	}
	resp := dto.UserinfoResp{}
	if err := copier.Copy(&resp, object); err != nil {
		core.Log.Error("AdminLogin", "err", err)
		e.Error(c, err)
		return
	}

	e.Ok(c, resp)
}

// 获取个人信息
// GetUserInfo 获取个人信息
// @Summary 获取个人信息
// @Description 获取个人信息
// @Tags sys-AdminLogin
// @Success 200 {object} base.Resp{data=dto.MyUserinfoResp} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin/myUserinfo [get]
// @Security Bearer
func (e *AdminLogin) MyUserInfo(c *gin.Context) {
	uid := utils.GetAdminId(c)
	if uid == 0 {
		e.Code(c, codes.InvalidToken_401)
		return
	}
	var object models.Admin
	err := service.SerAdmin.Get(uid, &object)
	if err != nil {
		e.Error(c, err)
		return
	}
	resp := dto.MyUserinfoResp{}
	if err := copier.Copy(&resp, object); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, resp)

}

// 修改密码
// ChangePwd 修改密码
// @Summary 修改密码
// @Description 修改密码
// @Tags sys-AdminLogin
// @Accept application/json
// @Product application/json
// @Param data body dto.ChangePwdReq true "data"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin/changePwd [post]
// @Security Bearer
func (e *AdminLogin) ChangePwd(c *gin.Context) {
	req := dto.ChangePwdReq{}
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if req.NewPassword != req.RePassword {
		e.Code(c, codes.ErrRePassword)
		return
	}
	if !regexps.CheckPwd(req.NewPassword) {
		e.Code(c, codes.ErrPasswordFMT)
		return
	}
	uid := utils.GetAdminId(c)
	if uid == 0 {
		e.Code(c, codes.InvalidToken_401)
		return
	}

	if err := service.SerAdmin.ChangePwdByOld(uid, req.OldPassword, req.NewPassword, req.InviteCode); err != nil {
		e.Err(c, err)
		return
	}
	e.Ok(c)
}

// 绑定手机号或者邮箱
// Bind 绑定手机号或者邮箱
// @Summary 绑定手机号或者邮箱
// @Description 绑定手机号或者邮箱
// @Tags sys-AdminLogin
// @Accept application/json
// @Product application/json
// @Param data body dto.BindReq true "data"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin/auth/bind [post]
// @Security Bearer
func (e *AdminLogin) Bind(c *gin.Context) {
	req := dto.BindReq{}
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	uid := utils.GetAdminId(c)
	if uid == 0 {
		e.Code(c, codes.InvalidToken_401)
		return
	}
	//是否手机
	if regexps.CheckMobile(req.Username) {
		s := service.SysSms{}

		if !s.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
	} else if regexps.CheckEmail(req.Username) { //是否邮箱
		s := service.SysEmail{}

		if !s.Verify(req.Username, req.Code) {
			e.Code(c, codes.ErrVerifyCode)
			return
		}
	} else {
		e.Code(c, codes.ErrMobileOrEmail)
		return
	}

	if err := service.SerAdmin.Bind(uid, &req); err != nil {
		e.Code(c, codes.ErrBind)
		return
	}

	e.Ok(c)
}

// 退出登录
// Logout 退出登录
// @Summary 退出登录
// @Description 退出登录（调用后清空本地token）
// @Tags sys-AdminLogin
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin/auth/logout [post]
// @Security Bearer
func (e *AdminLogin) Logout(c *gin.Context) {
	e.Ok(c)
}

func (e *AdminLogin) LoginWithFeishu(c *gin.Context) {
	state := uuid.NewString()
	// 生成 PKCE 需要的 code verifier
	verifier := oauth2.GenerateVerifier()
	// 将 code verifier 存入 session 中
	core.Cache.Set(ckey.LoginFeishuVerifyKey(state), verifier, 5*time.Minute)
	oauth2C := feishus.OAuth2Config(config.Ext.Feishu.AppId, config.Ext.Feishu.AppSecret)
	url := oauth2C.AuthCodeURL(state, oauth2.S256ChallengeOption(verifier))
	// 管理员点击登录时，重定向到授权页面
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (e *AdminLogin) FeishuLoginCallback(c *gin.Context) {
	ctx := context.Background()
	// 从 session 中获取 state

	state := c.Query("state")
	if state == "" {
		log.Printf("error: %s", c.Query("error"))
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	code := c.Query("code")
	// 如果 code 为空，说明管理员拒绝了授权
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
