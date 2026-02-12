package apis

import (
	"dilu/common/codes"
	"dilu/common/utils"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type BrowserApi struct {
	base.BaseApi
}

var ApiBrowser = BrowserApi{}

// QueryPage 获取Browser列表
// @Summary Page接口
// @Tags sys-Browser
// @Accept application/json
// @Product application/json
// @Param data body dto.BrowserGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.Browser}} "{"code": 200, "data": [...]}"
// @Router /api/app/browser/page [post]
// @Security Bearer
func (e *BrowserApi) QueryPage(c *gin.Context) {
	uid := utils.GetAppUid(c)
	if uid == 0 {
		e.Code(c, codes.InvalidToken_401)
		return
	}
	var req dto.BrowserGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.Browser, 10)
	var total int64

	var model models.Browser
	if err := copier.Copy(&model, req); err != nil {
		e.Error(c, err)
		return
	}
	model.UserId = uid

	if err := service.SerBrowser.Page(model, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取Browser
// @Summary 获取Browser
// @Tags sys-Browser
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.Browser} "{"code": 200, "data": [...]}"
// @Router /api/app/browser/get [post]
// @Security Bearer
func (e *BrowserApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Browser
	if err := service.SerBrowser.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建Browser
// @Summary 创建Browser
// @Tags sys-Browser
// @Accept application/json
// @Product application/json
// @Param data body dto.BrowserDto true "body"
// @Success 200 {object} base.Resp{data=models.Browser} "{"code": 200, "data": [...]}"
// @Router /api/app/browser/create [post]
// @Security Bearer
func (e *BrowserApi) Create(c *gin.Context) {
	var req dto.BrowserDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Browser
	copier.Copy(&data, req)

	data.UserId = utils.GetAppUid(c)
	data.EnvId = utils.GenUint()
	data.Status = 1
	if err := service.SerBrowser.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新Browser
// @Summary 更新Browser
// @Tags sys-Browser
// @Accept application/json
// @Product application/json
// @Param data body dto.BrowserDto true "body"
// @Success 200 {object} base.Resp{data=models.Browser} "{"code": 200, "data": [...]}"
// @Router /api/app/browser/update [post]
// @Security Bearer
func (e *BrowserApi) Update(c *gin.Context) {
	var req dto.BrowserDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Browser
	copier.Copy(&data, req)
	if err := service.SerBrowser.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除Browser
// @Summary 删除Browser
// @Tags sys-Browser
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.Browser} "{"code": 200, "data": [...]}"
// @Router /api/app/browser/del [post]
// @Security Bearer
func (e *BrowserApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerBrowser.DelIds(&models.Browser{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
