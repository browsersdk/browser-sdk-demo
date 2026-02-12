package apis

import (
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type SysEmailApi struct {
	base.BaseApi
}

var ApiSysEmail = SysEmailApi{}

// QueryPage 获取邮件列表
// @Summary 获取邮件列表
// @Tags sys-SysEmail
// @Accept application/json
// @Product application/json
// @Param data body dto.SysEmailGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysEmail}} "{"code": 200, "data": [...]}"
// @Router /api/sys/sys-email/page [post]
// @Security Bearer
func (e *SysEmailApi) QueryPage(c *gin.Context) {
	var req dto.SysEmailGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysEmail, 0, req.GetSize())
	var total int64

	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}

	if err := service.SerEmail.QueryPage(req, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取邮件
// @Summary 获取邮件
// @Tags sys-SysEmail
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysEmail} "{"code": 200, "data": [...]}"
// @Router /api/sys/sys-email/get [post]
// @Security Bearer
func (e *SysEmailApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysEmail
	if err := service.SerEmail.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建邮件
// @Summary 创建邮件
// @Tags sys-SysEmail
// @Accept application/json
// @Product application/json
// @Param data body dto.SysEmailDto true "body"
// @Success 200 {object} base.Resp{data=models.SysEmail} "{"code": 200, "data": [...]}"
// @Router /api/sys/sys-email/create [post]
// @Security Bearer
func (e *SysEmailApi) Create(c *gin.Context) {
	var req dto.SysEmailDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysEmail
	copier.Copy(&data, req)
	if err := service.SerEmail.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新邮件
// @Summary 更新邮件
// @Tags sys-SysEmail
// @Accept application/json
// @Product application/json
// @Param data body dto.SysEmailDto true "body"
// @Success 200 {object} base.Resp{data=models.SysEmail} "{"code": 200, "data": [...]}"
// @Router /api/sys/sys-email/update [post]
// @Security Bearer
func (e *SysEmailApi) Update(c *gin.Context) {
	var req dto.SysEmailDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysEmail
	copier.Copy(&data, req)
	if err := service.SerEmail.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除邮件
// @Summary 删除邮件
// @Tags sys-SysEmail
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysEmail} "{"code": 200, "data": [...]}"
// @Router /api/sys/sys-email/del [post]
// @Security Bearer
func (e *SysEmailApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerEmail.DelIds(&models.SysEmail{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
