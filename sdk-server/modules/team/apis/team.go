package apis

import (
	"dilu/common/utils"
	"dilu/modules/sys/enums"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type TeamApi struct {
	base.BaseApi
}

var ApiTeam = TeamApi{}

// Get 获取团队
// @Summary 获取团队
// @Tags Team
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysTeam} "{"code": 200, "data": [...]}"
// @Router /api/team/myinfo [post]
// @Security Bearer
func (e *TeamApi) Get(c *gin.Context) {
	var req base.ReqId
	// if err := c.ShouldBind(&req); err != nil {
	// 	e.Error(c, err)
	// 	return
	// }
	req.Id = utils.GetReqTeamId(c, req.Id)
	var data models.SysTeam
	if err := service.SerSysTeam.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新团队
// @Summary 更新团队
// @Tags Team
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysTeamDto true "body"
// @Success 200 {object} base.Resp{data=models.SysTeam} "{"code": 200, "data": [...]}"
// @Router /api/team/update [post]
// @Security Bearer
func (e *TeamApi) Update(c *gin.Context) {
	var req dto.SysTeamDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	req.Id = utils.GetReqTeamId(c, req.Id)
	var data models.SysTeam
	copier.Copy(&data, req)
	data.Status = 2
	if err := service.SerSysTeam.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// // Update 更新团队
// // @Summary 更新团队
// // @Tags Team
// // @Accept application/json
// // @Product application/json
// // @Param teamId header int false "团队id"
// // @Param data body dto.SysTeamDto true "body"
// // @Success 200 {object} base.Resp{data=models.SysTeam} "{"code": 200, "data": [...]}"
// // @Router /api/team/change [post]
// // @Security Bearer
// func (e *TeamApi) ChangeName(c *gin.Context) {
// 	var req dto.SysTeamDto
// 	if err := c.ShouldBind(&req); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	req.Id = utils.GetTeamId(c)
// 	var data models.SysTeam
// 	copier.Copy(&data, req)
// 	if err := service.SerSysTeam.UpdateById(&data); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	e.Ok(c, data)
// }

// GetUserMenus 获取用户菜单
// @Summary 获取用户菜单
// @Tags Team
// @Accept application/json
// @Product application/json
// @Success 200 {object} base.Resp{data=[]dto.MenuVo} "{"code": 200, "data": [...]}"
// @Router /api/team/user/menus [post]
// @Security Bearer
func (e *TeamApi) GetUserMenus(c *gin.Context) {
	teamId := utils.GetTeamId(c)
	uid := utils.GetAppUid(c)
	var ms []dto.MenuVo
	if err := service.SerSysMenu.GetUserMenus(teamId, uid, &ms); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, ms)
}

// GetGrantMenus 获取授权菜单
// @Summary 获取授权菜单
// @Tags sys-Menu
// @Accept application/json
// @Product application/json
// @Success 200 {object} base.Resp{data=[]models.SysMenu} "{"code": 200, "data": [...]}"
// @Router /api/team/menu/grant [post]
// @Security Bearer
func (e *TeamApi) GetGrantMenus(c *gin.Context) {
	var req dto.SysMenuReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	req.PlatformType = int(enums.MenuPriTeam)
	list := make([]models.SysMenu, 0)
	if err := service.SerSysMenu.GetMenus(&req, &list); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, list)
}
