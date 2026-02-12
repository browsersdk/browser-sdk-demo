package service

import (
	"dilu/common/ckey"
	"dilu/modules/notice/models"
	"dilu/modules/notice/service/dto"
	"encoding/json"
	"fmt"
	"time"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/jinzhu/copier"
)

type UserNoticeService struct {
	*base.BaseService
}

var SerUserNotice = UserNoticeService{
	base.NewService("notice"),
}

func (s *UserNoticeService) UserNotices(req *dto.UserNoticeGetPageReq, res *dto.NoticeDto) error {
	pns := SerPubNotice.GetNotice(req.TeamId)
	change := false
	if len(pns) > 0 {
		var nu models.UserNotice
		s.DB().Where("team_id = ?", req.TeamId).Where("user_id = ?", req.UserId).Where("pub_id > 0").
			Order("pub_id desc").Limit(1).Find(&nu)

		if nu.Id > 0 {
			var is []models.UserNotice
			for _, v := range pns {
				if v.Id <= nu.PubId {
					break
				}
				var nuT models.UserNotice
				copier.Copy(&nuT, v)
				nuT.Id = 0
				nuT.PubId = v.Id
				nuT.Status = 1
				nuT.TeamId = req.TeamId
				nuT.UserId = req.UserId
				is = append(is, nuT)
			}
			if len(is) > 0 {
				s.DB().Create(is)
				change = true
			}
		}
	}
	if req.GetOffset() == 0 && !change {
		data, _ := core.Cache.Get(ckey.MemberNotice(req.TeamId, req.UserId))
		if data != "" {
			err := json.Unmarshal([]byte(data), res)
			if err == nil {
				return nil
			}
		}
	}

	list := make([]models.UserNotice, 10)
	var total int64
	var unReadCnt int64
	db := s.DB().Where("team_id = ?", req.TeamId).Where("user_id = ?", req.UserId)
	if req.Status != 0 {
		db.Where("status = ?", req.Status)
	}
	if err := db.Order("status asc ,created_at desc").Offset(req.GetOffset()).Limit(req.GetSize()).
		Find(&list).Offset(-1).Limit(-1).Count(&total).Error; err != nil {
		return err
	}
	if req.Status == 0 {
		if err := db.Where("status = ?", 1).Count(&unReadCnt).Error; err != nil {
			return err
		}
	} else {
		unReadCnt = total
	}

	*res = dto.NoticeDto{
		Key:   "1",
		Name:  "通知",
		Count: unReadCnt,
		Total: total,
	}
	for _, v := range list {
		var item dto.NoticeItem
		copier.Copy(&item, v)
		item.Type = 1
		item.NoticeType = v.NoticeType
		item.CreatedAt = v.CreatedAt.Unix()
		res.List = append(res.List, item)
	}
	if req.GetOffset() == 0 {
		core.Cache.Set(ckey.MemberNotice(req.TeamId, req.UserId), res, time.Minute*10)
	}
	return nil
}

func (s *UserNoticeService) ReadUserNotice(req *dto.ReadNoticeDto, reqId string, teamId int, userId int) error {
	var nu models.UserNotice
	if err := s.DB().First(&nu, req.Id).Error; err != nil {
		return err
	}
	if nu.UserId != userId {
		return fmt.Errorf("无权限")
	}
	if nu.Status == 1 {
		nu.Status = 2
		nu.UpdatedAt = time.Now()
		if err := s.DB().Updates(&nu).Error; err != nil {
			return err
		}
	}
	core.Cache.Del(ckey.MemberNotice(nu.TeamId, nu.UserId))
	return nil
}

func (s *UserNoticeService) Create(req *dto.UserNoticeDto) error {
	var model models.UserNotice
	copier.Copy(&model, req)
	model.Status = 1
	if err := s.DB().Create(&model).Error; err != nil {
		return err
	}
	core.Cache.Del(ckey.MemberNotice(model.TeamId, model.UserId))
	return nil
}

func (s *UserNoticeService) Update(req *dto.UserNoticeDto) error {
	var model models.UserNotice
	copier.Copy(&model, req)
	if err := s.DB().Updates(&model).Error; err != nil {
		return err
	}
	core.Cache.Del(ckey.MemberNotice(model.TeamId, model.UserId))
	return nil
}

func (s *UserNoticeService) Del(id int) error {
	var model models.UserNotice
	if err := s.DB().First(&model, id).Error; err != nil {
		return err
	}
	if err := s.DB().Delete(&model).Error; err != nil {
		return err
	}
	core.Cache.Del(ckey.MemberNotice(model.TeamId, model.UserId))
	return nil
}
