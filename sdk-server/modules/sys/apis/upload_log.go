package apis

import (
	"dilu/common/consts"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type UploadLogApi struct {
	base.BaseApi
}

var ApiUploadLog = UploadLogApi{}

// 上传文件
// @Summary 上传文件
// @Tags sys-Cfg
// @Accept multipart/form-data
// @Product application/json
// @Param file formData file true "文件"
// @Success 200 {object} base.Resp{data=string} "{"code": 200, "data": [...]}"
// @Router /api/sys/cfg/upload [post]
// @Security Bearer
func (e *UploadLogApi) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		e.Error(c, err)
		return
	}
	url, err := service.SerUploadLog.UploadFile(file, consts.UPLOAD_MNG)
	if err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, url)
}

// QueryPage 获取UploadLog列表
// @Summary 获取UploadLog列表
// @Tags sys-UploadLog
// @Accept application/json
// @Product application/json
// @Param data body dto.UploadLogGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.UploadLog}} "{"code": 200, "data": [...]}"
// @Router /api/sys/upload-log/page [post]
// @Security Bearer
func (e *UploadLogApi) QueryPage(c *gin.Context) {
	var req dto.UploadLogGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.UploadLog, 0, req.GetSize())
	var total int64

	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}

	if err := service.SerUploadLog.QueryPage(req, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取UploadLog
// @Summary 获取UploadLog
// @Tags sys-UploadLog
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.UploadLog} "{"code": 200, "data": [...]}"
// @Router /api/sys/upload-log/get [post]
// @Security Bearer
func (e *UploadLogApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.UploadLog
	if err := service.SerUploadLog.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建UploadLog
// @Summary 创建UploadLog
// @Tags sys-UploadLog
// @Accept application/json
// @Product application/json
// @Param data body dto.UploadLogDto true "body"
// @Success 200 {object} base.Resp{data=models.UploadLog} "{"code": 200, "data": [...]}"
// @Router /api/sys/upload-log/create [post]
// @Security Bearer
func (e *UploadLogApi) Create(c *gin.Context) {
	var req dto.UploadLogDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.UploadLog
	copier.Copy(&data, req)
	if err := service.SerUploadLog.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新UploadLog
// @Summary 更新UploadLog
// @Tags sys-UploadLog
// @Accept application/json
// @Product application/json
// @Param data body dto.UploadLogDto true "body"
// @Success 200 {object} base.Resp{data=models.UploadLog} "{"code": 200, "data": [...]}"
// @Router /api/sys/upload-log/update [post]
// @Security Bearer
func (e *UploadLogApi) Update(c *gin.Context) {
	var req dto.UploadLogDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.UploadLog
	copier.Copy(&data, req)
	if err := service.SerUploadLog.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除UploadLog
// @Summary 删除UploadLog
// @Tags sys-UploadLog
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.UploadLog} "{"code": 200, "data": [...]}"
// @Router /api/sys/upload-log/del [post]
// @Security Bearer
func (e *UploadLogApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerUploadLog.DelIds(&models.UploadLog{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
