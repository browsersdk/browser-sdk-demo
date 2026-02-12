package apis

import (
	"dilu/common/codes"
	"dilu/common/consts"
	"dilu/common/utils"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	sService "dilu/modules/sys/service"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
)

type AppUser struct {
	base.BaseApi
}

var ApiAppUser = &AppUser{}

// 获取用户信息接口
// @Summary 获取用户信息接口
// @Description 获取用户信息接口
// @Tags app-user
// @Accept application/json
// @Product application/json
// @Success 200 {object} base.Resp{data=models.SysUser} "{"code": 200, "data": [...]}"
// @Router /api/app/user/info [get]
// @Security Bearer
func (e *AppUser) GetUserInfo(c *gin.Context) {

	uid := utils.GetAppUid(c)
	if uid == 0 {
		e.Code(c, codes.InvalidToken_401)
		return
	}
	var user models.SysUser

	if err := service.SerSysUser.Get(uid, &user); err != nil {
		e.Code(c, codes.ErrBind)
		return
	}

	e.Ok(c, user)
}

// 用户上传文件
// @Summary 用户上传文件
// @Tags app-user
// @Accept multipart/form-data
// @Product application/json
// @Param file formData file true "文件"
// @Success 200 {object} base.Resp{data=string} "{"code": 200, "data": [...]}"
// @Router /api/app/user/upload [post]
// @Security Bearer
func (e *AppUser) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		e.Error(c, err)
		return
	}
	url, err := sService.SerUploadLog.UploadFile(file, consts.UPLOAD_USER)
	if err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, url)
}
