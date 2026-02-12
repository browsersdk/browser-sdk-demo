package apis

import (
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type SysSmsApi struct {
	base.BaseApi
}

var ApiSysSms = SysSmsApi{}

// QueryPage 获取短信列表
// @Summary 获取短信列表
// @Tags sys-SysSms
// @Accept application/json
// @Product application/json
// @Param data body dto.SysSmsGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysSms}} "{"code": 200, "data": [...]}"
// @Router /api/sys/sys-sms/page [post]
// @Security Bearer
func (e *SysSmsApi) QueryPage(c *gin.Context) {
	var req dto.SysSmsGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysSms, 0, req.GetSize())
	var total int64

	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}

	if err := service.SerSms.QueryPage(req, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取短信
// @Summary 获取短信
// @Tags sys-SysSms
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysSms} "{"code": 200, "data": [...]}"
// @Router /api/sys/sys-sms/get [post]
// @Security Bearer
func (e *SysSmsApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysSms
	if err := service.SerSms.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建短信
// @Summary 创建短信
// @Tags sys-SysSms
// @Accept application/json
// @Product application/json
// @Param data body dto.SysSmsDto true "body"
// @Success 200 {object} base.Resp{data=models.SysSms} "{"code": 200, "data": [...]}"
// @Router /api/sys/sys-sms/create [post]
// @Security Bearer
func (e *SysSmsApi) Create(c *gin.Context) {
	var req dto.SysSmsDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysSms
	copier.Copy(&data, req)
	if err := service.SerSms.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新短信
// @Summary 更新短信
// @Tags sys-SysSms
// @Accept application/json
// @Product application/json
// @Param data body dto.SysSmsDto true "body"
// @Success 200 {object} base.Resp{data=models.SysSms} "{"code": 200, "data": [...]}"
// @Router /api/sys/sys-sms/update [post]
// @Security Bearer
func (e *SysSmsApi) Update(c *gin.Context) {
	var req dto.SysSmsDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysSms
	copier.Copy(&data, req)
	if err := service.SerSms.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除短信
// @Summary 删除短信
// @Tags sys-SysSms
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysSms} "{"code": 200, "data": [...]}"
// @Router /api/sys/sys-sms/del [post]
// @Security Bearer
func (e *SysSmsApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSms.DelIds(&models.SysSms{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
