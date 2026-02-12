package service

import (
	"dilu/modules/notice/models"
	"time"

	"github.com/baowk/dilu-core/core/base"
)

type PubNoticeService struct {
	*base.BaseService
}

var SerPubNotice = PubNoticeService{
	base.NewService("notice"),
}

func (s *PubNoticeService) GetNotice(teamId int) []models.PubNotice {
	teamIds := []int{teamId}
	if teamId != 0 {
		teamIds = append(teamIds, 0)
	}
	var notice []models.PubNotice
	if err := s.DB().Where("expired > ?", time.Now()).Where("status = ?", 1).
		Where("team_id in ?", teamIds).Limit(100).Order("id desc").Find(&notice).Error; err != nil {
		return nil
	}
	return notice
}
