package apis

import (
	"dilu/common/utils"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
)

type AdminRoleApi struct {
	base.BaseApi
}

var ApiAdminRole = AdminRoleApi{}

// QueryPage 获取角色列表
// @Summary 获取角色列表
// @Tags sys-AdminRole
// @Accept application/json
// @Product application/json
// @Param data body dto.AdminRoleGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.AdminRole}} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin-role/page [post]
// @Security Bearer
func (e *AdminRoleApi) QueryPage(c *gin.Context) {
	var req dto.AdminRoleGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.AdminRole, 0, req.GetSize())
	var total int64

	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}

	if err := service.SerAdminRole.QueryPage(req, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// List 获取角色列表
// @Summary 获取角色列表
// @Tags sys-Role
// @Accept application/json
// @Product application/json
// @Param data body dto.SysRoleGetPageReq true "body"
// @Success 200 {object} base.Resp{data=[]models.SysRole} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin-role/list [post]
// @Security Bearer
func (e *AdminRoleApi) List(c *gin.Context) {
	var req dto.SysRoleGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.AdminRole, 0)

	if err := service.SerAdminRole.Query(req.Status, &list); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, list)
}

// Get 获取角色
// @Summary 获取角色
// @Tags sys-AdminRole
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.AdminRole} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin-role/get [post]
// @Security Bearer
func (e *AdminRoleApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.AdminRole
	if err := service.SerAdminRole.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建角色
// @Summary 创建角色
// @Tags sys-AdminRole
// @Accept application/json
// @Product application/json
// @Param data body dto.AdminRoleDto true "body"
// @Success 200 {object} base.Resp{data=models.AdminRole} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin-role/create [post]
// @Security Bearer
func (e *AdminRoleApi) Create(c *gin.Context) {
	var req dto.AdminRoleDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	admin := utils.GetAdminId(c)
	if err := service.SerAdminRole.Create(admin, &req); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, "")
}

// Update 更新角色
// @Summary 更新角色
// @Tags sys-AdminRole
// @Accept application/json
// @Product application/json
// @Param data body dto.AdminRoleDto true "body"
// @Success 200 {object} base.Resp{data=models.AdminRole} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin-role/update [post]
// @Security Bearer
func (e *AdminRoleApi) Update(c *gin.Context) {
	var req dto.AdminRoleDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	admin := utils.GetAdminId(c)
	if err := service.SerAdminRole.Update(admin, &req); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, "")
}

// Del 删除角色
// @Summary 删除角色
// @Tags sys-AdminRole
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.AdminRole} "{"code": 200, "data": [...]}"
// @Router /api/sys/admin-role/del [post]
// @Security Bearer
func (e *AdminRoleApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	admin := utils.GetAdminId(c)
	for _, id := range req.Ids {
		if err := service.SerAdminRole.Del(id, admin); err != nil {
			e.Error(c, err)
			return
		}
	}
	e.Ok(c)
}
