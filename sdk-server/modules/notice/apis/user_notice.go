package apis

import (
	"dilu/common/codes"
	"dilu/common/utils"
	"dilu/modules/notice/models"
	"dilu/modules/notice/service"
	"dilu/modules/notice/service/dto"
	"dilu/modules/sys/enums"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
)

type UserNoticeApi struct {
	base.BaseApi
}

var ApiUserNotice = UserNoticeApi{}

// QueryPage 获取用户通知列表
// @Summary 获取用户通知列表
// @Tags notice-UserNotice
// @Accept application/json
// @Product application/json
// @Param data body dto.UserNoticeGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.UserNotice}} "{"code": 200, "data": [...]}"
// @Router /api/notice/user-notice/page [post]
// @Security Bearer
func (e *UserNoticeApi) QueryPage(c *gin.Context) {
	var req dto.UserNoticeGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.UserNotice, 10)
	var total int64

	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}

	if err := service.SerUserNotice.QueryPage(req, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取用户通知
// @Summary 获取用户通知
// @Tags notice-UserNotice
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.UserNotice} "{"code": 200, "data": [...]}"
// @Router /api/notice/user-notice/get [post]
// @Security Bearer
func (e *UserNoticeApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.UserNotice
	if err := service.SerUserNotice.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建用户通知
// @Summary 创建用户通知
// @Tags notice-UserNotice
// @Accept application/json
// @Product application/json
// @Param data body dto.UserNoticeDto true "body"
// @Success 200 {object} base.Resp{data=string} "{"code": 200, "data": [...]}"
// @Router /api/notice/user-notice/create [post]
// @Security Bearer
func (e *UserNoticeApi) Create(c *gin.Context) {
	var req dto.UserNoticeDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	teamId := utils.GetTeamId(c)
	if teamId != -1 {
		req.TeamId = teamId
	}

	if err := service.SerUserNotice.Create(&req); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}

// Update 更新用户通知
// @Summary 更新用户通知
// @Tags notice-UserNotice
// @Accept application/json
// @Product application/json
// @Param data body dto.UserNoticeDto true "body"
// @Success 200 {object} base.Resp{data=string} "{"code": 200, "data": [...]}"
// @Router /api/notice/user-notice/update [post]
// @Security Bearer
func (e *UserNoticeApi) Update(c *gin.Context) {
	var req dto.UserNoticeDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	teamId := utils.GetTeamId(c)
	if teamId != -1 {
		if req.TeamId != teamId {
			e.Err(c, codes.Err403(nil))
			return
		}
	}
	if err := service.SerUserNotice.Update(&req); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}

// Del 删除用户通知
// @Summary 删除用户通知
// @Tags notice-UserNotice
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.UserNotice} "{"code": 200, "data": [...]}"
// @Router /api/notice/user-notice/del [post]
// @Security Bearer
func (e *UserNoticeApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	for _, id := range req.Ids {
		if err := service.SerUserNotice.Del(id); err != nil {
			e.Error(c, err)
			return
		}
	}

	e.Ok(c)
}

// GetUserNotice 获取用户通知列表
// @Summary 获取用户通知列表
// @Tags notice-UserNotice
// @Accept application/json
// @Product application/json
// @Param data body dto.UserNoticeGetPageReq true "body"
// @Success 200 {object} base.Resp{data=dto.NoticeDto} "{"code": 200, "data": [...]}"
// @Router /api/notice/user-notice/my [post]
// @Security Bearer
func (e *UserNoticeApi) GetUserNotice(c *gin.Context) {
	var req dto.UserNoticeGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	req.TeamId = enums.PlatformTeamId
	req.UserId = utils.GetAdminId(c)

	var res dto.NoticeDto

	if err := service.SerUserNotice.UserNotices(&req, &res); err != nil {
		e.Error(c, err)
		return
	}

	e.Ok(c, res)
}

// Read 读一条消息
// @Summary 获取用户通知
// @Tags notice-UserNotice
// @Accept application/json
// @Product application/json
// @Param data body dto.ReadNoticeDto true "body"
// @Success 200 {object} base.Resp{data=string} "{"code": 200, "data": [...]}"
// @Router /api/notice/user-notice/read [post]
// @Security Bearer
func (e *UserNoticeApi) Read(c *gin.Context) {
	var req dto.ReadNoticeDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	if req.Key == 1 {
		if err := service.SerUserNotice.ReadUserNotice(&req, e.GetReqId(c), enums.PlatformTeamId, utils.GetAdminId(c)); err != nil {
			e.Error(c, err)
			return
		}
	}

	e.Ok(c)
}
