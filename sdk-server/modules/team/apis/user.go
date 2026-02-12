package apis

import (
	"dilu/common/codes"
	"dilu/common/regexps"
	"dilu/common/utils"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type SysUserApi struct {
	base.BaseApi
}

var ApiSysUser = SysUserApi{}

// @Summary Page接口
// @Tags Team
// @Accept application/json
// @Product application/json
// @Param data body dto.SysUserGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysUser}} "{"code": 200, "data": [...]}"
// @Router /api/team/user/page [post]
// @Security Bearer
func (e *SysUserApi) QueryPage(c *gin.Context) {
	var req dto.SysUserGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	if regexps.CheckMobile(req.Username) {
		req.Phone = req.Username
		req.Username = ""
	} else if regexps.CheckEmail(req.Username) {
		req.Email = req.Username
		req.Username = ""
	}

	list := make([]models.SysUser, 10)
	var total int64
	var model models.SysUser
	if err := copier.Copy(&model, req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysUser.Page(&req, &list, &total); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// 忘记密码
// ForgetPwd 忘记密码
// @Summary 忘记密码
// @Description 忘记密码
// @Tags sys-sso
// @Accept application/json
// @Product application/json
// @Param data body dto.ForgetPwdReq true "data"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/sys/forgetPwd [post]
func (e *SSO) ForgetPwd(c *gin.Context) {
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
	if err := service.SerSysUser.ChangePwd(mobile, email, req.Password); err != nil {
		core.Log.Error("sso", "err", err)
		e.Error(c, err)
		return
	} else {
		e.Ok(c)
	}

}

// 获取个人信息
// GetUserInfo 获取个人信息
// @Summary 获取个人信息
// @Description 获取个人信息
// @Tags sys-sso
// @Success 200 {object} base.Resp{data=dto.MyUserinfoResp} "{"code": 200, "data": [...]}"
// @Router /api/sys/myUserinfo [get]
// @Security Bearer
func (e *SSO) MyUserInfo(c *gin.Context) {
	uid := utils.GetAppUid(c)
	if uid == 0 {
		e.Code(c, codes.InvalidToken_401)
		return
	}
	var object models.SysUser
	err := service.SerSysUser.Get(uid, &object)
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
// @Tags sys-sso
// @Accept application/json
// @Product application/json
// @Param data body dto.ChangePwdReq true "data"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/sys/changePwd [post]
// @Security Bearer
func (e *SSO) ChangePwd(c *gin.Context) {
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
	uid := utils.GetAppUid(c)
	if uid == 0 {
		e.Code(c, codes.InvalidToken_401)
		return
	}

	if err := service.SerSysUser.ChangePwdByOld(uid, req.OldPassword, req.NewPassword, req.InviteCode); err != nil {
		e.Err(c, err)
		return
	}
	e.Ok(c)
}
