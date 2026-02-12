package config

type SmsSelect string

const (
	SmsAliyun  SmsSelect = "ali"
	SmsTencent SmsSelect = "tencent"
)

type SmsConfig struct {
	Select     SmsSelect `json:"select" yaml:"select" mapstructure:"select"`
	AppId      string    `json:"app-id" yaml:"app-id" mapstructure:"app-id"`
	SecretId   string    `json:"secret-id" yaml:"secret-id" mapstructure:"secret-id"`
	SecretKey  string    `json:"secret-key" yaml:"secret-key" mapstructure:"secret-key"`
	Sign       string    `json:"sign" yaml:"sign" mapstructure:"sign"`
	TemplateId string    `json:"template-id" yaml:"template-id" mapstructure:"template-id"`
	Region     string    `json:"region" yaml:"region" mapstructure:"region"`
}
