package service

import (
	"dilu/common/ckey"
	"dilu/modules/sys/enums"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/jinzhu/copier"
)

type AdminRoleService struct {
	*base.BaseService
}

var SerAdminRole = AdminRoleService{
	base.NewService("sys"),
}

func (s *AdminRoleService) Query(status int, list *[]models.AdminRole) error {
	db := s.DB()
	if status != 0 {
		db = db.Where("status = ?", status)
	}
	return db.Find(list).Error
}

func (s *AdminRoleService) Create(userId int, req *dto.AdminRoleDto) error {
	var model models.AdminRole
	copier.Copy(&model, req)
	model.CreateBy = userId
	model.Status = enums.STATUS_DEF
	tx := s.DB().Begin()
	if err := tx.Create(&model).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

}

func (s *AdminRoleService) Update(userId int, req *dto.AdminRoleDto) error {
	var model models.AdminRole
	copier.Copy(&model, req)
	model.UpdateBy = userId

	tx := s.DB().Begin()
	if err := tx.Updates(&model).Error; err != nil {
		tx.Rollback()
		return err
	}
	core.Cache.Del(ckey.RoleUserMenu(int(enums.MenuPriPlatform), int(req.Id)))
	return tx.Commit().Error
}

func (s *AdminRoleService) Del(id, userId int) error {
	var model models.SysRole
	if err := s.Get(id, &model); err != nil {
		return err
	}

	model.UpdateBy = userId
	model.Status = enums.STATUS_DEL
	err := s.DB().Save(&model).Error
	if err == nil {
		core.Cache.Del(ckey.RoleUserMenu(int(enums.MenuPriPlatform), model.Id))
	}
	return err
}
