package apis

import (
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
)

type SysUserLoginLogApi struct {
	base.BaseApi
}

var ApiSysUserLoginLog = SysUserLoginLogApi{}

// QueryPage 获取SysUserLoginLog列表
// @Summary 获取SysUserLoginLog列表
// @Tags sys-UserLoginLog
// @Accept application/json
// @Product application/json
// @Param data body dto.SysUserLoginLogGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysUserLoginLog}} "{"code": 200, "data": [...]}"
// @Router /api/sys/user-login-log/page [post]
// @Security Bearer
func (e *SysUserLoginLogApi) QueryPage(c *gin.Context) {
	var req dto.SysUserLoginLogGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysUserLoginLog, 0, req.GetSize())
	var total int64

	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}

	if err := service.SerSysUserLoginLog.QueryPage(req, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取SysUserLoginLog
// @Summary 获取SysUserLoginLog
// @Tags sys-UserLoginLog
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysUserLoginLog} "{"code": 200, "data": [...]}"
// @Router /api/sys/user-login-log/:id [get]
// @Security Bearer
func (e *SysUserLoginLogApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysUserLoginLog
	if err := service.SerSysUserLoginLog.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// // Create 创建SysUserLoginLog
// // @Summary 创建SysUserLoginLog
// // @Tags sys-UserLoginLog
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.SysUserLoginLogDto true "body"
// // @Success 200 {object} base.Resp{data=models.SysUserLoginLog} "{"code": 200, "data": [...]}"
// // @Router /api/sys/user-login-log/create [post]
// // @Security Bearer
// func (e *SysUserLoginLogApi) Create(c *gin.Context) {
// 	var req dto.SysUserLoginLogDto
// 	if err := c.ShouldBind(&req); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	var data models.SysUserLoginLog
// 	copier.Copy(&data, req)
// 	if err := service.SerSysUserLoginLog.Create(&data); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	e.Ok(c, data)
// }

// // Update 更新SysUserLoginLog
// // @Summary 更新SysUserLoginLog
// // @Tags sys-UserLoginLog
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.SysUserLoginLogDto true "body"
// // @Success 200 {object} base.Resp{data=models.SysUserLoginLog} "{"code": 200, "data": [...]}"
// // @Router /api/sys/user-login-log/update [post]
// // @Security Bearer
// func (e *SysUserLoginLogApi) Update(c *gin.Context) {
// 	var req dto.SysUserLoginLogDto
// 	if err := c.ShouldBind(&req); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	var data models.SysUserLoginLog
// 	copier.Copy(&data, req)
// 	if err := service.SerSysUserLoginLog.UpdateById(&data); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	e.Ok(c, data)
// }

// // Del 删除SysUserLoginLog
// // @Summary 删除SysUserLoginLog
// // @Tags sys-UserLoginLog
// // @Accept application/json
// // @Product application/json
// // @Param data body base.ReqIds true "body"
// // @Success 200 {object} base.Resp{data=models.SysUserLoginLog} "{"code": 200, "data": [...]}"
// // @Router /api/sys/user-login-log/del [post]
// // @Security Bearer
// func (e *SysUserLoginLogApi) Del(c *gin.Context) {
// 	var req base.ReqIds
// 	if err := c.ShouldBind(&req); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	if err := service.SerSysUserLoginLog.DelIds(&models.SysUserLoginLog{}, req.Ids); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	e.Ok(c)
// }
