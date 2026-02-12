package apis

import (
	"dilu/common/utils"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type AdminApi struct {
	base.BaseApi
}

var ApiAdmin = AdminApi{}

// QueryPage 获取用户列表
// @Summary 获取用户列表
// @Tags sys-Admin
// @Accept application/json
// @Product application/json
// @Param data body dto.AdminGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.Admin}} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin/page [post]
// @Security Bearer
func (e *AdminApi) QueryPage(c *gin.Context) {
	var req dto.AdminGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.Admin, 0, req.GetSize())
	var total int64

	if err := service.SerAdmin.QueryPage(&req, &list, &total); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取用户
// @Summary 获取用户
// @Tags sys-Admin
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.Admin} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin/get [post]
// @Security Bearer
func (e *AdminApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Admin
	if err := service.SerAdmin.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建用户
// @Summary 创建用户
// @Tags sys-Admin
// @Accept application/json
// @Product application/json
// @Param data body dto.AdminDto true "body"
// @Success 200 {object} base.Resp{data=models.Admin} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin/create [post]
// @Security Bearer
func (e *AdminApi) Create(c *gin.Context) {
	var req dto.AdminDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Admin
	copier.Copy(&data, req)
	if err := service.SerAdmin.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新用户
// @Summary 更新用户
// @Tags sys-Admin
// @Accept application/json
// @Product application/json
// @Param data body dto.AdminDto true "body"
// @Success 200 {object} base.Resp{data=models.Admin} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin/update [post]
// @Security Bearer
func (e *AdminApi) Update(c *gin.Context) {
	var req dto.AdminDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Admin
	copier.Copy(&data, req)
	if err := service.SerAdmin.Update(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除用户
// @Summary 删除用户
// @Tags sys-Admin
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.Admin} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin/del [post]
// @Security Bearer
func (e *AdminApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	admin := utils.GetAdminId(c)
	for _, id := range req.Ids {
		if err := service.SerAdmin.Remove(id, admin); err != nil {
			e.Error(c, err)
			return
		}
	}
	e.Ok(c)
}
