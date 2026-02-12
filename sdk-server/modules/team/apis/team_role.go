package apis

import (
	"dilu/common/utils"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
)

type TeamRoleApi struct {
	base.BaseApi
}

var ApiTeamRole = TeamRoleApi{}

// QueryPage 获取SysRole列表
// @Summary Page接口
// @Tags team-Role
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysRoleGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysRole}} "{"code": 200, "data": [...]}"
// @Router /api/team/role/page [post]
// @Security Bearer
func (e *TeamRoleApi) QueryPage(c *gin.Context) {
	var req dto.SysRoleGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysRole, 0)
	var total int64
	req.TeamId = utils.GetReqTeamId(c, req.TeamId)

	if err := service.SerSysRole.Page(&req, &list, &total); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// List 获取角色列表
// @Summary 获取角色列表
// @Tags team-Role
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysRoleGetPageReq true "body"
// @Success 200 {object} base.Resp{data=[]models.SysRole} "{"code": 200, "data": [...]}"
// @Router /api/team/role/list [post]
// @Security Bearer
func (e *TeamRoleApi) List(c *gin.Context) {
	var req dto.SysRoleGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysRole, 0)

	if err := service.SerSysRole.Query(utils.GetReqTeamId(c, req.TeamId), req.Status, &list); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, list)
}

// Get 获取SysRole
// @Summary 获取SysRole
// @Tags team-Role
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=dto.SysRoleDtoResp} "{"code": 200, "data": [...]}"
// @Router /api/team/role/get [post]
// @Security Bearer
func (e *TeamRoleApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data dto.SysRoleDtoResp
	if err := service.SerSysRole.GetRole(req.Id, utils.GetAppUid(c), utils.GetTeamId(c), &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建SysRole
// @Summary 创建SysRole
// @Tags team-Role
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysRoleDto true "body"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/team/role/create [post]
// @Security Bearer
func (e *TeamRoleApi) Create(c *gin.Context) {
	var req dto.SysRoleDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysRole.Create(utils.GetAppUid(c), utils.GetTeamId(c), req); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}

// Update 更新SysRole
// @Summary 更新SysRole
// @Tags team-Role
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysRoleDto true "body"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/team/role/update [post]
// @Security Bearer
func (e *TeamRoleApi) Update(c *gin.Context) {
	var req dto.SysRoleDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysRole.Update(utils.GetAppUid(c), utils.GetTeamId(c), req); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}

// Del 删除SysRole
// @Summary 删除SysRole
// @Tags team-Role
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysRole} "{"code": 200, "data": [...]}"
// @Router /api/team/role/del [post]
// @Security Bearer
func (e *TeamRoleApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	for _, id := range req.Ids {
		if err := service.SerSysRole.Del(id, utils.GetAppUid(c), utils.GetTeamId(c)); err != nil {
			e.Error(c, err)
			return
		}
	}
	e.Ok(c)
}
