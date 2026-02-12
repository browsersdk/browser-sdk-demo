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

type TeamCfgApi struct {
	base.BaseApi
}

var ApiTeamCfg = TeamCfgApi{}

// QueryPage 获取配置列表
// @Summary 获取配置列表
// @Tags team-Cfg
// @Accept application/json
// @Product application/json
// @Param data body dto.SysCfgGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysCfg}} "{"code": 200, "data": [...]}"
// @Router /api/team/cfg/page [post]
// @Security Bearer
func (e *TeamCfgApi) QueryPage(c *gin.Context) {
	var req dto.SysCfgGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysCfg, 0, req.GetSize())
	var total int64

	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}

	req.TeamId = utils.GetTeamId(c)

	if err := service.SerSysCfg.QueryPage(req, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取配置
// @Summary 获取配置
// @Tags team-Cfg
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysCfg} "{"code": 200, "data": [...]}"
// @Router /api/team/cfg/get [post]
// @Security Bearer
func (e *TeamCfgApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysCfg
	if err := service.SerSysCfg.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建配置
// @Summary 创建配置
// @Tags team-Cfg
// @Accept application/json
// @Product application/json
// @Param data body dto.SysCfgDto true "body"
// @Success 200 {object} base.Resp{data=models.SysCfg} "{"code": 200, "data": [...]}"
// @Router /api/team/cfg/create [post]
// @Security Bearer
func (e *TeamCfgApi) Create(c *gin.Context) {
	var req dto.SysCfgDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	req.TeamId = utils.GetTeamId(c)
	var data models.SysCfg
	copier.Copy(&data, req)
	if err := service.SerSysCfg.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新配置
// @Summary 更新配置
// @Tags team-Cfg
// @Accept application/json
// @Product application/json
// @Param data body dto.SysCfgDto true "body"
// @Success 200 {object} base.Resp{data=models.SysCfg} "{"code": 200, "data": [...]}"
// @Router /api/team/cfg/update [post]
// @Security Bearer
func (e *TeamCfgApi) Update(c *gin.Context) {
	var req dto.SysCfgDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysCfg
	copier.Copy(&data, req)
	if err := service.SerSysCfg.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除配置
// @Summary 删除配置
// @Tags team-Cfg
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysCfg} "{"code": 200, "data": [...]}"
// @Router /api/team/cfg/del [post]
// @Security Bearer
func (e *TeamCfgApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysCfg.DelIds(&models.SysCfg{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
