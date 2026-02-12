package models

import (
	"time"
)

// UploadLog
type UploadLog struct {
	Id         int       `json:"id" gorm:"type:bigint unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	Url        string    `json:"url" gorm:"type:varchar(255);comment:资源地址"`                          //资源地址
	Provider   string    `json:"provider" gorm:"type:varchar(16);comment:提供方"`                       //提供方
	Status     int8      `json:"status" gorm:"type:tinyint;comment:1使用中 -1 删除"`                      //1使用中 -1 删除
	SyncStatus int8      `json:"syncStatus" gorm:"type:tinyint;comment:同步状态"`                        //同步状态
	CreatedAt  time.Time `json:"createdAt" gorm:"type:datetime;comment:创建时间"`                        //创建时间
	UpdatedAt  time.Time `json:"updatedAt" gorm:"type:datetime;comment:更新时间"`                        //更新时间
}

const TBUploadLog = "upload_log"

func (UploadLog) TableName() string {
	return TBUploadLog
}
