package apis

import (
	"dilu/common/utils"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
)

type AdminDeptApi struct {
	base.BaseApi
}

var ApiAdminDept = AdminDeptApi{}

// QueryPage 获取部门列表
// @Summary 获取部门列表
// @Tags sys-AdminDept
// @Accept application/json
// @Product application/json
// @Param data body dto.AdminDeptGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.AdminDept}} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin-dept/page [post]
// @Security Bearer
func (e *AdminDeptApi) QueryPage(c *gin.Context) {
	var req dto.AdminDeptGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.AdminDept, 0, req.GetSize())
	var total int64

	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}

	if err := service.SerAdminDept.QueryPage(req, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// List 获取全部部门
// @Summary 获取全部部门
// @Tags sys-Dept
// @Accept application/json
// @Product application/json
// @Success 200 {object} base.Resp{data=[]models.AdminDept} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin-dept/all [post]
// @Security Bearer
func (e *AdminDeptApi) List(c *gin.Context) {

	list := make([]models.AdminDept, 10)

	if err := service.SerAdminDept.List(&list); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, list)
}

// Get 获取部门
// @Summary 获取部门
// @Tags sys-AdminDept
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.AdminDept} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin-dept/get [post]
// @Security Bearer
func (e *AdminDeptApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.AdminDept
	if err := service.SerAdminDept.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建部门
// @Summary 创建部门
// @Tags sys-AdminDept
// @Accept application/json
// @Product application/json
// @Param data body dto.AdminDeptDto true "body"
// @Success 200 {object} base.Resp{data=models.AdminDept} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin-dept/create [post]
// @Security Bearer
func (e *AdminDeptApi) Create(c *gin.Context) {
	var req dto.AdminDeptDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	admin := utils.GetAdminId(c)
	if data, err := service.SerAdminDept.Create(req, admin); err != nil {
		e.Error(c, err)
		return
	} else {
		e.Ok(c, data)
	}
}

// Update 更新部门
// @Summary 更新部门
// @Tags sys-AdminDept
// @Accept application/json
// @Product application/json
// @Param data body dto.AdminDeptDto true "body"
// @Success 200 {object} base.Resp{data=models.AdminDept} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin-dept/update [post]
// @Security Bearer
func (e *AdminDeptApi) Update(c *gin.Context) {
	var req dto.AdminDeptDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	admin := utils.GetAdminId(c)
	if data, err := service.SerAdminDept.Update(req, admin); err != nil {
		e.Error(c, err)
		return
	} else {
		e.Ok(c, data)
	}
}

// Del 删除部门
// @Summary 删除部门
// @Tags sys-AdminDept
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.AdminDept} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin-dept/del [post]
// @Security Bearer
func (e *AdminDeptApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerAdminDept.DelIds(&models.AdminDept{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
