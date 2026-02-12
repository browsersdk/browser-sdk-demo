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

type SysMemberApi struct {
	base.BaseApi
}

var ApiSysMember = SysMemberApi{}

// QueryPage 获取会员列表
// @Summary 获取会员列表
// @Tags sys-Member
// @Accept application/json
// @Product application/json
// @Param data body dto.SysMemberGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysMember}} "{"code": 200, "data": [...]}"
// @Router /api/sys/member/page [post]
// @Security Bearer
func (e *SysMemberApi) QueryPage(c *gin.Context) {
	var req dto.SysMemberGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	if req.TeamId == 0 {
		e.Error(c, codes.ErrInvalidParameter("", "teamId is required"))
		return
	}
	list := make([]models.SysMember, 10)
	var total int64

	if err := service.SerSysMember.Query(req, &list, &total); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// GetMembers 获取我的会员列表
// @Summary 获取会员列表
// @Tags sys-Member
// @Accept application/json
// @Product application/json
// @Param data body dto.SysMemberGetPageReq true "body"
// @Success 200 {object} base.Resp{data=[]models.SysMember} "{"code": 200, "data": [...]}"
// @Router /api/sys/member/members [post]
// @Security Bearer
func (e *SysMemberApi) GetMembers(c *gin.Context) {
	var req dto.SysMemberGetReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	// teamId := utils.GetTeamId(c)
	// userId := utils.GetUserId(c)
	// var tu models.SysMember
	// if err := service.SerSysMember.GetMember(teamId, userId, &tu); err != nil {
	// 	e.Error(c, err)
	// 	return
	// }
	// uid := 0
	// if tu.PostId == enums.Staff.Id {
	// 	uid = tu.UserId
	// } else if tu.PostId > enums.Admin.Id {
	// 	req.DeptPath = tu.DeptPath
	// }
	if req.TeamId == 0 {
		e.Error(c, codes.ErrInvalidParameter("", "teamId is required"))
		return
	}
	list := make([]models.SysMember, 10)

	if err := service.SerSysMember.GetMembers(req.TeamId, 0, req.DeptPath, req.Name, req.Status, &list); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, list)
}

// // MyTeams 获取加入的团队
// // @Summary 获取加入的团队
// // @Tags sys-Member
// // @Accept application/json
// // @Product application/json
// // @Success 200 {object} base.Resp{data=[]dto.TeamMemberResp} "{"code": 200, "data": [...]}"
// // @Router /api/sys/member/myTeams [post]
// // @Security Bearer
// func (e *SysMemberApi) MyTeams(c *gin.Context) {
// 	uid := utils.GetUserId(c)
// 	fmt.Println("myTeam", uid)
// 	if uid < 1 {
// 		e.Code(c, codes.InvalidToken_401)
// 		return
// 	}
// 	var list []dto.TeamMemberResp
// 	if err := service.SerSysMember.GetUserTeams(uid, &list); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	e.Ok(c, list)
// }

// Exits 用户是否在团队中
// @Summary Exits 用户是否在团队中
// @Tags sys-Member
// @Accept application/json
// @Product application/json
// @Param data body dto.ExistMemberReq true "body"
// @Success 200 {object} base.Resp{data=dto.ExistMemberResp} "{"code": 200, "data": [...]}"
// @Router /api/sys/member/exist [post]
// @Security Bearer
func (e *SysMemberApi) Exits(c *gin.Context) {
	var req dto.ExistMemberReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if req.TeamId == 0 {
		e.Error(c, codes.ErrInvalidParameter("", "teamId is required"))
		return
	}
	b, err := service.SerSysMember.Exist(req.TeamId, req.UserId)
	if err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, dto.ExistMemberResp{
		Exist: b,
	})
}

// // MyInfo 获取我的信息
// // @Summary 获取我的信息
// // @Tags sys-Member
// // @Accept application/json
// // @Product application/json
// // @Success 200 {object} base.Resp{data=[]dto.TeamMemberResp} "{"code": 200, "data": [...]}"
// // @Router /api/sys/member/myInfo [post]
// // @Security Bearer
// func (e *SysMemberApi) MyInfo(c *gin.Context) {
// 	uid := utils.GetUserId(c)
// 	if uid < 1 {
// 		e.Code(c, codes.InvalidToken_401)
// 		return
// 	}
// 	var data dto.TeamMemberResp
// 	teamId := utils.GetTeamId(c)
// 	if teamId == 0 && utils.GetRoleId(c) != 0 {
// 		data.TeamName = "Dilu后台管理"
// 		data.UserId = uid
// 		// data.Nickname = utils.GetNickname(c)
// 		// data.Phone = utils.GetPhone(c)
// 	} else {
// 		if err := service.SerSysMember.GetTeamMember(teamId, uid, &data); err != nil {
// 			e.Error(c, err)
// 			return
// 		}
// 	}
// 	e.Ok(c, data)
// }

// Get 获取会员
// @Summary 获取会员
// @Tags sys-Member
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysMember} "{"code": 200, "data": [...]}"
// @Router /api/sys/member/get [post]
// @Security Bearer
func (e *SysMemberApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysMember
	if err := service.SerSysMember.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建会员
// @Summary 创建会员
// @Tags sys-Member
// @Accept application/json
// @Product application/json
// @Param data body dto.SysMemberDto true "body"
// @Success 200 {object} base.Resp{data=models.SysMember} "{"code": 200, "data": [...]}"
// @Router /api/sys/member/create [post]
// @Security Bearer
func (e *SysMemberApi) Create(c *gin.Context) {
	var req dto.SysMemberDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	if req.TeamId == 0 {
		e.Error(c, codes.ErrInvalidParameter("", "teamId is required"))
		return
	}

	var data models.SysMember
	copier.Copy(&data, req)
	if err := service.SerSysMember.Create(&data, nil); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新会员
// @Summary 更新会员
// @Tags sys-Member
// @Accept application/json
// @Product application/json
// @Param data body dto.SysMemberDto true "body"
// @Success 200 {object} base.Resp{data=models.SysMember} "{"code": 200, "data": [...]}"
// @Router /api/sys/member/update [post]
// @Security Bearer
func (e *SysMemberApi) Update(c *gin.Context) {
	var req dto.SysMemberDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysMember
	copier.Copy(&data, req)
	if err := service.SerSysMember.Update(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除会员
// @Summary 删除会员
// @Tags sys-Member
// @Accept application/json
// @Product application/json
// @Param data body dto.TeamReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysMember} "{"code": 200, "data": [...]}"
// @Router /api/sys/member/del [post]
// @Security Bearer
func (e *SysMemberApi) Del(c *gin.Context) {
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
		if err := service.SerSysMember.Del(id, req.TeamId, admin); err != nil {
			e.Error(c, err)
			return
		}
	}
	e.Ok(c)
}

// // Update 更新会员
// // @Summary 更新会员
// // @Tags sys-Member
// // @Accept application/json
// // @Product application/json
// // @Param data body dto.ChangeMyMemberDto true "body"
// // @Success 200 {object} base.Resp{data=models.SysMember} "{"code": 200, "data": [...]}"
// // @Router /api/sys/member/changeMyInfo [post]
// // @Security Bearer
// func (e *SysMemberApi) ChangeMyInfo(c *gin.Context) {
// 	var req dto.ChangeMyMemberDto
// 	if err := c.ShouldBind(&req); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	teamId := utils.GetTeamId(c)
// 	userId := utils.GetUserId(c)
// 	var data models.SysMember
// 	if err := service.SerSysMember.GetMember(teamId, userId, &data); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	data.Birthday = req.Birthday
// 	data.Gender = req.Gender
// 	data.Name = req.Name
// 	data.Nickname = req.Nickname
// 	data.Phone = req.Phone

// 	if err := service.SerSysMember.Update(&data); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	e.Ok(c, data)
// }
