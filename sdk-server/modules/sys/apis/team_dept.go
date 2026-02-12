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

type SysDeptApi struct {
	base.BaseApi
}

var ApiSysDept = SysDeptApi{}

// QueryPage 获取SysDept列表
// @Summary Page接口
// @Tags sys-Dept
// @Accept application/json
// @Product application/json
// @Param data body dto.SysDeptGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysDept}} "{"code": 200, "data": [...]}"
// @Router /api/sys/dept/page [post]
// @Security Bearer
func (e *SysDeptApi) QueryPage(c *gin.Context) {
	var req dto.SysDeptGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if req.TeamId == 0 {
		e.Error(c, codes.ErrInvalidParameter("", "teamId is required"))
		return
	}
	list := make([]models.SysDept, 10)
	var total int64
	//req.TeamId = utils.GetReqTeamId(c, req.TeamId)
	if err := service.SerSysDept.Page(&req, &list, &total); err != nil {
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
// @Param data body dto.SysDeptGetPageReq true "body"
// @Success 200 {object} base.Resp{data=[]models.SysDept} "{"code": 200, "data": [...]}"
// @Router /api/sys/dept/all [post]
// @Security Bearer
func (e *SysDeptApi) List(c *gin.Context) {
	var req dto.SysDeptGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if req.TeamId == 0 {
		e.Error(c, codes.ErrInvalidParameter("", "teamId is required"))
		return
	}
	list := make([]models.SysDept, 10)

	if err := service.SerSysDept.GetDepts(req.TeamId, &list); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, list)
}

// Get 获取SysDept
// @Summary 获取SysDept
// @Tags sys-Dept
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysDept} "{"code": 200, "data": [...]}"
// @Router /api/sys/dept/get [post]
// @Security Bearer
func (e *SysDeptApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysDept
	if err := service.SerSysDept.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建SysDept
// @Summary 创建SysDept
// @Tags sys-Dept
// @Accept application/json
// @Product application/json
// @Param data body dto.SysDeptDto true "body"
// @Success 200 {object} base.Resp{data=models.SysDept} "{"code": 200, "data": [...]}"
// @Router /api/sys/dept/create [post]
// @Security Bearer
func (e *SysDeptApi) Create(c *gin.Context) {
	var req dto.SysDeptDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysDept
	if req.TeamId == 0 {
		e.Error(c, codes.ErrInvalidParameter("", "teamId is required"))
		return
	}
	adminId := utils.GetAdminId(c)
	if err := service.SerSysDept.CreateDept(req, adminId, e.GetReqId(c)); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新SysDept
// @Summary 更新SysDept
// @Tags sys-Dept
// @Accept application/json
// @Product application/json
// @Param data body dto.SysDeptDto true "body"
// @Success 200 {object} base.Resp{data=models.SysDept} "{"code": 200, "data": [...]}"
// @Router /api/sys/dept/update [post]
// @Security Bearer
func (e *SysDeptApi) Update(c *gin.Context) {
	var req dto.SysDeptDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	var data models.SysDept

	adminId := utils.GetAdminId(c)
	if err := service.SerSysDept.UpdateDept(req, adminId, e.GetReqId(c)); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除SysDept
// @Summary 删除SysDept
// @Tags sys-Dept
// @Accept application/json
// @Product application/json
// @Param data body dto.TeamReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysDept} "{"code": 200, "data": [...]}"
// @Router /api/sys/dept/del [post]
// @Security Bearer
func (e *SysDeptApi) Del(c *gin.Context) {
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
		if err := service.SerSysDept.Del(id, req.TeamId, admin); err != nil {
			e.Error(c, err)
			return
		}
	}
	e.Ok(c)
}
