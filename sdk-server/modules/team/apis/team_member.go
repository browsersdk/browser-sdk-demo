package apis

import (
	"dilu/common/codes"
	"dilu/common/utils"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"
	"fmt"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type TeamMemberApi struct {
	base.BaseApi
}

var ApiTeamMember = TeamMemberApi{}

// QueryPage 获取会员列表
// @Summary 获取会员列表
// @Tags team-Member
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysMemberGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysMember}} "{"code": 200, "data": [...]}"
// @Router /api/team/member/page [post]
// @Security Bearer
func (e *TeamMemberApi) QueryPage(c *gin.Context) {
	var req dto.SysMemberGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	req.TeamId = utils.GetReqTeamId(c, req.TeamId)
	list := make([]models.SysMember, 10)
	var total int64

	if err := service.SerSysMember.Query(req, &list, &total); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// // GetMembers 获取我的会员列表
// // @Summary 获取会员列表
// // @Tags team-Member
// // @Accept application/json
// // @Product application/json
// // @Param teamId header int false "团队id"
// // @Param data body dto.SysMemberGetPageReq true "body"
// // @Success 200 {object} base.Resp{data=[]models.SysMember} "{"code": 200, "data": [...]}"
// // @Router /api/team/member/members [post]
// // @Security Bearer
// func (e *TeamMemberApi) GetMembers(c *gin.Context) {
// 	var req dto.SysMemberGetReq
// 	if err := c.ShouldBind(&req); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	teamId := utils.GetTeamId(c)
// 	userId := utils.GetAppUid(c)
// 	var tu models.SysMember
// 	if err := service.SerSysMember.GetMember(teamId, userId, &tu); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	uid := 0
// 	if tu.PostId == enums.Staff.Id {
// 		uid = tu.UserId
// 	} else if tu.PostId > enums.Admin.Id {
// 		req.DeptPath = tu.DeptPath
// 	}
// 	req.TeamId = utils.GetReqTeamId(c, req.TeamId)
// 	list := make([]models.SysMember, 10)

// 	if err := service.SerSysMember.GetMembers(req.TeamId, uid, req.DeptPath, req.Name, req.Status, &list); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	e.Ok(c, list)
// }

// MyTeams 获取加入的团队
// @Summary 获取加入的团队
// @Tags team-Member
// @Accept application/json
// @Product application/json
// @Success 200 {object} base.Resp{data=[]dto.TeamMemberResp} "{"code": 200, "data": [...]}"
// @Router /api/team/member/myTeams [post]
// @Security Bearer
func (e *TeamMemberApi) MyTeams(c *gin.Context) {
	uid := utils.GetAppUid(c)
	fmt.Println("myTeam uid:", uid)
	if uid < 1 {
		e.Code(c, codes.InvalidToken_401)
		return
	}
	var list []dto.TeamMemberResp
	if err := service.SerSysMember.GetUserTeams(uid, &list); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, list)
}

// Exits 用户是否在团队中
// @Summary Exits 用户是否在团队中
// @Tags team-Member
// @Accept application/json
// @Product application/json
// @Param data body dto.ExistMemberReq true "body"
// @Param teamId header int false "团队id"
// @Success 200 {object} base.Resp{data=dto.ExistMemberResp} "{"code": 200, "data": [...]}"
// @Router /api/team/member/exist [post]
// @Security Bearer
func (e *TeamMemberApi) Exits(c *gin.Context) {
	var req dto.ExistMemberReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	req.TeamId = utils.GetReqTeamId(c, req.TeamId)
	b, err := service.SerSysMember.Exist(req.TeamId, req.UserId)
	if err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, dto.ExistMemberResp{
		Exist: b,
	})
}

// MyInfo 获取我的信息
// @Summary 获取我的信息
// @Tags team-Member
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Success 200 {object} base.Resp{data=[]dto.TeamMemberResp} "{"code": 200, "data": [...]}"
// @Router /api/team/member/myInfo [post]
// @Security Bearer
func (e *TeamMemberApi) MyInfo(c *gin.Context) {
	uid := utils.GetAppUid(c)
	if uid < 1 {
		e.Code(c, codes.InvalidToken_401)
		return
	}
	var data dto.TeamMemberResp
	teamId := utils.GetTeamId(c)
	// if teamId == 0 && utils.GetRoleId(c) != 0 {
	// 	data.TeamName = "Dilu后台管理"
	// 	data.UserId = uid
	// 	// data.Nickname = utils.GetNickname(c)
	// 	// data.Phone = utils.GetPhone(c)
	// } else {
	if err := service.SerSysMember.GetTeamMember(teamId, uid, &data); err != nil {
		e.Error(c, err)
		return
	}
	//}
	e.Ok(c, data)
}

// Get 获取会员
// @Summary 获取会员
// @Tags team-Member
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysMember} "{"code": 200, "data": [...]}"
// @Router /api/team/member/get [post]
// @Security Bearer
func (e *TeamMemberApi) Get(c *gin.Context) {
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
// @Tags team-Member
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysMemberDto true "body"
// @Success 200 {object} base.Resp{data=models.SysMember} "{"code": 200, "data": [...]}"
// @Router /api/team/member/create [post]
// @Security Bearer
func (e *TeamMemberApi) Create(c *gin.Context) {
	var req dto.SysMemberDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	req.TeamId = utils.GetReqTeamId(c, req.TeamId)

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
// @Tags team-Member
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysMemberDto true "body"
// @Success 200 {object} base.Resp{data=models.SysMember} "{"code": 200, "data": [...]}"
// @Router /api/team/member/update [post]
// @Security Bearer
func (e *TeamMemberApi) Update(c *gin.Context) {
	var req dto.SysMemberDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	req.TeamId = utils.GetReqTeamId(c, req.TeamId)
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
// @Tags team-Member
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysMember} "{"code": 200, "data": [...]}"
// @Router /api/team/member/del [post]
// @Security Bearer
func (e *TeamMemberApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	uid := utils.GetAppUid(c)
	teamId := utils.GetTeamId(c)
	for _, id := range req.Ids {
		if err := service.SerSysMember.Del(id, teamId, uid); err != nil {
			e.Error(c, err)
			return
		}
	}
	e.Ok(c)
}

// Update 更新会员
// @Summary 更新会员
// @Tags team-Member
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.ChangeMyMemberDto true "body"
// @Success 200 {object} base.Resp{data=models.SysMember} "{"code": 200, "data": [...]}"
// @Router /api/team/member/changeMyInfo [post]
// @Security Bearer
func (e *TeamMemberApi) ChangeMyInfo(c *gin.Context) {
	var req dto.ChangeMyMemberDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	teamId := utils.GetTeamId(c)
	userId := utils.GetAppUid(c)
	var data models.SysMember
	if err := service.SerSysMember.GetMember(teamId, userId, &data); err != nil {
		e.Error(c, err)
		return
	}
	data.Birthday = req.Birthday
	data.Gender = req.Gender
	data.Name = req.Name
	data.Nickname = req.Nickname
	data.Phone = req.Phone

	if err := service.SerSysMember.Update(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}
