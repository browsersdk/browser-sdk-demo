package sms

import (
	"dilu/common/third/sms/aliyun"
	"dilu/common/third/sms/config"
	"dilu/common/third/sms/tencent"
	"math/rand/v2"
)

type SmsSend interface {
	Send(phone string, code string, tmpId string) error
}

func Setup(cfg *config.SmsConfig) SmsSend {
	switch cfg.Select {
	case config.SmsAliyun:
		return aliyun.NewAliyunSms(cfg)
	case config.SmsTencent:
		return tencent.NewTencentSms(cfg)
	}
	return nil
}

var (
	numberic = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
)

func GenerateSmsCode(length int) string {
	sb := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		c := numberic[rand.IntN(len(numberic))]
		sb = append(sb, c)
	}
	return string(sb)
}
