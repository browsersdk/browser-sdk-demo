package service

import (
	//"github.com/acmestack/gorm-plus/gplus"

	"dilu/common/ckey"
	"dilu/modules/sys/enums"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service/dto"
	"encoding/json"
	"fmt"
	"time"

	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/jinzhu/copier"
)

type SysRoleService struct {
	*base.BaseService
}

var SerSysRole = SysRoleService{
	base.NewService(consts.DB_DEF),
}

func (s *SysRoleService) Create(userId, teamId int, req dto.SysRoleDto) error {
	var model models.SysRole
	copier.Copy(&model, req)
	model.TeamId = teamId
	model.CreateBy = userId
	if model.Status == 0 {
		model.Status = enums.STATUS_DEF
	}
	tx := s.DB().Begin()
	if err := tx.Create(&model).Error; err != nil {
		tx.Rollback()
		return err
	}
	//m := make(map[int]bool, 0)
	var rms []models.SysRoleMenu
	for _, mids := range req.MenuIds {
		rms = append(rms, models.SysRoleMenu{RoleId: model.Id, MenuId: mids})
	}
	if err := tx.Create(rms).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

}

func (s *SysRoleService) Update(userId, teamId int, req dto.SysRoleDto) error {
	var model models.SysRole
	copier.Copy(&model, req)
	model.TeamId = teamId
	model.UpdateBy = userId

	tx := s.DB().Begin()
	if err := tx.Updates(&model).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("role_id = ?", model.Id).Delete(&models.SysRoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	//m := make(map[int]bool, 0)
	var rms []models.SysRoleMenu
	for _, mids := range req.MenuIds {
		rms = append(rms, models.SysRoleMenu{RoleId: model.Id, MenuId: mids})
	}
	if err := tx.Create(rms).Error; err != nil {
		tx.Rollback()
		return err
	}
	err := tx.Commit().Error
	if err == nil {
		core.Cache.Del(ckey.RoleUserMenu(int(enums.MenuPriTeam), model.Id))
		core.Cache.Del(ckey.RoleApis(int(enums.MenuPriTeam), model.Id))
	}
	return err
}

func (s *SysRoleService) GetRole(id, userId, teamId int, data *dto.SysRoleDtoResp) error {
	var model models.SysRole
	if err := s.DB().First(&model, id).Error; err != nil {
		return err
	}
	copier.Copy(data, model)
	var menuIds []int
	if err := s.DB().Model(&models.SysRoleMenu{}).Select("menu_id").Where("role_id = ?", id).Find(&menuIds).Error; err != nil {
		return err
	}
	data.MenuIds = menuIds
	return nil
}

func (s *SysRoleService) Query(teamId, status int, list *[]models.SysRole) error {
	db := s.DB()
	if teamId != 0 {
		db = db.Where("team_id = ?", teamId)
	}
	if status != 0 {
		db = db.Where("status = ?", status)
	}
	return db.Find(list).Error
}

func (s *SysRoleService) Page(req *dto.SysRoleGetPageReq, list *[]models.SysRole, total *int64) error {
	db := s.DB().Offset(req.GetOffset()).Limit(req.GetSize())
	if req.TeamId != 0 {
		db.Where("team_id = ?", req.TeamId)
	}
	if req.Status != 0 {
		db.Where("status = ?", req.Status)
	}
	return db.Find(list).Offset(-1).Limit(-1).Count(total).Error
}

func (s *SysRoleService) Del(id, userId, teamId int) error {
	var model models.SysRole
	if err := s.Get(id, &model); err != nil {
		return err
	}
	if model.TeamId != teamId {
		return fmt.Errorf("团队id不匹配")
	}
	model.UpdateBy = userId
	model.Status = enums.STATUS_DEL
	err := s.DB().Save(&model).Error
	if err == nil {
		core.Cache.Del(ckey.RoleUserMenu(int(enums.MenuPriTeam), model.Id))
		core.Cache.Del(ckey.RoleApis(int(enums.MenuPriTeam), model.Id))
	}
	return err
}

func (s *SysRoleService) GetApis(roleId int, ids *[]int) error {
	data, _ := core.Cache.Get(ckey.RoleApis(int(enums.MenuPriTeam), roleId))
	if data != "" {
		err := json.Unmarshal([]byte(data), ids)
		if err == nil {
			return nil
		}
	}
	if err := s.DB().Raw("select sys_api_id from sys_menu_api_rule r,team_role_menu rm  where  rm.role_id = ? and rm.menu_id = r.sys_menu_id", roleId).
		Find(&ids).Error; err != nil {
		return err
	}
	core.Cache.Set(ckey.RoleApis(int(enums.MenuPriTeam), roleId), ids, time.Hour)
	return nil
}
