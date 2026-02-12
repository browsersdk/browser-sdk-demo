package apis

import (
	"dilu/common/codes"
	"dilu/common/utils"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
)

type SysRoleApi struct {
	base.BaseApi
}

var ApiSysRole = SysRoleApi{}

// QueryPage 获取SysRole列表
// @Summary Page接口
// @Tags sys-Role
// @Accept application/json
// @Product application/json
// @Param data body dto.SysRoleGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysRole}} "{"code": 200, "data": [...]}"
// @Router /api/sys/role/page [post]
// @Security Bearer
func (e *SysRoleApi) QueryPage(c *gin.Context) {
	var req dto.SysRoleGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysRole, 0)
	var total int64
	if req.TeamId == 0 {
		e.Error(c, codes.ErrInvalidParameter("", "teamId is required"))
		return
	}

	if err := service.SerSysRole.Page(&req, &list, &total); err != nil {
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
// @Router /api/sys/role/list [post]
// @Security Bearer
func (e *SysRoleApi) List(c *gin.Context) {
	var req dto.SysRoleGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysRole, 0)

	if req.TeamId == 0 {
		e.Error(c, codes.ErrInvalidParameter("", "teamId is required"))
		return
	}

	if err := service.SerSysRole.Query(req.TeamId, req.Status, &list); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, list)
}

// Get 获取SysRole
// @Summary 获取SysRole
// @Tags sys-Role
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=dto.SysRoleDtoResp} "{"code": 200, "data": [...]}"
// @Router /api/sys/role/get [post]
// @Security Bearer
func (e *SysRoleApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data dto.SysRoleDtoResp
	if err := service.SerSysRole.GetRole(req.Id, utils.GetAdminId(c), utils.GetTeamId(c), &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建SysRole
// @Summary 创建SysRole
// @Tags sys-Role
// @Accept application/json
// @Product application/json
// @Param data body dto.SysRoleDto true "body"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/sys/role/create [post]
// @Security Bearer
func (e *SysRoleApi) Create(c *gin.Context) {
	var req dto.SysRoleDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if req.TeamId == 0 {
		e.Error(c, codes.ErrInvalidParameter("", "teamId is required"))
		return
	}
	if err := service.SerSysRole.Create(utils.GetAdminId(c), req.TeamId, req); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}

// Update 更新SysRole
// @Summary 更新SysRole
// @Tags sys-Role
// @Accept application/json
// @Product application/json
// @Param data body dto.SysRoleDto true "body"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/sys/role/update [post]
// @Security Bearer
func (e *SysRoleApi) Update(c *gin.Context) {
	var req dto.SysRoleDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysRole.Update(utils.GetAdminId(c), req.TeamId, req); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}

// Del 删除SysRole
// @Summary 删除SysRole
// @Tags sys-Role
// @Accept application/json
// @Product application/json
// @Param data body dto.TeamReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysRole} "{"code": 200, "data": [...]}"
// @Router /api/sys/role/del [post]
// @Security Bearer
func (e *SysRoleApi) Del(c *gin.Context) {
	var req dto.TeamReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if req.TeamId == 0 {
		e.Error(c, codes.ErrInvalidParameter("", "teamId is required"))
		return
	}

	admin := utils.GetAdminId(c)

	for _, id := range req.Ids {
		if err := service.SerSysRole.Del(id, admin, req.TeamId); err != nil {
			e.Error(c, err)
			return
		}
	}
	e.Ok(c)
}
