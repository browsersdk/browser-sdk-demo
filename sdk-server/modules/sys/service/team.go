package service

import (
	"dilu/common/utils"
	"dilu/modules/sys/enums"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service/dto"
	"time"

	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core/base"
	"github.com/pkg/errors"
)

type SysTeamService struct {
	*base.BaseService
}

var SerSysTeam = SysTeamService{
	base.NewService(consts.DB_DEF),
}

func (s *SysTeamService) Page(req *dto.SysTeamGetPageReq, data *[]dto.SysTeamData, total *int64) error {
	db := s.DB()
	if req.Name != "" {
		db = db.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Status != 0 {
		db = db.Where("status = ?", req.Status)
	}
	var list []models.SysTeam
	err := db.Limit(req.GetSize()).Offset(req.GetOffset()).Find(&list).Limit(-1).Offset(-1).Count(total).Error
	if err != nil {
		return errors.Wrap(err, "获取团队列表失败")
	}
	for _, v := range list {
		*data = append(*data, dto.SysTeamData{
			Id:        v.Id,
			Name:      v.Name,
			Owner:     v.Owner,
			Status:    v.Status,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			AppSk:     v.GetAppKey(),
		})
	}
	return nil
}

func (s *SysTeamService) Create(team *models.SysTeam, user *models.SysUser, owner *models.SysMember) error {
	if team.Name == "" {
		return errors.New("团队名不能为空")
	}
	if team.Owner == 0 {
		return errors.New("用户id不能为空")
	}
	if user == nil {
		if err := SerSysUser.Get(team.Owner, user); err != nil {
			return err
		}
	}
	if user.Status != 1 {
		return errors.New("用户状态异常")
	}
	team.Status = enums.STATUS_OK
	team.AppKeyCreatedAt = time.Now()
	if err := s.DB().Create(team).Error; err != nil {
		return err
	}
	if owner != nil && owner.Id > 0 {
		return nil
	}
	owner = &models.SysMember{
		TeamId:   team.Id,
		UserId:   user.Id,
		Nickname: user.Nickname,
		Name:     user.Name,
		PostId:   enums.Admin.Id,
		Status:   1,
		RoleId:   enums.RoleTeamOwner,
	}

	if err := SerSysMember.Create(owner, user); err != nil {
		return err
	}

	return nil
}

func (s *SysTeamService) GetTeamByOwner(owner int, team *models.SysTeam) error {
	return s.DB().Where("owner = ?", owner).Find(team).Error
}

func (s *SysTeamService) DelIds(ids []int, admin int) error {
	if len(ids) == 0 {
		return nil
	}
	u := map[string]any{
		"status":     enums.STATUS_DEL,
		"updated_at": time.Now(),
		"update_by":  0,
	}
	return s.DB().Model(&models.SysTeam{}).Where("id in (?)", ids).Updates(u).Error
}

func (s *SysTeamService) ValidAppKey(appKey string) (int, error) {
	tid, exp, err := utils.DecodeTeamKey(appKey)
	if err != nil {
		return 0, err
	}
	var team models.SysTeam
	if err := s.Get(tid, &team); err != nil {
		return 0, err
	}
	if team.Status < enums.STATUS_DEF {
		return 0, errors.New("team is lock")
	}
	if exp < team.AppKeyCreatedAt.Unix() {
		return 0, errors.New("team appKey is expired")
	}
	return tid, nil
}
