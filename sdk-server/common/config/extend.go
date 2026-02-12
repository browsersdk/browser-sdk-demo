package config

import (
	fsCfg "dilu/common/third/file_store/cfg"
	smsCfg "dilu/common/third/sms/config"

	"github.com/baowk/dilu-rd/config"
)

var Ext *Extend

func init() {
	Ext = new(Extend)
}

type Extend struct {
	AesKey        string             `mapstructure:"aes-key" json:"aes-key" yaml:"aes-key"` //aes加密key
	RdConfig      config.Config      `mapstructure:"rd-config" json:"rd-config" yaml:"rd-config"`
	OneClickLogin OneClickLogin      `mapstructure:"one-click-login" json:"one-click-login" yaml:"one-click-login"`
	Feishu        Feishu             `mapstructure:"feishu" json:"feishu" yaml:"feishu"`
	FSCfg         fsCfg.FileStoreCfg `mapstructure:"file-store" json:"file-store" yaml:"file-store"`
	Sms           smsCfg.SmsConfig   `mapstructure:"sms" json:"sms" yaml:"sms"`
	OneClick      string             `mapstructure:"one-click" json:"one-click" yaml:"one-click"` // 一键登录的配置
}

type ChatService struct {
	Endpoint string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	ApiKey   string `mapstructure:"api-key" json:"api-key" yaml:"api-key"`
}

type OneClickLogin struct {
	AppKey string `mapstructure:"app-key" json:"app-key" yaml:"app-key"`
	AppId  string `mapstructure:"app-id" json:"app-id" yaml:"app-id"`
}

type Feishu struct {
	AppId             string `mapstructure:"app-id" json:"app-id" yaml:"app-id"`
	AppSecret         string `mapstructure:"app-secret" json:"app-secret" yaml:"app-secret"`
	VerificationToken string `mapstructure:"verification-token" json:"verification-token" yaml:"verification-token"`
	EncryptKey        string `mapstructure:"encrypt-key" json:"encrypt-key" yaml:"encrypt-key"`
	RedirectUri       string `mapstructure:"redirect-uri" json:"redirect-uri" yaml:"redirect-uri"`
}

// type DingCfg struct {
// 	AgentId   string `mapstructure:"agent-id" json:"agent-id" yaml:"agent-id"`
// 	AppKey    string `mapstructure:"app-key" json:"app-key" yaml:"app-key"`
// 	AppSecret string `mapstructure:"app-secret" json:"app-secret" yaml:"app-secret"`
// 	CropId    string `mapstructure:"crop-id" json:"crop-id" yaml:"crop-id"`
// }

// type WechatMp struct {
// 	AppId          string `mapstructure:"app-id" json:"app-id" yaml:"app-id"`
// 	AppSecret      string `mapstructure:"app-secret" json:"app-secret" yaml:"app-secret"`
// 	WxToken        string `mapstructure:"wx-token" json:"wx-token" yaml:"wx-token"`
// 	EncodingAESKey string `mapstructure:"encoding-aes-key" json:"encoding-aes-key" yaml:"encoding-aes-key"`
// }

// type Ai struct {
// 	Ali Ali `mapstructure:"ali" json:"ali" yaml:"ali"`
// }

// type Ali struct {
// 	SK string `mapstructure:"sk" json:"sk" yaml:"sk"`
// }
