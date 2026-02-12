package cfg

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type StoreType string

const (
	LocalStore   StoreType = "local"
	OssStore     StoreType = "oss"
	TencentStore StoreType = "cos"
	AWSStore     StoreType = "s3"
)

type FileStoreCfg struct {
	Select    StoreType `mapstructure:"select" json:"select" yaml:"select"`             //选择存储类型
	SecretID  string    `mapstructure:"secret-id" json:"secret-id" yaml:"secret-id"`    //访问id accesskey
	SecretKey string    `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"` //安全码
	Bucket    string    `mapstructure:"bucket" json:"bucket" yaml:"bucket"`             //桶
	Region    string    `mapstructure:"region" json:"region" yaml:"region"`             //区域
	Domain    string    `mapstructure:"domain" json:"domain" yaml:"domain"`             //cdn域名
	Prefix    string    `mapstructure:"prefix" json:"prefix" yaml:"prefix"`             //目录前缀
	Internal  bool      `mapstructure:"internal" json:"internal" yaml:"internal"`       //内网 默认false
}

func (cfg *FileStoreCfg) GetFileKey(prefix, ext string) string {

	if prefix == "" {
		prefix = cfg.Prefix
	}
	if !strings.HasSuffix(prefix, "/") {
		prefix = prefix + "/"
	}

	fileKey := prefix + time.Now().Format("2006/01") + "/" + uuid.NewString()
	if ext != "" {
		fileKey = fileKey + "." + ext
	}
	return fileKey
}
