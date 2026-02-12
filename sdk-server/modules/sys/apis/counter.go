package apis

import (
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type CounterApi struct {
	base.BaseApi
}

var ApiCounter = CounterApi{}

// QueryPage 获取计数器列表
// @Summary 获取计数器列表
// @Tags sys-Counter
// @Accept application/json
// @Product application/json
// @Param data body dto.CounterGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.Counter}} "{"code": 200, "data": [...]}"
// @Router /api/sys/counter/page [post]
// @Security Bearer
func (e *CounterApi) QueryPage(c *gin.Context) {
	var req dto.CounterGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.Counter, 0, req.GetSize())
	var total int64

	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}

	if err := service.SerCounter.QueryPage(req, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取计数器
// @Summary 获取计数器
// @Tags sys-Counter
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.Counter} "{"code": 200, "data": [...]}"
// @Router /api/sys/counter/get [post]
// @Security Bearer
func (e *CounterApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Counter
	if err := service.SerCounter.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建计数器
// @Summary 创建计数器
// @Tags sys-Counter
// @Accept application/json
// @Product application/json
// @Param data body dto.CounterDto true "body"
// @Success 200 {object} base.Resp{data=models.Counter} "{"code": 200, "data": [...]}"
// @Router /api/sys/counter/create [post]
// @Security Bearer
func (e *CounterApi) Create(c *gin.Context) {
	var req dto.CounterDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Counter
	copier.Copy(&data, req)
	if err := service.SerCounter.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新计数器
// @Summary 更新计数器
// @Tags sys-Counter
// @Accept application/json
// @Product application/json
// @Param data body dto.CounterDto true "body"
// @Success 200 {object} base.Resp{data=models.Counter} "{"code": 200, "data": [...]}"
// @Router /api/sys/counter/update [post]
// @Security Bearer
func (e *CounterApi) Update(c *gin.Context) {
	var req dto.CounterDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Counter
	copier.Copy(&data, req)
	if err := service.SerCounter.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除计数器
// @Summary 删除计数器
// @Tags sys-Counter
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.Counter} "{"code": 200, "data": [...]}"
// @Router /api/sys/counter/del [post]
// @Security Bearer
func (e *CounterApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerCounter.DelIds(&models.Counter{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
