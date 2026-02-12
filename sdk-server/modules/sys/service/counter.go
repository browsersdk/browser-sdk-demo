package service

import (
	"dilu/modules/sys/models"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"gorm.io/gorm"
)

type CounterService struct {
	*base.BaseService
}

var SerCounter = CounterService{
	base.NewService("sys"),
}

func (s *CounterService) Incr(key string, min int64) int64 {
	cacheKey := "counter:" + key
	r, err := core.Cache.Incr(cacheKey)
	//fmt.Println(r, err)
	if err != nil {
		var c models.Counter
		if err := s.DB().Where("`key` = ?", key).Find(&c).Error; err != nil {
			return 0
		}
		if c.Key == "" {
			c.Key = key
			c.Seq = min
			go s.DB().Create(&c)
		} else {
			s.DB().Model(&models.Counter{}).Where("`key` = ?", key).Update("seq", gorm.Expr("seq+1"))
		}
		if c.Key != "" {
			r = c.Seq + 1
			core.Cache.Set(cacheKey, r, 0)
		}
		return r
	} else {
		if r < min {
			var c models.Counter
			if err := s.DB().Where("`key` = ?", key).Find(&c).Error; err != nil {
				return 0
			}
			if c.Key != "" {
				r = c.Seq + 1
			} else {
				c.Key = key
				c.Seq = min
				go s.DB().Create(&c)
			}
			core.Cache.Set(cacheKey, r, 0)
		}
		go func() {
			s.DB().Model(&models.Counter{}).Where("`key` = ?", key).Update("seq", r)
		}()
		return r
	}
}
