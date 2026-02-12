package service

import (
	"dilu/modules/sys/enums"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service/dto"
	"fmt"

	"github.com/baowk/dilu-core/core/base"
	"github.com/jinzhu/copier"
)

type AdminDeptService struct {
	*base.BaseService
}

var SerAdminDept = AdminDeptService{
	base.NewService("sys"),
}

func (s *AdminDeptService) List(list *[]models.AdminDept) error {
	return s.DB().Where("status = 1").Find(list).Error
}

func (s *AdminDeptService) Page(req *dto.AdminDeptGetPageReq, list *[]models.AdminDept, total *int64) error {
	db := s.DB().Offset(req.GetOffset()).Limit(req.GetSize())
	if req.Status != 0 {
		db.Where("status = ?", req.Status)
	}
	return db.Find(list).Offset(-1).Limit(-1).Count(total).Error
}

func (s *AdminDeptService) GetDepts(teamId int, list *[]models.AdminDept) error {
	if teamId != 0 {
		return s.DB().Where("team_id = ?", teamId).Find(list).Error
	}
	return s.DB().Find(list).Error
}

func (s *AdminDeptService) Create(req dto.AdminDeptDto, adminId int) (models.AdminDept, error) {
	var data models.AdminDept
	var pdept models.AdminDept
	if req.ParentId > 0 {
		if err := s.Get(req.ParentId, &pdept); err != nil {
			return data, err
		}
	}

	copier.Copy(&data, req)
	data.CreateBy = adminId
	data.Status = enums.STATUS_DEF
	if err := s.BaseService.Create(&data); err != nil {
		return data, err
	}
	if req.ParentId > 0 {
		data.DeptPath = fmt.Sprintf("%s%d/", pdept.DeptPath, data.Id)
	} else {
		data.DeptPath = fmt.Sprintf("/0/%d/", data.Id)
	}
	err := s.UpdateById(data)
	return data, err
}

func (s *AdminDeptService) Update(req dto.AdminDeptDto, adminId int) (models.AdminDept, error) {
	var old models.AdminDept
	if err := s.Get(req.Id, &old); err != nil {
		return old, err
	}
	var pdept models.AdminDept
	if req.ParentId > 0 && old.ParentId != req.ParentId {
		if err := s.Get(req.ParentId, &pdept); err != nil {
			return old, err
		}
	}
	copier.Copy(&old, req)
	old.UpdateBy = adminId

	if pdept.Id > 0 {
		old.DeptPath = fmt.Sprintf("%s%d/", pdept.DeptPath, old.Id)
	}
	err := s.UpdateById(old)
	return old, err
}

func (s *AdminDeptService) Del(id, teamId int, uid int) error {
	var dept models.AdminDept
	if err := s.Get(id, &dept); err != nil {
		return err
	}

	if err := s.DB().Where("parent_id = ?", id).Find(&dept).Error; err == nil {
		return fmt.Errorf("存在子部门，不允许删除")
	}
	cnt, err := SerSysMember.CountMembers(dept.Id)
	if err != nil {
		return err
	}
	if cnt > 0 {
		return fmt.Errorf("存在成员，不允许删除")
	}

	dept.Status = enums.STATUS_DEL
	dept.UpdateBy = uid

	return s.UpdateById(dept)
}
