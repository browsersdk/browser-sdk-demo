package service

import (
	"dilu/common/config"
	"dilu/common/third/sms"
	"dilu/modules/sys/models"
	"fmt"
	"sync"
	"time"

	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
)

type SysSms struct {
	*base.BaseService
}

var SerSms = SysSms{
	base.NewService(consts.DB_DEF),
}

var onceSms sync.Once
var smsSend sms.SmsSend

func (e *SysSms) getSms() sms.SmsSend {
	onceSms.Do(func() {
		smsSend = sms.Setup(&config.Ext.Sms)
	})
	return smsSend
}

func (e *SysSms) Send(phone string, tempId string) error {
	var err error
	code := sms.GenerateSmsCode(6)
	data := models.SysSms{
		Phone:     phone,
		Code:      code,
		Type:      "Code",
		Status:    0,
		UseStatus: 0,
	}
	data.CreatedAt = time.Now().Unix()
	data.UpdatedAt = data.CreatedAt
	err = core.DB().Create(&data).Error
	if err != nil {
		return err
	}
	return e.getSms().Send(phone, code, tempId)
}

// 验证
func (e *SysSms) Verify(phone, code string) bool {
	if core.Cfg.Server.Mode != "prod" && code == "666666" {
		return true
	}
	var err error
	var data models.SysSms
	err = core.DB().Model(&data).Where("phone = ? ", phone).Order("id desc").Find(&data).Error
	if err != nil {
		return false
	}
	if data.Id == 0 {
		fmt.Println("没有找到验证码", phone, code)
		return false
	}
	if data.UseStatus == 1 {
		fmt.Println("已使用")
		return false
	}
	if data.CreatedAt+int64(time.Minute.Minutes()*10) < time.Now().Unix() {
		fmt.Println("1111")
		if code == data.Code {
			fmt.Println("111333")
			data.UseStatus = 1
			data.UpdatedAt = time.Now().Unix()
			core.DB().Save(&data)
			return true
		}
	}

	return false
}
